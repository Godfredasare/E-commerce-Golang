package services

import (
	"context"
	"fmt"
	"log"

	"github.com/Godfredasare/go-ecommerce/database"
	"github.com/Godfredasare/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	userDB = "users"
)

func CreateUser(u *models.Users) error {
	col := database.Collection(userDB)

	u.ID = primitive.NewObjectID()
	u.CreatedAt= u.ID.Timestamp().String()
	u.UpdatedAt= u.ID.Timestamp().String()

	result, err := col.InsertOne(context.Background(), u)
	if err != nil {
		log.Fatalf("Error inserting to users %v", err)
	}

	fmt.Println(result.InsertedID)

	return nil

}
