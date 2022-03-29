package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sklad/models"
)

func ListProductCategories(c *gin.Context) {
	productCategoryRepo := c.MustGet("productCategoryRepo").(models.ProductCategoryRepo)

	listProductCategory, err := productCategoryRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, listProductCategory)
}

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

func GetProductCategory(c *gin.Context) {
	productCategoryRepo := c.MustGet("productCategoryRepo").(models.ProductCategoryRepo)

	id := c.Param("id")
	productCategory, err := productCategoryRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, productCategory)
}

func UpdateProductCategory(c *gin.Context) {
	var data *models.ProductCategory
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productCategoryRepo := c.MustGet("productCategoryRepo").(models.ProductCategoryRepo)
	err := productCategoryRepo.Update(data)
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

func DeleteProductCategory(c *gin.Context) {
	productCategoryRepo := c.MustGet("productCategoryRepo").(models.ProductCategoryRepo)

	id := c.Param("id")
	err := productCategoryRepo.Delete(id)
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
