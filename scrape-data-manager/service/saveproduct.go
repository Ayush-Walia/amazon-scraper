package service

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/dto"
	"github.com/Ayush-Walia/amazon-scraper/scrape-data-manager/utils"
)

// SaveData takes the scrapedData and stores it in MongoDB.
func SaveData(scrapedData dto.ScrapedData) dto.ResponseMessage {
	var saveProductResponse dto.ResponseMessage
	scrapedData.LastUpdatedTime = time.Now()

	mongoHost := utils.GetEnv("MONGO_HOST", "mongo")
	mongoPort := utils.GetEnv("MONGO_PORT", "27017")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongoHost + ":" + mongoPort))
	if err != nil {
		log.Println(err)
		saveProductResponse.Success = false
		saveProductResponse.Message = err.Error()
		return saveProductResponse
	}
	ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)

	// Connect to the client
	err = client.Connect(ctx)
	if err != nil {
		log.Println(err)
		saveProductResponse.Success = false
		saveProductResponse.Message = err.Error()
		return saveProductResponse
	}
	log.Println("Connected to MongoDB!")
	defer client.Disconnect(ctx)

	// Access MongoDB collection through a database
	collection := client.Database("scrapedb").Collection("product_details")

	// Update document in Mongo.
	filter := bson.M{"url": scrapedData.URL}
	update := bson.M{
		"$set": bson.M{
			"product":         scrapedData.Product,
			"lastupdatedtime": scrapedData.LastUpdatedTime,
		},
	}
	updated, _ := collection.UpdateOne(ctx, filter, update)
	if updated.ModifiedCount == 0 {
		_, err := collection.InsertOne(ctx, scrapedData)
		if err != nil {
			log.Println(err)
			saveProductResponse.Success = false
			saveProductResponse.Message = err.Error()
			return saveProductResponse
		}
	}
	log.Println("Saved data in DB successfully.")
	saveProductResponse.Success = true
	saveProductResponse.Message = "Saved data in DB successfully."
	return saveProductResponse
}
