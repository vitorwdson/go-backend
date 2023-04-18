package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"vitorwdson/go-backend/db/models"
)

func GenerateAuthenticationTokens(user *models.User) (string, string, error) {
	if user == nil || user.ID.String() == "" {
		return "", "", errors.New("User must be saved in the database before generating authentication tokens.")
	}

	accessSecretKey := []byte(os.Getenv("ACCESS_SECRET_KEY"))
	accessToken := jwt.New(jwt.SigningMethodHS512)

	accessTokenClaims := accessToken.Claims.(jwt.MapClaims)
	accessTokenClaims["exp"] = time.Now().Add(10 * time.Minute) // 10 minutes
	accessTokenClaims["userId"] = user.ID.String()

	accessTokenString, err := accessToken.SignedString(accessSecretKey)
	if err != nil {
		return "", "", errors.New("Unexpected error")
	}

	refreshToken := jwt.New(jwt.SigningMethodHS512)

	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["exp"] = time.Now().Add(24 * 7 * time.Hour) // One week
	refreshTokenClaims["userId"] = user.ID.String()

	refreshSecretKey := []byte(os.Getenv("REFRESH_SECRET_KEY"))
	refreshTokenString, err := refreshToken.SignedString(refreshSecretKey)
	if err != nil {
		return "", "", errors.New("Unexpected error")
	}

	return accessTokenString, refreshTokenString, nil
}
