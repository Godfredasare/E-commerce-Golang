package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Godfredasare/go-ecommerce/database"
	"github.com/Godfredasare/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var name = "products"

func setDefault(p models.Product) {
	if p.Currency == "" {
		p.Currency = "USD"
	}
}

func CreateProduct(p models.Product) error {
	col := database.Collection(name)

	setDefault(p)

	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now().Format(time.RFC3339)
	p.UpdatedAt = time.Now().Format(time.RFC3339)

	result, err := col.InsertOne(context.Background(), p)
	if err != nil {
		log.Printf("Error inserting product %v", err)
		return err
	}

	fmt.Println("Product inserted succesfully into db:", result.InsertedID)
	return nil
}
