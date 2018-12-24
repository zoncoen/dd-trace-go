package mongo_test

import (
	"context"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"

	mongotrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/mongodb/mongo-go-driver/mongo"
)

func Example() {
	// connect to MongoDB
	client, err := mongo.Connect(context.Background(), "mongodb://localhost:27017",
		options.Client().SetMonitor(mongotrace.NewMonitor()))
	if err != nil {
		panic(err)
	}
	db := client.Database("example")
	inventory := db.Collection("inventory")

	inventory.InsertOne(context.Background(), bson.D{
		{"item", "canvas"},
		{"qty", 100},
	})
}
