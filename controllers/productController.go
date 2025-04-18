package controllers

import (
	"day23Assignment/config"
	"day23Assignment/models"
	"day23Assignment/utils"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := config.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&product)
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func GetProductImage(c *gin.Context) {
	productID := c.Param("image")

	var product models.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	imagePath := filepath.Join("uploads", "products", product.Image)

	_, err := os.Stat(imagePath)
	if product.Image == "" || (err != nil && os.IsNotExist(err)) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(imagePath)
}

func UpdateProductImage(c *gin.Context) {
	productID := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, productID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is empty"})
		return
	}

	const maxFileSize = 2 << 20 //const maxFileSize = 2 * 1024 * 1024
	if file.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File too large. Max size is 2MB"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only JPG, JPEG, and PNG files are allowed"})
		return
	}

	filename := utils.GenerateImageFilename(productID, product.Name)

	savePath := filepath.Join("uploads", "products", filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
		return
	}

	product.Image = filename
	if err := config.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product image in DB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Image uploaded successfully",
		"product": product,
	})
}
