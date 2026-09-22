package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	res "github.com/mohamed-karam/go-food-delivery/api-gateway-service/response"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

		getError(err, c)
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

		getError(err, c)
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
		c.JSON(http.StatusInternalServerError, res.APIResponse{
            Success: false,
            Message: "Internal server error",
        })
		return
	}

	switch grpcStatus.Code() {
		case codes.InvalidArgument:
			c.JSON(http.StatusBadRequest, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
				Errors: extractFieldErrors(grpcStatus),
			})

		case codes.AlreadyExists:
			c.JSON(http.StatusConflict, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
				Errors: extractFieldErrors(grpcStatus),
			})

		case codes.Unauthenticated:
			c.JSON(http.StatusUnauthorized, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
			})

		case codes.NotFound:
			c.JSON(http.StatusNotFound, res.APIResponse{
            	Success: false,
				Message: grpcStatus.Message(),
			})

		case codes.PermissionDenied:
			c.JSON(http.StatusForbidden, res.APIResponse{
				Success: false,
				Message: grpcStatus.Message(),
			})

		default:
			c.JSON(http.StatusInternalServerError, res.APIResponse{
				Success: false,
				Message: "Internal server error",
			})
	}
}

func extractFieldErrors(
    grpcStatus *status.Status,
) map[string]string {

    errorsMap := make(map[string]string)

    for _, detail := range grpcStatus.Details() {

        badRequest, ok := detail.(*errdetails.BadRequest)

        if !ok {
            continue
        }

        for _, violation := range badRequest.FieldViolations {
            errorsMap[violation.Field] = violation.Description
        }
    }

    return errorsMap
}