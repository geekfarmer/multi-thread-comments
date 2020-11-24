package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/geekfarmer/multi-thread-comments/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type appRepo struct {
	CollectionName     string
	CollectionNameUser string
	Database           *mongo.Database
}

func NewAppMongoRepository(data *mongo.Database) AppRepository {

	return &appRepo{
		CollectionName:     "app-collection",
		CollectionNameUser: "user-collection",
		Database:           data}
}

func (r *appRepo) FindAllApps() ([]entity.App, error) {
	ctx := context.Background()

	var apps []entity.App

	cursor, err := r.Database.Collection(r.CollectionName).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var app entity.App
		cursor.Decode(&app)
		apps = append(apps, app)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return apps, nil
}

func (r *appRepo) GetAllRequestedApps() ([]entity.App, error) {
	ctx := context.Background()

	var apps []entity.App

	cursor, err := r.Database.Collection(r.CollectionName).Find(ctx, bson.D{{"status", "pending"}})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var app entity.App
		cursor.Decode(&app)
		apps = append(apps, app)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return apps, nil
}
func (r *appRepo) CreateApp(app *entity.App) (*entity.App, error) {
	ctx := context.Background()

	_, err := r.Database.Collection(r.CollectionName).InsertOne(ctx, app)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return app, nil
}

func (r *appRepo) CreateApps(apps []entity.App) ([]entity.App, error) {
	ctx := context.Background()

	var allApps []interface{}

	for _, app := range apps {
		app.CreatedAt = time.Now()
		app.UpdatedAt = time.Now()
		allApps = append(allApps, app)
	}

	_, err := r.Database.Collection(r.CollectionName).InsertMany(ctx, allApps)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return apps, nil
}

func (r *appRepo) UpdateAppsForUser(user *entity.User) (*entity.User, error) {
	ctx := context.Background()

	filter := bson.M{"user_token": user.UserToken}
	update := bson.M{
		"$set": user,
	}

	_, err := r.Database.Collection(r.CollectionNameUser).UpdateOne(ctx, filter, update)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
