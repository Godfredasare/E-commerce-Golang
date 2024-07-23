package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Price       float64            `json:"price" bson:"price" validate:"required"`
	Currency    string             `json:"currency" bson:"currency"`
	Stock       int64              `json:"stock" bson:"stock"  validate:"required"`
	// Images      []string           `json:"images" validate:"required"`
	UserId    primitive.ObjectID `json:"user_id" bson:"user_id" validate:"required"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
	UpdatedAt string             `json:"updated_at" bson:"updated_at"`
}
