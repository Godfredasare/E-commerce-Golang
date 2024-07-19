package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id"`
	Username  string             `json:"username" bson:"username" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	Password  string             `json:"password" bson:"password" validate:"required"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at"`
}
