package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/leulshawel/go-blogger-app/routes"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	mongoConnection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal("Couldn't connect to mongoDB")
	}

	fmt.Println("Connected to DB")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pingError := mongoConnection.Ping(ctx, nil)
	if pingError != nil {
		log.Fatal("DB Connection lost")
	}

	app, error := routes.New(mongoConnection)
	if error != nil {
		fmt.Println("Error returned from blogger app")
	}
	app.Listen("127.0.0.1:8080")

}
