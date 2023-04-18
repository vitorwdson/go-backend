package db

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"vitorwdson/go-backend/db/models"
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

	db.AutoMigrate(&models.User{})

	DB = db
}
