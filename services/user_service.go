package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Godfredasare/go-ecommerce/database"
	"github.com/Godfredasare/go-ecommerce/models"
	"github.com/Godfredasare/go-ecommerce/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	userDB = "users"
)

func CreateUser(u *models.Users) error {
	col := database.Collection(userDB)

	u.ID = primitive.NewObjectID()
	u.CreatedAt = u.ID.Timestamp().String()
	u.UpdatedAt = u.ID.Timestamp().String()

	//email exist
	filter := bson.D{{Key: "email", Value: u.Email}}
	count, err := col.CountDocuments(context.Background(), filter)
	if err != nil {
		log.Printf("Error %v", err)
		return errors.New("error with count document")
	}

	if count > 0 {
		return errors.New("email already exists")
	}


	hassPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Printf("Error %v", err)
		return errors.New("error hashing password")
	}

	u.Password = hassPassword

	result, err := col.InsertOne(context.Background(), u)
	if err != nil {
		log.Printf("Error inserting to users %v", err)
		return errors.New("can not insert user")
	}

	fmt.Println(result.InsertedID)

	return nil

}

func FindAllUsers() ([]models.Users, error) {
	col := database.Collection(userDB)
	cur, err := col.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Error getting all to users %v", err)
		return nil, err
	}

	var users []models.Users
	for cur.Next(context.Background()) {
		var user models.Users
		err := cur.Decode(&user)
		if err != nil {
			log.Printf("Error %v", err)
			return nil, err
		}

		users = append(users, user)

	}

	return users, nil
}

func UpdateUser(userId string, u *models.Users) (int64, error) {
	col := database.Collection(userDB)
	id, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "username", Value: u.Username},
		{Key: "email", Value: u.Email},
		{Key: "password", Value: u.Password},
		{Key: "updated_at", Value: u.ID.Timestamp().String()},
	}}}

	result, err := col.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error updating user %v", err)
		return 0, err
	}

	return result.ModifiedCount, nil
}
