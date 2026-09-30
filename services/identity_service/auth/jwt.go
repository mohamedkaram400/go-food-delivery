package auth

import (
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
