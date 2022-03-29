package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sklad/models"
	"net/http"
)

func ListEmployees(c *gin.Context) {
	employeeRepo := c.MustGet("employeeRepo").(models.EmployeeRepo)
	employees, err := employeeRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func GetEmployee(c *gin.Context) {
	employeeRepo := c.MustGet("employeeRepo").(models.EmployeeRepo)
	id := c.Params.ByName("id")
	employee, err := employeeRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func CreateEmployee(c *gin.Context) {
	employeeRepo := c.MustGet("employeeRepo").(models.EmployeeRepo)
	var data *models.Employee
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := employeeRepo.Create(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func UpdateEmployee(c *gin.Context) {
	employeeRepo := c.MustGet("employeeRepo").(models.EmployeeRepo)
	var data *models.Employee
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := employeeRepo.Update(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func DeleteEmployee(c *gin.Context) {
	employeeRepo := c.MustGet("employeeRepo").(models.EmployeeRepo)
	id := c.Params.ByName("id")
	err := employeeRepo.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
