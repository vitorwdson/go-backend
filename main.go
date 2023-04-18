package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"vitorwdson/go-backend/controllers/auth"
	"vitorwdson/go-backend/db"
)

func main() {
	godotenv.Load()

	db.SetupDB()

	r := gin.Default()

	r.POST("auth/register", auth.RegisterUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
