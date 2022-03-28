package controllers

import (
	"bookCRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProductCategory(c *gin.Context) {
	var data *models.ProductCategory
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productCategoryRepo := c.MustGet("productCategoryRepo").(models.ProductCategoryRepo)
	err := productCategoryRepo.Create(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
