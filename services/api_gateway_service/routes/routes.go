package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/handler"
)


func AuthRoutes(rg *gin.RouterGroup, authHandler *handler.IdentityHandler) {
	auth := rg.Group("/auth")

	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
		auth.POST("/refresh-token", authHandler.RefreshToken)
	}
} 