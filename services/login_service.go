package services

import (
	"context"
	"log"

	"github.com/Godfredasare/go-ecommerce/database"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginModel struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}

func ValidCredenial(u *LoginModel) error {
	col := database.Collection("users")

	filter := bson.D{{Key: "email", Value: u.Email}}

	err := col.FindOne(context.Background(), filter).Decode(&u)
	if err != nil {
		log.Printf("Error %v", err)
		return err
	}

	return err
}
