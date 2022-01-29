package config

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDB() *gorm.DB {
	host := "localhost"
	user := "admin"
	dbname := "admin"
	password := "admin-password"
	port := "5432"

	migratePath := "file://tool/db/migration"

	dsnForORM := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	dsnForMigrate := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	db, err := gorm.Open("postgres", dsnForORM)
	if err != nil {
		panic(err)
	}

	m, err := migrate.New(
		migratePath,
		dsnForMigrate,
	)
	if err != nil {
		panic(err)
	}
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}

	return db
}
