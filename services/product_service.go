package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Godfredasare/go-ecommerce/database"
	"github.com/Godfredasare/go-ecommerce/models"
	"go.mongodb.org/mongo-driver/bson"
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

func FindAll() ([]models.Product, error) {
	col := database.Collection(name)

	cur, err := col.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("Error Finding product %v", err)
		return nil, err
	}

	var products []models.Product

	for cur.Next(context.Background()) {
		var product models.Product
		err := cur.Decode(&product)
		if err != nil {
			log.Printf("Error decoding product %v", err)
			return nil, err
		}

		products = append(products, product)

	}

	return products, nil

}

func FindOne(productID string) (*models.Product, error) {
	col := database.Collection(name)

	id, _ := primitive.ObjectIDFromHex(productID)

	var product models.Product

	filter := bson.D{{Key: "_id", Value: id}}

	err := col.FindOne(context.Background(), filter).Decode(&product)
	if err != nil {
		log.Printf("Error finding one product %v", err)
		return nil, err
	}

	return &product, nil

}
