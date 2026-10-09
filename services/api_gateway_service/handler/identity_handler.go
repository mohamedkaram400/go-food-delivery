package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/errors"
	res "github.com/mohamed-karam/go-food-delivery/api-gateway-service/response"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
)


type IdentityHandler struct {
	identityClient pb.IdentityServiceClient
}

func NewIdentityHandler(identityClient pb.IdentityServiceClient) *IdentityHandler {
	return &IdentityHandler{
		identityClient: identityClient,
	}
}

func (h *IdentityHandler) Register(c *gin.Context) {
	var request struct {
		Name    string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone string `json:"phone"`
		RoleID  int32 `json:"role_id"`
	}

	log.Println("🔥 1. Register HTTP handler started")

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("❌ 2. JSON binding error: %v", err)

		errors.WriteGRPCError(err, c)
		return
	}

	log.Println("✅ 3. JSON binding successful")
	log.Printf("Name: %v", request.Name)
	log.Printf("Email: %v", request.Email)

	// Call Identity Service using gRPC.
	data, err := h.identityClient.Register(
		c.Request.Context(),
		&pb.RegisterRequest{
			Name:    request.Name,
			Email:    request.Email,
			Password: request.Password,
			Phone: request.Phone,
			RoleID: request.RoleID,
		},
	)

	log.Println("📥 5. Returned from Identity Service gRPC call")

	if err != nil {
		log.Printf("❌ 6. gRPC error: %v", err)

		errors.WriteGRPCError(err, c)
		return
	}

	log.Printf("✅ 7. User received: %+v", data)

	c.JSON(http.StatusCreated, res.APIResponse{
		Success: true,
		Message: "Registration successful",
		Data: data,
	})
}

func (h *IdentityHandler) Login(c *gin.Context) {

	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.WriteGRPCError(err, c)
		return
	}

	// Call Identity Service using gRPC.
	response, err := h.identityClient.Login(
		c.Request.Context(),
		&pb.LoginRequest{
			Email:    request.Email,
			Password: request.Password,
		},
	)

	if err != nil {
		errors.WriteGRPCError(err, c)
		return
	}

	c.JSON(http.StatusOK, res.APIResponse{
		Success: true,
		Message: "Login successful",
		Data: response,
	})
}

func (h *IdentityHandler) Logout(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.WriteGRPCError(err, c)
		return
	}

	// Call Identity Service using gRPC.
	_, err := h.identityClient.Logout(
		c.Request.Context(),
		&pb.LogoutRequest{
			RefreshToken:    request.RefreshToken,
		},
	)

	if err != nil {
		errors.WriteGRPCError(err, c)
		return
	}

	c.JSON(http.StatusOK, res.APIResponse{
		Success: true,
		Message: "Logout successful",
	})
}

func (h *IdentityHandler) RefreshToken(c *gin.Context) {

	var request struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		errors.WriteGRPCError(err, c)
		return 
	}

	// Refresh token enpoint through gRPC
	accessToken, err := h.identityClient.RefreshToken(
		c.Request.Context(),
		&pb.RefreshTokenRequest{
			RefreshToken: request.RefreshToken,
		},
	)

	if err != nil {
		errors.WriteGRPCError(err, c)
		return 
	}

	c.JSON(http.StatusCreated, res.APIResponse{
		Success: true,
		Message: "Access token generated successful",
		Data: accessToken,
	})
}