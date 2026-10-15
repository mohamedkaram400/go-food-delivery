package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/errors"
	res "github.com/mohamed-karam/go-food-delivery/api-gateway-service/response"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
)


type AddressHandler struct {
	addressClient pb.AddressServiceClient
}

func NewAddressHandler(addressClient pb.AddressServiceClient) *AddressHandler {
	return &AddressHandler{
		addressClient: addressClient,
	}
}

func (h *AddressHandler) AddAddress(c *gin.Context) {

	var request struct {
        Addresses []string `json:"addresses" binding:"required,min=1,max=5,dive,required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.WriteGRPCError(err, c)
		return 
	}

	addresses, err := h.addressClient.CreateAddress(
		c.Request.Context(),
		&pb.CreateAddressesRequest{
			Addresses: request.Addresses,
		},
	)

	if err != nil {
		errors.WriteGRPCError(err, c)
		return 
	}

	c.JSON(http.StatusCreated, res.APIResponse{
		Success: true,
		Message: "Address Added successful",
		Data: addresses,
	})
}