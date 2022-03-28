// controllers/products.go

package controllers

import (
	"bookCRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListProducts(c *gin.Context) {
	productRepo := c.MustGet("productRepo").(models.ProductRepo)
	products, err := productRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	productRepo := c.MustGet("productRepo").(models.ProductRepo)
	id := c.Params.ByName("id")
	product, err := productRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var data *models.Product
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	productRepo := c.MustGet("productRepo").(models.ProductRepo)
	err := productRepo.Create(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
