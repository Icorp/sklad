package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sklad/models"
)

func ListOrders(c *gin.Context) {
	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	orders, err := orderRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	order, err := orderRepo.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func CreateOrder(c *gin.Context) {
	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	clientRepo := c.MustGet("clientRepo").(models.ClientRepo)
	productRepo := c.MustGet("productRepo").(models.ProductRepo)

	var data *models.Order
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check CLient Access
	client, err := clientRepo.GetByID(data.ClientID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    client,
			"error": err.Error(),
		})
		return
	}

	// Check Product
	product, err := productRepo.GetByID(data.ProductID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":    product,
			"error": err.Error(),
		})
		return
	}
	if product.Count <= data.Count {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "count_of_product_not_enouth",
		})
		return
	}

	err = productRepo.Update(&models.Product{
		ID:    data.ProductID,
		Count: product.Count - data.Count,
	})
	if err != nil {
		_ = err.Error()
		return
	}

	err = orderRepo.Create(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func UpdateOrder(c *gin.Context) {
	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	var data *models.Order
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")
	data.ID = id

	err := orderRepo.Update(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func DeleteOrder(c *gin.Context) {
	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	err := orderRepo.Delete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
