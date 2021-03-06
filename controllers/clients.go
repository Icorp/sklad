package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sklad/models"
)

func ListClients(c *gin.Context) {
	clientRepo := c.MustGet("clientRepo").(models.ClientRepo)
	clients, err := clientRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}

func GetClient(c *gin.Context) {
	clientRepo := c.MustGet("clientRepo").(models.ClientRepo)
	client, err := clientRepo.GetByID(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func UpdateClient(c *gin.Context) {
	clientRepo := c.MustGet("clientRepo").(models.ClientRepo)
	var data *models.Client
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id := c.Param("id")
	data.ID = id

	err := clientRepo.Update(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func DeleteClient(c *gin.Context) {
	clientRepo := c.MustGet("clientRepo").(models.ClientRepo)
	err := clientRepo.Delete(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func CreateClient(c *gin.Context) {
	clientRepo := c.MustGet("clientRepo").(models.ClientRepo)
	var data *models.Client
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	clients, err := clientRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	for i := range clients {
		if clients[i].Name == data.Name {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "product_category_already_exists",
			})
			return
		}
	}

	err = clientRepo.Create(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
