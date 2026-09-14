package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
)


var jwtSecret = []byte(os.Getenv("SECRET_KEY"))

func GenerateAccessToken(user *entity.User, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"role": user.RoleID,
		"email": user.Email,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(duration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
