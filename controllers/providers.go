package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sklad/models"
	"net/http"
)

func CreateProvider(c *gin.Context) {
	providerRepo := c.MustGet("providerRepo").(models.ProviderRepo)
	var data *models.Provider
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := providerRepo.Create(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func ListProviders(c *gin.Context) {
	providerRepo := c.MustGet("providerRepo").(models.ProviderRepo)
	providers, err := providerRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   providers,
	})
}

func GetProvider(c *gin.Context) {
	providerRepo := c.MustGet("providerRepo").(models.ProviderRepo)
	id := c.Param("id")
	provider, err := providerRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   provider,
	})
}

func UpdateProvider(c *gin.Context) {
	providerRepo := c.MustGet("providerRepo").(models.ProviderRepo)
	var data *models.Provider
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := providerRepo.Update(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func DeleteProvider(c *gin.Context) {
	providerRepo := c.MustGet("providerRepo").(models.ProviderRepo)
	id := c.Param("id")
	err := providerRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
