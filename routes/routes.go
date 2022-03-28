package routes

import (
	"github.com/go-pg/pg/v10"
	"github.com/sklad/controllers"
	"github.com/sklad/repository"

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
	r.POST("/products/:id", controllers.UpdateProduct)
	r.POST("/products", controllers.CreateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)

	r.GET("/product_categories", controllers.ListProductCategories)
	r.GET("/product_categories/:id", controllers.GetProductCategory)
	r.POST("/product_categories/:id", controllers.UpdateProductCategory)
	r.POST("/product_categories", controllers.CreateProductCategory)
	r.DELETE("/product_categories/:id", controllers.DeleteProductCategory)

	r.GET("/clients", controllers.ListClients)
	r.GET("/clients/:id", controllers.GetClient)
	r.POST("/clients/:id", controllers.UpdateClient)
	r.POST("/clients", controllers.CreateClient)
	r.DELETE("/clients/:id", controllers.DeleteClient)

	r.GET("/orders", controllers.ListOrders)
	r.GET("/orders/:id", controllers.GetOrder)
	r.POST("/orders/:id", controllers.UpdateOrder)
	r.POST("/orders", controllers.CreateOrder)
	r.DELETE("/orders/:id", controllers.DeleteOrder)

	r.GET("/employees", controllers.ListEmployees)
	r.GET("/employees/:id", controllers.GetEmployee)
	r.POST("/employees/:id", controllers.UpdateEmployee)
	r.POST("/employees", controllers.CreateEmployee)
	r.DELETE("/employees/:id", controllers.DeleteEmployee)

	r.GET("/providers", controllers.ListProviders)
	r.GET("/providers/:id", controllers.GetProvider)
	r.POST("/providers/:id", controllers.UpdateProvider)
	r.POST("/providers", controllers.CreateProvider)
	r.DELETE("/providers/:id", controllers.DeleteProvider)
	return r
}
