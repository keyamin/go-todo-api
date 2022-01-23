package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=admin dbname=admin password=admin-password sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
