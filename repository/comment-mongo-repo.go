package repository

import (
	"context"
	"fmt"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type commentRepo struct {
	CollectionName string
	Database       *mongo.Database
}

var lookUpAuthor = bson.D{{"$lookup", bson.D{
	{"from", "users"},
	{"localField", "author"},
	{"foreignField", "_id"},
	{"as", "authorDetails"},
}}}
var unwindAuthor = bson.D{{"$unwind", bson.D{
	{"path", "$authorDetails"},
	{"preserveNullAndEmptyArrays", true},
}}}
var projectStage = bson.D{{"$project", bson.D{
	{"authorDetails.email", 0},
	{"authorDetails.created_at", 0},
	{"authorDetails.updated_at", 0},
}}}

func NewCommentMongoRepository(data *mongo.Database) CommentRepository {
	// Access a MongoDB collection through a database
	return &commentRepo{CollectionName: "comments", Database: data}
}

func (r *commentRepo) CreateComment(comment *entity.Comment, userID string) (*entity.Comment, error) {
	ctx := context.Background()
	comment.Author = userID
	_, err := r.Database.Collection(r.CollectionName).InsertOne(ctx, comment)
	if err != nil {
		fmt.Println(err)
	}
	return comment, err
}

func (r *commentRepo) UpdateComment(comment *entity.Comment) (*entity.Comment, error) {
	ctx := context.Background()
	filter := bson.M{"_id": comment.ID}
	update := bson.M{
		"$set": comment,
	}

	_, err := r.Database.Collection(r.CollectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return comment, nil
}

func (r *commentRepo) UpdateChildComments(id string, comment *entity.Comment) (*entity.Comment, error) {
	var parentComment *entity.Comment
	ctx := context.Background()
	filter := bson.M{"_id": id}
	err := r.Database.Collection(r.CollectionName).FindOne(ctx, filter).Decode(&parentComment)
	parentComment.ChildComments = append(parentComment.ChildComments, comment.ID)
	update := bson.M{
		"$set": parentComment,
	}

	_, err = r.Database.Collection(r.CollectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return parentComment, nil
}

func (r *commentRepo) FindAllComments() ([]entity.Comment, error) {
	ctx := context.Background()
	var comments []entity.Comment
	cursor, err := r.Database.Collection(r.CollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var comment entity.Comment
		cursor.Decode(&comment)
		comments = append(comments, comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *commentRepo) FindCommentByID(id string) (*entity.Comment, error) {
	ctx := context.Background()
	var comment *entity.Comment
	err := r.Database.Collection(r.CollectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentRepo) DeleteComment(id string) error {
	ctx := context.Background()
	_, err := r.Database.Collection(r.CollectionName).DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *commentRepo) GetChildren(ids []string) []entity.Comment {
	ctx := context.Background()
	var comments []entity.Comment
	for _, id := range ids {
		commentMatchStage := bson.D{{"$match", bson.D{
			{"_id", id},
		}}}
		var comment *entity.Comment
		var temp []entity.Comment
		cursor, err := r.Database.Collection(r.CollectionName).Aggregate(ctx, mongo.Pipeline{commentMatchStage, lookUpAuthor, unwindAuthor, projectStage})
		if err = cursor.All(ctx, &temp); err != nil {
			return comments
		}
		comment = &temp[0]
		if len(comment.ChildComments) > 0 {
			comment.ChildCommentsList = r.GetChildren(comment.ChildComments)
		}
		comments = append(comments, *comment)
	}
	return comments
}

func (r *commentRepo) FindCommentsByPostID(id string) ([]entity.Comment, error) {
	ctx := context.Background()
	var comments []entity.Comment
	matchStage := bson.D{{"$match", bson.D{
		{"postId", id},
		{"parentId", bson.D{{"$exists", false}}},
	}}}

	cursor, err := r.Database.Collection(r.CollectionName).Aggregate(ctx, mongo.Pipeline{matchStage, lookUpAuthor, unwindAuthor, projectStage})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var comment entity.Comment
		cursor.Decode(&comment)
		comment.ChildCommentsList = r.GetChildren(comment.ChildComments)
		comments = append(comments, comment)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}
