package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID          primitive.ObjectID `json:"id"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description,omitempty"`
	Stock       int64              `json:"stock"`
	Images      []string           `json:"images" validate:"required"`
	// User_id     int64              `json:"user_id" validate:"required"`
	Created_at  string             `json:"created_at"`
	Updated_at  string             `json:"updated_at"`
}
