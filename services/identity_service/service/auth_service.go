package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mohamed-karam/go-food-delivery/identity-service/response"
	"github.com/mohamed-karam/go-food-delivery/identity-service/auth"
	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/pkg"
	"github.com/mohamed-karam/go-food-delivery/identity-service/repo"
	"github.com/mohamed-karam/go-food-delivery/identity-service/requests"
	"gorm.io/gorm"
)

type AuthService struct {
	identityRepo  *repo.IdentityRepo
	TokenDuration int
}

func NewAuthService(identityRepo *repo.IdentityRepo, tokenDuration int) *AuthService {
	return &AuthService{
		identityRepo:  identityRepo,
		TokenDuration: tokenDuration,
	}
}

func (s *AuthService) Register(ctx context.Context, req *requests.RegisterRequest) (string, *entity.User, error) {

	// Check validation errors
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return "", nil, err
	}

	var fieldErrors = map[string]string{}

	// Check email
	emailExisting, err := s.identityRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && emailExisting != nil {
		fieldErrors["email"] = "email already exists"
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, err
	}

	// Check phone
	phoneExisting, err := s.identityRepo.GetUserByPhone(ctx, req.Phone)
	if err == nil && phoneExisting != nil {
		fieldErrors["phone"] = "phone already exists"
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, err
	}

	if len(fieldErrors) > 0 {
		return "", nil, response.RegisterError{
			Fields: fieldErrors,
		}
	}

	// Hash password
	hashedPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return "", nil, err
	}

	// Prepare user object
	userObj := &entity.User{
		Name:         req.Name,
		Email:        req.Email,
		Phone:        req.Phone,
		RoleID:       req.RoleID,
		PasswordHash: hashedPassword,
	}

	// Pass the user object to repo for creation
	user, err := s.identityRepo.Register(ctx, userObj)
	if err != nil {
		return "", nil, err
	}

	log.Printf("User created: %+v", user)

	// Generate access token for that user
	accessToken, err := auth.GenerateAccessToken(user, time.Duration(s.TokenDuration)*time.Hour)
	if err != nil {
		return "", nil, err
	}

	return accessToken, user, nil
}

func (s *AuthService) Login(ctx context.Context, req *requests.LoginRequest) (*entity.User, error) {

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

func (s *AuthService) GetUser(ctx context.Context, userId int) (*entity.User, error) {

	return nil, nil
}

func (s *AuthService) Logout(ctx context.Context, userId int) (*entity.User, error) {
	return nil, nil

}

// func (s *AuthService) RefreshToken(ctx context.Context, req *requests.RefreshTokenRequest) (*requests.AuthResponse, error) {
// 	return nil, nil
// }

// func (s *AuthService) ValidateToken(ctx context.Context, req *requests.ValidateTokenRequest) (*requests.ValidateTokenResponse, error) {
// 	return nil, nil
// }
