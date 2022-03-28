package models

import (
	"github.com/go-pg/pg/v10"
)

// SetupDB : initializing postgresql database
func SetupDB() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "admin",
		Database: "sklad",
	})
	return db
}
