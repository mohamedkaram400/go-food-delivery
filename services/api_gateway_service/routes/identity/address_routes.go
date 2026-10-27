package identity

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/handler/identity"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/middleware"
)


func AddressRoutes(rg *gin.RouterGroup, h *handler.AddressHandler, jwtSecret string) {
	address := rg.Group("/addresses")

	// All routes in this group require authentication.
    address.Use(middleware.AuthMiddleware([]byte(jwtSecret)))

	{
		address.POST("", h.AddAddresses)
	}
}  