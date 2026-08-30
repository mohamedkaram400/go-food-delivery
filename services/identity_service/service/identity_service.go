package service

import (
	"context"

	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/pkg"
	"github.com/mohamed-karam/go-food-delivery/identity-service/repo"
	"github.com/mohamed-karam/go-food-delivery/identity-service/requests"
)


type IdentityService struct {
	identityRepo *repo.IdentityRepo
}

func NewIdentityService(identityRepo *repo.IdentityRepo) *IdentityService {
	return &IdentityService{
		identityRepo: identityRepo,
	}
}

func (s *IdentityService) Register(ctx context.Context, req *requests.RegisterRequest) (*entity.User, error) {

	hashedPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	userObj := &entity.User{
		Name: req.Name,
		Email: req.Email,
		Phone: req.Phone,
		RoleID: req.RoleID,
		PasswordHash: hashedPassword,
	}

	user, err := s.identityRepo.Register(ctx, userObj)
	
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *IdentityService) Login(ctx context.Context, req *requests.LoginRequest) (*entity.User, error) {

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


func (s *IdentityService) GetUser(ctx context.Context, userId int) (*entity.User, error) {

	return nil, nil
}


func (s *IdentityService) Logout(ctx context.Context, userId int) (*entity.User, error) {
	return nil, nil

}

// func (s *IdentityService) RefreshToken(ctx context.Context, req *requests.RefreshTokenRequest) (*requests.AuthResponse, error) {
// 	return nil, nil
// }

// func (s *IdentityService) ValidateToken(ctx context.Context, req *requests.ValidateTokenRequest) (*requests.ValidateTokenResponse, error) {
// 	return nil, nil
// }
