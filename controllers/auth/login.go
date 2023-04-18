package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vitorwdson/go-backend/application/auth"
	"vitorwdson/go-backend/db"
	"vitorwdson/go-backend/db/models"
)

func Login(c *gin.Context) {
	var input LoginInput

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	result := db.DB.Where("username = ?", input.Username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Username or password is invalid"})
		return
	}

	// TODO: Check against hashed password
	if user.Password != input.Password {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or password is invalid"})
		return
	}

	accessToken, refreshToken, err := auth.GenerateAuthenticationTokens(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("auth_token", accessToken, 60*10, "/", "", true, true)
	c.SetCookie("refresh_token", refreshToken, 60*60*24*7, "/", "", true, true)

	c.JSON(http.StatusOK, user)
}
