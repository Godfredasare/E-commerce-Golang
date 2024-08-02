package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" form:"name" bson:"name" validate:"required"`
	Description string             `json:"description,omitempty" form:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price" form:"price" bson:"price" validate:"required"`
	Currency    string             `json:"currency" form:"currency" bson:"currency"`
	Stock       int64              `json:"stock" form:"stock" bson:"stock"  validate:"required"`
	Category    string             `json:"category" form:"category" bson:"category"  validate:"required"`
	ImagesID    string             `json:"images_id" form:"images" bson:"images_id"`
	UserId      primitive.ObjectID `json:"user_id" form:"user_id" bson:"user_id"`
	CreatedAt   string             `json:"created_at" form:"created_at" bson:"created_at"`
	UpdatedAt   string             `json:"updated_at" form:"updated_at" bson:"updated_at"`
}
