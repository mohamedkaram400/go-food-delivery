package auth

import (
	"errors"
	"fmt"
	"go/token"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
)


var jwtSecret = []byte(os.Getenv("SECRET_KEY"))


func generateToken(user *entity.User, duration time.Duration, tokenType string) (string, error) {

	claims := jwt.MapClaims{
		"sub": user.ID,
		"role": user.RoleID,
		"email": user.Email,
		"type": tokenType,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateAccessToken(user *entity.User, hours int) (string, error) {
	return generateToken(user, time.Duration(hours)*time.Hour, "AccessToken")
}

func GenerateRefreshToken(user *entity.User, days int) (string, error) {
	return generateToken(user, time.Duration(days)*24*time.Hour, "RefreshToken")
}

func ValidateRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			// Ensure the token uses the expected signing algo.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(RefreshTokenSecret), nil
		},
	)

	if err != nil {
		return "", fmt.Errorf("Invalid refresh token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", errors.New("Invalid refresh token")
	}

	// Ensure the token is intended for refreshing, not API access.
    if claims.TokenType != "refresh" {
        return "", errors.New("invalid token type")
    }

	return claims, nil
}