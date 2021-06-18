package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Sensor struct {
	Time        string `json:"Time"`
	Temperature string `json:"Temperature"`
	Pressure    string `json:"Pressure"`
	Humidity    string `json:"Humidity"`
}

var sensor_data = []Sensor{}

func GetData(c *gin.Context) {
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
}
