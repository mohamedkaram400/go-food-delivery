package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	log.Println("✅ 3. JSON binding successful")
	log.Printf("Name: %v", request.Name)
	log.Printf("Email: %v", request.Email)

	// Call Identity Service using gRPC.
	user, err := h.identityClient.Register(
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
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
		})
		return
	}

	log.Printf("✅ 7. User received: %+v", user)

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}


func (h *IdentityHandler) Login(c *gin.Context) {

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