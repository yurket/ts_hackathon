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
	r.GET("/", controllers.GetDevices)
	r.POST("/", controllers.MakeDevice)
	r.DELETE("/:id", controllers.DeleteDevice)
	r.Run()
}

// curl -i -X GET -H "Content-Type: application/json" http://localhost:8080/
// curl -i -X POST -H "Content-Type: application/json" -d "{\"name\":\"RPi\"}" http://localhost:8080/
// curl -i -X DELETE http://localhost:8080/52
