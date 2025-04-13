package controllers

import (
	"day23Assignment/config"
	"day23Assignment/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStockByID(c *gin.Context) {
	id := c.Param("id")
	var inventory models.Inventory
	if err := config.DB.Where("id = ?", id).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, inventory)
}

func UpdateStock(c *gin.Context) {
	id := c.Param("id")
	var inventory models.Inventory
	if err := config.DB.Where("id = ?", id).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&inventory)
	c.JSON(http.StatusOK, inventory)
}
