package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
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
	response, err := h.identityClient.Register(
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

		getError(err, c)
		return
	}


	log.Printf("✅ 7. User received: %+v", response)

	c.JSON(http.StatusCreated, gin.H{
		"data": response,
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

func getError(err error, c *gin.Context) {
	grpcStatus, ok := status.FromError(err)

	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	switch grpcStatus.Code() {
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, gin.H{
				"message": grpcStatus.Message(),
			})

		case codes.AlreadyExists:
			c.JSON(http.StatusConflict, gin.H{
				"message": grpcStatus.Message(),
			})

		case codes.Unauthenticated:
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": grpcStatus.Message(),
			})

		case codes.NotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"message": grpcStatus.Message(),
			})

		case codes.PermissionDenied:
			c.JSON(http.StatusForbidden, gin.H{
				"message": grpcStatus.Message(),
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "internal server error",
			})
	}
}