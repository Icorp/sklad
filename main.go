package main

import (
	"github.com/sklad/models"
	"github.com/sklad/routes"
)

func main() {

	db := models.SetupDB()
	r := routes.SetupRoutes(db)
	r.Run()
}
