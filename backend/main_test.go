package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func TestMongoInteraction(t *testing.T) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo databases: ", databases)

	db := client.Database("my_db")
	collection1 := db.Collection("collection1")
	collection2 := db.Collection("collection2")

	// object1Result, err := collection1.InsertOne(ctx, bson.D{
	// 	{Key: "a", Value: "aa"},
	// 	{Key: "b", Value: "bb"},
	// 	{Key: "qqqqq", Value: 255555},
	// 	{Key: "as", Value: bson.A{1, 2, 3, 4, 5}},
	// })
	// fmt.Println(object1Result)

	// object2Result, err := collection2.InsertOne(ctx, bson.D{
	// 	{Key: "c", Value: "cc"},
	// 	{Key: "d", Value: "dd"},
	// 	{Key: "e", Value: "ee"},
	// })
	// fmt.Println(object2Result)

	cursor1, err := collection1.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var items []bson.M
	if err = cursor1.All(ctx, &items); err != nil {
		log.Fatal(err)
	}
	fmt.Println(items)

	cursor2, err := collection2.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var items2 []bson.M
	if err = cursor2.All(ctx, &items2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(items2)

	defer client.Disconnect(ctx)
}
