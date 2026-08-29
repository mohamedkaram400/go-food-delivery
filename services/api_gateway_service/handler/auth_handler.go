package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
)


type AuthHandler struct {
	identityClient pb.IdentityServiceClient
}

func NewAuthHandler(identityClient pb.IdentityServiceClient) *AuthHandler {
	return &AuthHandler{
		identityClient: identityClient,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request struct {
		Name    string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	// Call Identity Service using gRPC.
	user, err := h.identityClient.Register(
		c.Request.Context(),
		&pb.RegisterRequest{
			Name:    request.Name,
			Email:    request.Email,
			Password: request.Password,
		},
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}


func (h *AuthHandler) Login(c *gin.Context) {

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	// Call Identity Service using gRPC.
	user, err := h.identityClient.Login(
		c.Request.Context(),
		&pb.LoginRequest{
			Email:    request.Email,
			Password: request.Password,
		},
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}