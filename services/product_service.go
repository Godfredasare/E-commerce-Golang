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

func setDefault(p *models.Product) {
	if p.Currency == "" {
		p.Currency = "USD"
	}
}

func CreateProduct(p *models.Product) error {
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

		defer cur.Close(context.Background())

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

func Update(productID string, p *models.Product) (int64, error) {
	col := database.Collection(name)

	id, _ := primitive.ObjectIDFromHex(productID)

	filter := bson.D{{Key: "_id", Value: id}}

	setDefault(p)

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: p.Name},
		{Key: "description", Value: p.Description},
		{Key: "price", Value: p.Price},
		{Key: "currency", Value: p.Currency},
		{Key: "stock", Value: p.Stock},
		{Key: "updated_at", Value: time.Now().Format(time.RFC3339)},
	}}}

	result, err := col.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Printf("Error updating product %v", err)
		return 0, err
	}

	return result.ModifiedCount, nil

}

func Delete(productID string) (int64, error) {
	col := database.Collection(name)

	id, _ := primitive.ObjectIDFromHex(productID)

	filter := bson.D{{Key: "_id", Value: id}}
	result, err := col.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Printf("Error updating product %v", err)
		return 0, err
	}

	return result.DeletedCount, nil
}
