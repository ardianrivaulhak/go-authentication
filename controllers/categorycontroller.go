package controllers

import (
	"go-authentication/database"
	"go-authentication/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var category models.Categories

	if err := c.ShouldBindJSON(&category); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	return

	}

	if err := database.Instance.Create(&category).Error; err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return

	}

	c.JSON(http.StatusOK, category)
}


func GetCategories(c *gin.Context) {
	var categories []models.Categories

	database.Instance.Find(&categories)

	c.JSON(http.StatusOK, categories)
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")

	var category models.Categories

	if err := database.Instance.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Categories not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func UpdateCategory(c *gin.Context)  {
	type UpdateCategoryInput struct {
		Name string `json:"name"`
	}
	var category models.Categories

	if err := database.Instance.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found!"})
		return
	}

	// Validate input
	var input UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := database.Instance.Model(&category).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DestroyCategort(c *gin.Context)  {
	var category models.Categories

	if err := database.Instance.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found!"})
		return
	}

	database.Instance.Model(&category).Delete(&category)
	
	c.JSON(http.StatusOK, gin.H{"data": true})
}