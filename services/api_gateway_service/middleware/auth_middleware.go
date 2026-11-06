package middleware

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtSecret []byte) gin.HandlerFunc {
    return func(c *gin.Context) {
        authorization := c.GetHeader("Authorization")

        if !strings.HasPrefix(authorization, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "message": "Missing or invalid authorization header",
            })
            return 
        }

        tokenString := strings.TrimSpace(strings.TrimPrefix(authorization, "Bearer "))
        if tokenString == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "message": "Missing access token",
            })
            return
        }

        claims := jwt.MapClaims{}

        token, err := jwt.ParseWithClaims(
            tokenString,
            claims,
            func(token *jwt.Token) (interface{}, error) {
                return jwtSecret, nil
            },
            jwt.WithValidMethods([]string{"HS256"}),
            jwt.WithExpirationRequired(),
        )

        if err != nil {
            log.Printf("JWT validation failed: %v", err)

            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "message": "Invalid or expired access token",
            })
            return 
        }

        if !token.Valid {
            log.Println("JWT token is invalid")
            c.AbortWithStatus(http.StatusUnauthorized)
            return 
        }

        tokenType, ok := claims["type"].(string)
        if !ok || tokenType != "AccessToken" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "message": "Access token required",
            })
            return
        }

        // Your existing JWT uses "sub" to store the user ID.
        subject, err := claims.GetSubject()
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        userID, err := strconv.ParseInt(subject, 10, 64)
        if err != nil || userID <= 0 {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        c.Set("user_id", userID)
        c.Next()
    }
}
