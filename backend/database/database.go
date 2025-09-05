package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Collections collections
var Client *mongo.Client

type collections struct {
	Players *mongo.Collection
	Scores  *mongo.Collection
	Changes *mongo.Collection
}

// Initialise the database connection and fetch the collections
func Initialise() {
	databaseURI := os.Getenv("MONGO_URI")
	if databaseURI == "" {
		log.Fatal("Missing MONGO_URI environment variable")
	}
	client, err := mongo.Connect(options.Client().ApplyURI(databaseURI))
	if err != nil {
		panic(err)
	}
	Client = client
	databaseName := os.Getenv("MONGO_DATABASE")
	// Set the collections within the collections struct
	collections := collections{
		Players: client.Database(databaseName).Collection("players"),
		Scores:  client.Database(databaseName).Collection("scores"),
		Changes: client.Database(databaseName).Collection("changes"),
	}
	Collections = collections
}

// Fetch a document from the provided collection using the provided filter
func fetchDocument(collection *mongo.Collection, filter bson.M) (*mongo.SingleResult, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result := collection.FindOne(context, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return result, nil
}

// Fetch multiple documents from the provided collection using the provided filter, skip and limit
func fetchDocuments(collection *mongo.Collection, filter bson.M, options ...options.Lister[options.FindOptions]) (*mongo.Cursor, error) {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return collection.Find(
		context,
		filter,
		options...,
	)
}

// Insert a document into the provided collection
func InsertDocument(collection *mongo.Collection, document interface{}, opts ...options.InsertOneOptions) error {
	bsonBytes, err := bson.Marshal(document)
	if err != nil {
		return fmt.Errorf("error marshaling document: %v", err)
	}
	var bsonDoc bson.M
	err = bson.Unmarshal(bsonBytes, &bsonDoc)
	if err != nil {
		return fmt.Errorf("error converting to bson.M: %v", err)
	}
	_, err = collection.InsertOne(context.Background(), bsonDoc)
	if err != nil {
		return fmt.Errorf("error inserting document: %v", err)
	}
	return nil
}

// Delete the provided document
func DeleteDocument(collection *mongo.Collection, filter bson.M) error {
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error deleting document: %v", err)
	}
	return nil
}

// Update the provided document
func UpdateDocument(collection *mongo.Collection, filter bson.M, update bson.M) error {
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error updating document: %v", err)
	}
	return nil
}
