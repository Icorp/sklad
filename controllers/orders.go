package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sklad/models"
	"net/http"
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
	order := models.Order{}
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = orderRepo.Create(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func UpdateOrder(c *gin.Context) {
	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	order := models.Order{}
	err := c.BindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = orderRepo.Update(&order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	err := orderRepo.Delete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "deleted"})
}
