package auth

import (
	"errors"
	"fmt"
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

func ValidateRefreshToken(refreshToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(
		refreshToken,
		func(t *jwt.Token) (interface{}, error) {
			// Ensure the token uses the expected signing algo.
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(
					"unexpected signing method: %v", 
					t.Header["alg"],
				)
			}

			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Invalid refresh token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid refresh token")
	}

	// Ensure the token is intended for refreshing, not API access.
    if claims["type"] != "RefreshToken" {
        return nil, errors.New("invalid token type")
    }

	return claims, nil
}
