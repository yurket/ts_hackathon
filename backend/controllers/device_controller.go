package controllers

import (
	"context"
	// "fmt"
	"log"
	"time"
	// "iot/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Device struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	Tags        []string           `bson:"tags,omitempty"`
}

func GetDevices(c *gin.Context) {
	db := c.MustGet("db").(*mongo.Database)
	device_collection := db.Collection("devices")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := device_collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}
	var devices []bson.M
	if err = cursor.All(ctx, &devices); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"devices": devices})
}

func MakeDevice(c *gin.Context) {
	db := c.MustGet("db").(*mongo.Database)
	device_collection := db.Collection("devices")

	var device Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	deviceResult, err := device_collection.InsertOne(ctx, device)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"_id": deviceResult.InsertedID})
}

func DeleteDevice(c *gin.Context) {
	db := c.MustGet("db").(*mongo.Database)
	device_collection := db.Collection("devices")

	id, err := primitive.ObjectIDFromHex(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := device_collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": result.DeletedCount})
}
