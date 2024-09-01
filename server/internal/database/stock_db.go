package database

import (
	"context"
	"fmt"
	"log"

	"github.com/KladeRe/stock-server/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StockConfig struct {
	// ID is automatically generated by database
	Id           primitive.ObjectID `bson:"_id,omitempty"`
	Symbol       string             `bson:"symbol"`
	Value        float32            `bson:"value"`
	Buy          bool               `bson:"buy"`
	Notification string             `bson:"notification"`
}

func EstablishDBClient() (*mongo.Client, error) {
	DB_HOST, hostError := utils.GetEnvVariable("DB_HOST")

	DB_PORT, portError := utils.GetEnvVariable(("DB_PORT"))

	if portError != nil {
		return &mongo.Client{}, portError
	}

	if hostError != nil {
		return &mongo.Client{}, hostError
	}

	// MongoDB URI (connect to your local MongoDB instance)
	clientOptions := options.Client().ApplyURI("mongodb://" + DB_HOST + ":" + DB_PORT)

	// Connect to MongoDB
	client, connectionErr := mongo.Connect(context.TODO(), clientOptions)
	if connectionErr != nil {
		return &mongo.Client{}, connectionErr
	}

	// Check the connection
	pingErr := client.Ping(context.TODO(), nil)
	if pingErr != nil {
		return &mongo.Client{}, pingErr
	}

	return client, nil
}

func GetCollection(client *mongo.Client) *mongo.Collection {
	// Access a specific database and collection
	collection := client.Database("stockconfig").Collection("stocks")
	return collection
}

func CloseClient(client *mongo.Client) (string, error) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		return "", err
	}

	return "Connection to DB closed", nil
}

func AddDocument(collection *mongo.Collection, newConfig StockConfig) {

	// Insert the document into the collection
	insertResult, err := collection.InsertOne(context.TODO(), newConfig)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

}

func DeleteDocument(collection *mongo.Collection, id primitive.ObjectID) (string, error) {
	filter := bson.M{"_id": id}
	deleteResult, deleteErr := collection.DeleteOne(context.TODO(), filter)
	if deleteErr != nil {
		return "", deleteErr
	}
	if deleteResult.DeletedCount > 0 {
		return "Deleted document successfully", nil
	} else {
		return "No document found with the ID given", nil
	}

}

func GetAllDocuments(collection *mongo.Collection) ([]StockConfig, error) {
	results := []StockConfig{}
	cursor, cursorErr := collection.Find(context.TODO(), bson.D{})
	if cursorErr != nil {
		return []StockConfig{}, cursorErr
	}
	for cursor.Next(context.TODO()) {
		var result StockConfig
		nextErr := cursor.Decode(&result)
		if nextErr != nil {
			return []StockConfig{}, nextErr
		}
		results = append(results, result)

	}

	return results, nil
}