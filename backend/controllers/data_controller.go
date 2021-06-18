package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Sensor struct {
	Time        string `json:"Time"`
	Temperature string `json:"Temperature"`
	Pressure    string `json:"Pressure"`
	Humidity    string `json:"Humidity"`
}

var sensor_data = []Sensor{}

func GetData(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	c.JSON(http.StatusOK, gin.H{"data": sensor_data})
}

func PostData(c *gin.Context) {
	fmt.Println("Header:", c.Request.Header)
	fmt.Println("Host:", c.Request.Host)
	fmt.Println("Body:", c.Request.Body)
	var data Sensor
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sensor_data = append(sensor_data, data)
	c.JSON(http.StatusOK, gin.H{})
}
