package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vitorwdson/go-backend/db"
	"vitorwdson/go-backend/db/models"
)

func RegisterUser(c *gin.Context) {
	var input RegisterUserInput

	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	err = user.ValidatePassword()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		if db.IsErrorCode(result.Error, db.PgDuplicateError) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already in use"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		}

		return
	}

	c.JSON(http.StatusCreated, user)
}
