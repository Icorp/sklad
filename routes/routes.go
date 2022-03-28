package routes

import (
	"bookCRUD/controllers"
	"bookCRUD/repository"
	"github.com/go-pg/pg/v10"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *pg.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("clientRepo", repository.NewClientRepo(db))
		c.Set("productRepo", repository.NewProductRepo(db))
		c.Set("orderRepo", repository.NewOrderRepo(db))
		c.Set("productCategoryRepo", repository.NewProductCategoryRepo(db))
		c.Set("employeeRepo", repository.NewEmployeeRepo(db))
		c.Set("providerRepo", repository.NewProviderRepo(db))
	})

	r.GET("/products", controllers.ListProducts)
	r.GET("/products/:id", controllers.GetProduct)
	r.POST("/products", controllers.CreateProduct)

	r.POST("/product_categories", controllers.CreateProductCategory)
	return r
}
