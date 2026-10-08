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
	"github.com/mohamed-karam/go-food-delivery/identity-service/response"
	"gorm.io/gorm"
)

type AuthService struct {
	identityRepo  *repo.IdentityRepo
	AccessTokenDuration int
	RefreshTokenDuration int
}

func NewAuthService(identityRepo *repo.IdentityRepo, accessTokenDuration int, refreshTokenDuration int) *AuthService {
	return &AuthService{
		identityRepo:  identityRepo,
		AccessTokenDuration: accessTokenDuration,
		RefreshTokenDuration: refreshTokenDuration,
	}
}

func (s *AuthService) Register(ctx context.Context, req *requests.RegisterRequest) (string, string,  *entity.User, error) {

	// Check validation errors
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return "", "", nil, err
	}

	var fieldErrors = map[string]string{}

	// Check email
	emailExisting, err := s.identityRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && emailExisting != nil {
		fieldErrors["email"] = "email already exists"
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", nil, err
	}

	// Check phone
	phoneExisting, err := s.identityRepo.GetUserByPhone(ctx, req.Phone)
	if err == nil && phoneExisting != nil {
		fieldErrors["phone"] = "phone already exists"
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", "", nil, err
	}

	if len(fieldErrors) > 0 {
		return "", "", nil, response.RegisterError{
			Fields: fieldErrors,
		}
	}

	// Hash password
	hashedPassword, err := pkg.HashPassword(req.Password)
	if err != nil {
		return "", "", nil, err
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
		return "", "", nil, err
	}

	log.Printf("User created: %+v", user)

	// Generate access token for that user
	accessToken, err := auth.GenerateAccessToken(user, s.AccessTokenDuration)
	if err != nil {
		return "", "", nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken(user, s.RefreshTokenDuration)
	if err != nil {
		return "", "", nil, err
	}

	if err := s.storeRefreshToken(ctx, refreshToken, user); err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, user, nil
}

func (s *AuthService) Login(ctx context.Context, req *requests.LoginRequest) (string, string, *entity.User, error) {

	// Make validation in the request
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return "", "", nil, err
	}

	// Check if user/email exists or not 
	exists, err := s.identityRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", "", nil, response.InvalidCredentialsError{
				Message: "Invalid email or password",
			}
		}

		return "", "", nil, err
	}

	// Check from password
	if err := pkg.CheckPassword(req.Password, exists.PasswordHash); err != nil {
		return "", "", nil, response.InvalidCredentialsError{
			Message: "The sended password doesn't match the user password",
		}
	}

	accessToken, err := auth.GenerateAccessToken(exists, s.AccessTokenDuration)
	if err != nil {
		return "", "", nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken(exists, s.RefreshTokenDuration)
	if err != nil {
		return "", "", nil, err
	}

	if err := s.storeRefreshToken(ctx, refreshToken, exists); err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, exists, nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken string) error {

	hashedToken := pkg.HashToken(refreshToken)

	if err := s.identityRepo.RevokeRefreshToken(ctx, hashedToken); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	// Validate the refresh token and extract its claims.
    _, err := auth.ValidateRefreshToken(refreshToken)
    if err != nil {
        return "", err
    }

	// Hash the refresh token.
    hashedToken := pkg.HashToken(refreshToken)

	// Find the stored refresh token.
    storedToken, err := s.identityRepo.GetRefreshToken(ctx, hashedToken)
    if err != nil {
        return "", err
    }

	// Verify that the refresh token is still valid in the database.
    if storedToken.RevokedAt != nil {
        return "", errors.New("refresh token has been revoked")
    }

    if !storedToken.ExpiresAt.After(time.Now()) {
        return "", errors.New("refresh token has expired")
    }

	// Get the user
	user, err := s.identityRepo.GetUserByID(ctx, uint64(storedToken.UserID))
	if err != nil {
        return "", err
	}

	// Generate a new access token for this user
	accessToken, err := auth.GenerateAccessToken(user, s.AccessTokenDuration)
	if err != nil {
		return "", err
	}

	// Return access token
	return accessToken, nil
}

func (s *AuthService) storeRefreshToken(ctx context.Context, refreshToken string, user *entity.User) error {

	hashedToken := pkg.HashToken(refreshToken)
	expiresAt := time.Now().Add(time.Duration(s.RefreshTokenDuration))

	refreshTokenData := entity.RefreshToken{
		UserID: 		user.ID,
		TokenHash: 		hashedToken,
		ExpiresAt:		expiresAt,
	}

	if err := s.identityRepo.StoreRefreshToken(ctx, &refreshTokenData); err != nil {
		log.Printf("failed to store refresh token: %w", err)
		
		return err
	}
	return nil
}


// Create table for refresh_tokens
// Store refrash_token after generation at the table 
// When click logout -> revoke refrash token from the DB 