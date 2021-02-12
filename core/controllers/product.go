package controllers

import (
	"net/http"

	"github.com/YannaAlghamdi/xenelectronic/core/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// GET /products
// Get all products
func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Product
	db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// POST /products
// Create new products
func Create(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Validate input
	var input models.CreateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create product
	product := models.Product{Name: input.Name, Description: input.Description}
	db.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// GET /products/:id
// Find a product
func Get(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// PATCH /products/:id
// Update a product
func Update(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	// Validate input
	var input models.UpdateProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Model(&product).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// DELETE /products/:id
// Delete a product
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var product models.Product
	if err := db.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
