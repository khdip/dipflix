package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DBInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Unable to find .env file")
	}

	mongoDB := os.Getenv("MONGODB_URI")
	if mongoDB == "" {
		log.Fatal("MONGODB_URI not set!")
	}

	clientOptions := options.Client().ApplyURI(mongoDB)

	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(collectionName string) *mongo.Collection {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Unable to find .env file")
	}

	dbName := os.Getenv("DATABASE_NAME")

	collection := Client.Database(dbName).Collection(collectionName)
	return collection
}
