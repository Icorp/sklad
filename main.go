package main

import (
	"bookCRUD/models"
	"bookCRUD/routes"
)

func main() {

	db := models.SetupDB()
	r := routes.SetupRoutes(db)
	r.Run()
}
