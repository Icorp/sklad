package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sklad/models"
	"net/http"
	"time"
)

func GetSkladInfo(c *gin.Context) {
	type orderResponse struct {
		ID              string                  `json:"id"`
		Product         *models.Product         `json:"product"`
		ProductCategory *models.ProductCategory `json:"product_category"`
		Employee        *models.Employee        `json:"employee"`
		CreatedAt       string                  `json:"created_at"`
		Address         string                  `json:"address"`
		Count           int                     `json:"count"`
	}
	type response struct {
		Profit int             `json:"profit"`
		Orders []orderResponse `json:"orders"`
	}

	const RFC3339FullDate = "2006-01-02"
	fromQuery := c.Query("from")
	fromDate := time.Now().AddDate(0, -1, 0)
	if fromQuery != "" {
		var err error
		fromDate, err = time.Parse(RFC3339FullDate, c.Query("from"))
		if err != nil {
			_ = c.Error(err)
			return
		}
	}

	toQuery := c.Query("to")
	toDate := time.Now()
	if toQuery != "" {
		var err error
		toDate, err = time.Parse(RFC3339FullDate, c.Query("to"))
		if err != nil {
			_ = c.Error(err)
			return
		}
	}

	orderRepo := c.MustGet("orderRepo").(models.OrderRepo)
	productCategoryRepo := c.MustGet("productCategoryRepo").(models.ProductCategoryRepo)
	productRepo := c.MustGet("productRepo").(models.ProductRepo)
	employeeRepo := c.MustGet("employeeRepo").(models.EmployeeRepo)
	orders, err := orderRepo.GetAll(&models.OrderRequest{
		From: &fromDate,
		To:   &toDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	profit := 0
	for i := range orders {
		profit += orders[i].CurrentPrice * orders[i].Count
	}

	productCategoryList, err := productCategoryRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	orderResponses := make([]orderResponse, len(orders))
	for i := range orders {
		var productCategory *models.ProductCategory

		product, err := productRepo.GetByID(orders[i].ProductID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		employee, err := employeeRepo.GetByID(orders[i].EmployeeID)
		if err != nil {
			_ = c.Error(err)
			return
		}

		for j := range productCategoryList {
			if productCategoryList[j].ID == product.ProductCategoryID {
				productCategory = productCategoryList[j]
			}
		}

		orderResponses[i].ID = orders[i].ID
		orderResponses[i].Product = product
		orderResponses[i].ProductCategory = productCategory
		orderResponses[i].Employee = employee
		orderResponses[i].Address = orders[i].Address
		orderResponses[i].Count = orders[i].Count
	}

	c.JSON(http.StatusOK, &response{
		Profit: profit,
		Orders: orderResponses,
	})
}
