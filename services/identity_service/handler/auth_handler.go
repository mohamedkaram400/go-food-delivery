package handler

import (
	"context"
	"log"

	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/requests"
	"github.com/mohamed-karam/go-food-delivery/identity-service/service"
)


type IdentityHandler struct {
	IdentityService *service.IdentityService
	pb.UnimplementedIdentityServiceServer
}

func NewIdentityHandler(IdentityService *service.IdentityService) *IdentityHandler {
	return &IdentityHandler{
		IdentityService: IdentityService,
	}
}

func (s *IdentityHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {

	log.Println("🔥 IDENTITY: Register gRPC handler reached")

	registerRequest := &requests.RegisterRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		RoleID:     int(req.RoleID),
		Phone:    &req.Phone,
	}
	log.Printf("📥 IDENTITY: Request received: %+v", registerRequest)

	log.Println(registerRequest)

	accessToken, user, err := s.IdentityService.Register(ctx, registerRequest)
	if err != nil {
		log.Printf("❌ IDENTITY: Service error: %v", err)
		return nil, err
	}

	log.Printf("✅ IDENTITY: User created: %+v", user)

	return &pb.AuthResponse{
		AccessToken: accessToken,
		User: &pb.User{
			Id:    int64(user.ID),
			Name:  user.Name,
			Email: user.Email,
			Phone: getStringValue(user.Phone),
		},
	}, nil
}

func (s *IdentityHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {


	// Call Identity Service using gRPC.
	// user, err := h.identityRepo.Login(
	// 	c.Request.Context(),
	// 	&pb.LoginRequest{
	// 		Email:    request.Email,
	// 		Password: request.Password,
	// 	},
	// )
	return nil, nil

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


func getStringValue(value *string) string {
	if value == nil {
		return ""
	}

	return *value
}