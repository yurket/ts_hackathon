package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Sensor struct {
	Time        string `json:"Time"`
	Temperature string `json:"Temperature"`
	Pressure    string `json:"Pressure"`
	Humidity    string `json:"Humidity"`
}

func GetData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	db := c.MustGet("db").(*mongo.Database)
	data_collection := db.Collection("data")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := data_collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	var sensor_data []bson.M
	if err = cursor.All(ctx, &sensor_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sensor_data})
}

func PostData(c *gin.Context) {
	fmt.Println("Header:", c.Request.Header)
	fmt.Println("Host:", c.Request.Host)
	fmt.Println("Body:", c.Request.Body)

	db := c.MustGet("db").(*mongo.Database)
	data_collection := db.Collection("data")

	var data Sensor
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	data_result, err := data_collection.InsertOne(ctx, data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Fatal(err)
		return
	}

	fmt.Println("Successfully inserted sensor data with _id: ", data_result.InsertedID)
	c.JSON(http.StatusOK, gin.H{"_id": data_result.InsertedID})
}
