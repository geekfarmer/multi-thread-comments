package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//App Model Sctructure
type App struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	CreatedAt          time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt          time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	AppPackageName     string             `json:"app_package_name,omitempty" bson:"app_package_name,omitempty"`
	Name               string             `json:"name" bson:"name"`
	AppColor           string             `json:"app_color" bson:"app_color"`
	AppBackgroundColor string             `json:"app_bg" bson:"app_bg"`
	ActiveStatus       string             `json:"status,omitempty" bson:"status,omitempty"` //active, pending, cancelled
	Description        string             `json:"description,omitempty" bson:"description,omitempty"`
	AppLogo            string             `json:"app_logo,omitempty" bson:"app_logo,omitempty"`
	AddedBy            string             `json:"added_by,omitempty" bson:"added_by,omitempty"`
}

// Installed apps for specific user
type InstalledApps struct {
	ID                 primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	CreatedAt          time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt          time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
	AppPackageName     string             `json:"app_package_name,omitempty" bson:"app_package_name,omitempty"`
	Name               string             `json:"name" bson:"name"`
	AppColor           string             `json:"app_color" bson:"app_color"`
	AppBackgroundColor string             `json:"app_bg" bson:"app_bg"`
	ActiveStatus       string             `json:"status,omitempty" bson:"status,omitempty"` //active, pending, cancelled
	Description        string             `json:"description,omitempty" bson:"description,omitempty"`
	AppLogo            string             `json:"app_logo,omitempty" bson:"app_logo,omitempty"`
	SubscriptionData   string             `json:"subscritpion_date" bson:"subscritpion_date,omitempty"`
	SubscriptionType   string             `json:"subscritpion_type" bson:"subscritpion_type,omitempty"` //monthly,yearly
	Price              string             `json:"price" bson:"price,omitempty"`
	Currency           string             `json:"currency" bson:"currency,omitempty"`
}

//Question
//App is slow, hwo can i improve my app
//server fetch app of image from list view
//best way viewholder patter or not.
//question like do you use library.
