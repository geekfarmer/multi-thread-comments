package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	Phone      string             `json:"phone_number" bson:"phone_number"`
	IsVerified bool               `json:"isVerified" bson:"isVerified"`
	UserToken  string             `json:"user_token" bson:"user_token"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Apps       []InstalledApps    `json:"apps" bson:"apps"`
}
