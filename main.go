package main

import (
	"github.com/joho/godotenv"

	"vitorwdson/go-backend/models"
)

func main() {
	godotenv.Load()

	models.SetupDB()
}
