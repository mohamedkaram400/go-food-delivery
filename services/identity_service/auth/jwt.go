package auth

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
)


func generateToken(user *entity.User, duration time.Duration, tokenType string, jwtSecret []byte) (string, error) {

	claims := jwt.MapClaims{
		"sub": strconv.FormatInt(int64(user.ID), 10),
		"role": user.RoleID,
		"email": user.Email,
		"type": tokenType,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateAccessToken(user *entity.User, duration time.Duration, jwtSecret []byte) (string, error) {
	return generateToken(user, duration, "AccessToken", jwtSecret)
}

func GenerateRefreshToken(user *entity.User, duration time.Duration, jwtSecret []byte) (string, error) {
	return generateToken(user, duration, "RefreshToken", jwtSecret)
}

func ValidateRefreshToken(refreshToken string, jwtSecret string) (jwt.MapClaims, error) {
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
