package main

import (
	"github.com/gin-gonic/gin"
	"iot/controllers"
	"iot/models"
)

func main() {
	db := models.SetupModels()

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/", controllers.GetData)
	r.POST("/", controllers.PostData)
	r.GET("/devices", controllers.GetDevices)
	r.POST("/devices", controllers.MakeDevice)
	r.DELETE("/devices/:id", controllers.DeleteDevice)
	r.Run()
}

// curl -i -X GET -H "Content-Type: application/json" http://localhost:8080/devices
// curl -i -X POST -H "Content-Type: application/json" -d "{\"name\":\"RPi\"}" http://localhost:8080/devices
// curl -i -X DELETE http://localhost:8080/devices/60ce0bec8bd713427940abf7
