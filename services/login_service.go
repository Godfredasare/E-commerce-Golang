package services

import (
	"context"
	"errors"
	"log"

	"github.com/Godfredasare/go-ecommerce/database"
	"github.com/Godfredasare/go-ecommerce/models"
	"github.com/Godfredasare/go-ecommerce/utils"
	"go.mongodb.org/mongo-driver/bson"
)

type LoginModel struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required"`
}

func ValidCredenial(u *LoginModel) error {
	col := database.Collection("users")

	var user *models.Users
	filter := bson.D{{Key: "email", Value: u.Email}}

	err := col.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("Error %v", err)
		return errors.New("invalid email/password")
	}

	password := utils.CompareHashPassword(u.Password, user.Password)
	if !password {
		return errors.New("invalid email/password")
	}

	return err
}
