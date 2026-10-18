package routes

import (
	"github.com/gin-gonic/gin"
	identityHandler "github.com/mohamed-karam/go-food-delivery/api-gateway-service/handler/identity"
	identityRoutes "github.com/mohamed-karam/go-food-delivery/api-gateway-service/routes/identity"
)

type Handlers struct {
	Auth    *identityHandler.AuthHandler
	Address *identityHandler.AddressHandler
}

func RegisterRoutes(r *gin.Engine, h Handlers) {
	api := r.Group("/api/v1")
	identity := api.Group("/identity")

	identityRoutes.AuthRoutes(identity, h.Auth)
	identityRoutes.AddressRoutes(identity, h.Address)
} 