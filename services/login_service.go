package services

import (
	"context"
	"errors"
	"log"

	"github.com/Godfredasare/go-ecommerce/database"
	"github.com/Godfredasare/go-ecommerce/models"
	"github.com/Godfredasare/go-ecommerce/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Email    string             `json:"email" bson:"email" validate:"required,email"`
	Password string             `json:"password" bson:"password" validate:"required"`
}

//check if user with email exist 
//then decode the user from the db to get the user to perform compare pasword and signJwt
func ValidCredenial(u *LoginModel) (*models.Users, error) {
	col := database.Collection("users")

	var user *models.Users
	filter := bson.D{{Key: "email", Value: u.Email}}

	err := col.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		log.Printf("Error %v", err)
		return nil, errors.New("invalid email/password")
	}

	password := utils.CompareHashPassword(u.Password, user.Password)
	if !password {
		return nil, errors.New("invalid email/password")
	}

	return user, err
}
