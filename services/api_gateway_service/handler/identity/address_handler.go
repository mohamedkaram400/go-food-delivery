package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/errors"
	res "github.com/mohamed-karam/go-food-delivery/api-gateway-service/response"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"google.golang.org/grpc/metadata"
)


type AddressHandler struct {
	addressClient pb.AddressServiceClient
}

func NewAddressHandler(addressClient pb.AddressServiceClient) *AddressHandler {
	return &AddressHandler{
		addressClient: addressClient,
	}
}

func (h *AddressHandler) AddAddresses(c *gin.Context) {

	var request struct { 
        Addresses []string `json:"addresses" binding:"required,min=1,max=5,dive,required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.WriteGRPCError(err, c)
		return 
	}

	log.Println("GATEWAY: Request comes to address successfully")

	// Set by authentication middleware after JWT verification.
	userID := c.GetInt64("user_id")
	if userID <= 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
            "message": "Unauthorized",
        })
        return
	}

	ctx := metadata.AppendToOutgoingContext(
		c.Request.Context(),
		"x-user-id",
		strconv.FormatInt(userID, 10),
	)

	addresses, err := h.addressClient.CreateAddresses(
		ctx,
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