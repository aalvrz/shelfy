package database

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbURI = "mongodb://localhost:27017/"
const DatabaseName = "shelfy"

var ctx = context.TODO()
var MongoClient *mongo.Client

func ConnectDB() error {
	clientOptions := options.Client().ApplyURI(dbURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Create a default user
	// collection := client.Database(DatabaseName).Collection("users")
	// _, err = collection.InsertOne(ctx, model.User{
	// 	ID:       primitive.NewObjectID(),
	// 	Email:    "john@msn.com",
	// 	Password: "qwerty123!",
	// })
	// log.Info("Created default user")

	MongoClient = client
	return err
}
