package repository

import (
	"context"
	"fmt"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepo struct {
	CollectionName string
	Database       *mongo.Database
}

func NewUserMongoRepository(data *mongo.Database) UserRepository {
	// Access a MongoDB collection through a database
	coll := data.Collection("users")
	_, err := coll.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		fmt.Println(err)
	}
	return &userRepo{CollectionName: "users", Database: data}
}

func (r *userRepo) Create(user *entity.User) (*entity.User, error) {
	ctx := context.Background()
	_, err := r.Database.Collection(r.CollectionName).InsertOne(ctx, user)
	if err != nil {
		fmt.Println(err)
	}
	return user, err
}
func (r *userRepo) UpdateUser(user *entity.User) (*entity.User, error) {
	ctx := context.Background()
	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": user,
	}

	_, err := r.Database.Collection(r.CollectionName).UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}
	return user, nil
}

func (r *userRepo) FindAll() ([]entity.User, error) {
	ctx := context.Background()
	var users []entity.User
	cursor, err := r.Database.Collection(r.CollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user entity.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepo) FindByID(id string) (*entity.User, error) {
	ctx := context.Background()
	var user *entity.User
	err := r.Database.Collection(r.CollectionName).FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepo) Delete(id string) error {
	ctx := context.Background()
	_, err := r.Database.Collection(r.CollectionName).DeleteOne(ctx, bson.M{"_id": id})
	return err
}
