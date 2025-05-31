package handlers

import (
	"net/http"
	"rebac/db"
	"rebac/models"
	"github.com/gin-gonic/gin"
)

// CreateModel creates a new Model
func CreateModel(c *gin.Context) {
	var model models.Model
	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.DB.Create(&model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model)
}

// GetModels returns all Models
func GetModels(c *gin.Context) {
	var modelsList []models.Model
	if err := db.DB.Find(&modelsList).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, modelsList)
}

// GetModel returns a Model by ID
func GetModel(c *gin.Context) {
	var model models.Model
	id := c.Param("id")
	if err := db.DB.First(&model, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	c.JSON(http.StatusOK, model)
}

// UpdateModel updates a Model by ID
func UpdateModel(c *gin.Context) {
	var model models.Model
	id := c.Param("id")
	if err := db.DB.First(&model, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Model not found"})
		return
	}
	var input models.Model
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model.Name = input.Name
	model.Description = input.Description
	if err := db.DB.Save(&model).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, model)
}

// DeleteModel deletes a Model by ID
func DeleteModel(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Model{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Model deleted"})
}
