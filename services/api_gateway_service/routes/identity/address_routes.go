package identity

import (
	"github.com/gin-gonic/gin"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/handler/identity"
)


func AddressRoutes(rg *gin.RouterGroup, h *handler.AddressHandler) {
	address := rg.Group("/addresses")

	{
		address.POST("", h.AddAddresses)
	}
}  