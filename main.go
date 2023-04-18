package main

import (
	"github.com/joho/godotenv"

	"vitorwdson/go-backend/db"
)

func main() {
	godotenv.Load()

	db.SetupDB()
}
