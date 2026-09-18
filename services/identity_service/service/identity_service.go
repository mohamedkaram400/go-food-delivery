package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mohamed-karam/go-food-delivery/identity-service/auth"
	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/pkg"
	"github.com/mohamed-karam/go-food-delivery/identity-service/repo"
	"github.com/mohamed-karam/go-food-delivery/identity-service/requests"
	"gorm.io/gorm"
)


type IdentityService struct {
	identityRepo *repo.IdentityRepo
	TokenDuration int
}

func NewIdentityService(identityRepo *repo.IdentityRepo, tokenDuration int) *IdentityService {
	return &IdentityService{
		identityRepo: identityRepo,
		TokenDuration: tokenDuration,
	}
}

func (s *IdentityService) Register(ctx context.Context, req *requests.RegisterRequest) (string, *entity.User, error) {

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return "", nil, err
	}

	existingUser, err := s.identityRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return "", nil, errors.New("email already exists")
	}

	if ! errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, err
	}

	hashedPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return "", nil, err
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
		return "", nil, err
	}

	log.Printf("User created: %+v", user)

	accessToken, err := auth.GenerateAccessToken(user, time.Duration(s.TokenDuration)*time.Hour)
	if err != nil {
		return "", nil, err
	}

	return accessToken, user, nil
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
