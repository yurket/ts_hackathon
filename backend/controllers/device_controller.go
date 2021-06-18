package controllers

import (
	"iot/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDevices(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var devices []models.Device
	db.Find(&devices)

	c.JSON(http.StatusOK, gin.H{"data": devices})
}

func MakeDevice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var device models.Device
	if err := c.ShouldBindJSON(&device); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&device)
	c.JSON(http.StatusOK, gin.H{"data": device})
}

func DeleteDevice(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	id := c.Params.ByName("id")

	var device models.Device
	db.First(&device, id)

	if device.ID != 0 {
		db.Delete(&device)
		c.JSON(http.StatusOK, gin.H{"device": device})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Device not found"})
	}
}
