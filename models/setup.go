package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDB() {

	db, err := gorm.Open(
		postgres.Open(os.Getenv("PG_CONNECTION_STRING")),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	DB = db
}
