package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sklad/models"
)

func ListOrders(c *gin.Context) {
	type response struct {
		ID              string                  `json:"id"`
		Product         *models.Product         `json:"product"`
		ProductCategory *models.ProductCategory `json:"product_category"`
		Employee        *models.Employee        `json:"employee"`
		CreatedAt       string                  `json:"created_at"`
		Address         string                  `json:"address"`
		Count           int                     `json:"count"`
	}

	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	productCategoryRepo := c.MustGet("productCategoryRepo").(models.ProductCategoryRepo)
	orders, err := orderRepo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	productCategoryList, err := productCategoryRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ordersResponse := make([]response, len(orders))
	for i := range orders {
		var productCategory *models.ProductCategory
		for j := range productCategoryList {
			if productCategoryList[j].ID == orders[i].Product.ProductCategoryID {
				productCategory = productCategoryList[j]
			}
		}

		ordersResponse[i].ID = orders[i].ID
		ordersResponse[i].Product = orders[i].Product
		ordersResponse[i].ProductCategory = productCategory
		ordersResponse[i].Employee = orders[i].Employee
		ordersResponse[i].Address = orders[i].Address
		ordersResponse[i].Count = orders[i].Count
	}

	c.JSON(http.StatusOK, ordersResponse)
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

	data.CurrentPrice = product.Price

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
