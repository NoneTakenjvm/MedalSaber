package database

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Collections collections
var Client *mongo.Client

type collections struct {
	Players       *mongo.Collection
	Scores        *mongo.Collection
	CountryTopten *mongo.Collection
}

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
	databaseName := os.Getenv("MONG_DATABASE")
	// Set the collections within the collections struct
	collections := collections{
		Players:       client.Database(databaseName).Collection("players"),
		Scores:        client.Database(databaseName).Collection("scores"),
		CountryTopten: client.Database(databaseName).Collection("countrytopten"),
	}
	Collections = collections
}
