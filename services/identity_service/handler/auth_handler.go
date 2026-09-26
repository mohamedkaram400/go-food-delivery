package handler

import (
	"context"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/requests"
	"github.com/mohamed-karam/go-food-delivery/identity-service/service"

	"github.com/mohamed-karam/go-food-delivery/identity-service/response"
	
    "google.golang.org/genproto/googleapis/rpc/errdetails"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)


type IdentityHandler struct { 
	AuthService *service.AuthService
	pb.UnimplementedIdentityServiceServer
}

func NewIdentityHandler(AuthService *service.AuthService) *IdentityHandler {
	return &IdentityHandler{
		AuthService: AuthService,
	}
}

func (s *IdentityHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {

	log.Println("🔥 IDENTITY: Register gRPC handler reached")

	// Prepare the register request
	registerRequest := &requests.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		RoleID:     int(req.RoleID),
		Phone:    req.Phone,
	}
	log.Printf("📥 IDENTITY: Request received: %+v", registerRequest)

	// Call the register endpoint to get the token and user
	accessToken, user, err := s.AuthService.Register(ctx, registerRequest)
	if err != nil {
		log.Printf("❌ IDENTITY: Service error: %v", err)

		return nil, toGRPCError(err)
	}

	log.Printf("✅ IDENTITY: User created: %+v", user)

	// Return the auth response 
	return &pb.AuthResponse{
		RefreshToken: accessToken,
		AccessToken: accessToken,
		User: &pb.User{
			Id:    int64(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
	}, nil
}

func (s *IdentityHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {

	log.Println("🔥 IDENTITY: Login gRPC handler reached")

	// Prepare the request to login request
	loginRequest := &requests.LoginRequest{
		Email: req.Email,
		Password: req.Password,
	}

	log.Printf("📥 IDENTITY: Request received: %+v", loginRequest)

	// Call the login endpoint to get the token and user
	accessToken, user, err := s.AuthService.Login(ctx, loginRequest)
	if err != nil {
		log.Printf("❌ IDENTITY: Service error: %v", err)

		return nil, toGRPCError(err)
	}

	log.Printf("✅ IDENTITY: User created: %+v", user)

	// Return the auth response 
	return &pb.AuthResponse{
		AccessToken: accessToken,
		RefreshToken: accessToken,
		User: &pb.User{
			Name: user.Name,
			Email: user.Email,
			Phone: user.Phone,
		},
	}, nil
}

func (s *IdentityHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
	return nil, nil
}

func (s *IdentityHandler) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.Empty, error) {
	return nil, nil
}

func (s *IdentityHandler) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.AuthResponse, error) {
	return nil, nil
}

func (s *IdentityHandler) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {
	return nil, nil
}


func toGRPCError(err error) (error) {

	var registerErr response.RegisterError

	if errors.As(err, &registerErr) {

		st := status.New(
			codes.AlreadyExists,
			"registration failed",
		)

		var violations []*errdetails.BadRequest_FieldViolation

		for field, message := range registerErr.Fields {
			violations = append(
				violations,
				&errdetails.BadRequest_FieldViolation{
					Field:       field,
					Description: message,
				},
			)
		}

		detailedStatus, detailErr := st.WithDetails(
			&errdetails.BadRequest{
				FieldViolations: violations,
			},
		)

		if detailErr != nil {
			return status.Errorf(
				codes.Internal,
				"failed to create error details: %v",
                detailErr,
			)
		}

		return detailedStatus.Err()
	}

	// Invalid credentials
	var creadentialErr response.InvalidCredentialsError

	if errors.As(err, &creadentialErr) {
		return status.Errorf(
			codes.Unauthenticated,
			creadentialErr.Error(),
		)
	}

	// Validation error
	var validationErr validator.ValidationErrors

	if errors.As(err, &validationErr) {
		return status.Errorf(
			codes.InvalidArgument,
            "validation failed: %v",
            validationErr,
		)
	}

	// Unexpected error
	return status.Errorf(
		codes.Internal,
		"internal error: %v",
        err,
	)
}