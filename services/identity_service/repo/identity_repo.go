package repo

import (
	"context"
	"errors"
	"time"

	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"gorm.io/gorm"
)


type IdentityRepo struct {
    DB *gorm.DB
}

func NewIdentityRepo(db *gorm.DB) *IdentityRepo {
	return &IdentityRepo{
		DB: db,
	}
}
 
func (r *IdentityRepo) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	
	if err := r.DB.WithContext(ctx).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *IdentityRepo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *IdentityRepo) GetUserByPhone(ctx context.Context, phone string) (*entity.User, error) {
	var user entity.User

	if err := r.DB.WithContext(ctx).Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *IdentityRepo) StoreRefreshToken(ctx context.Context, refreshToken *entity.RefreshToken) (error) {
	if err := r.DB.WithContext(ctx).Create(refreshToken).Error; err != nil {
		return err
	}

	return nil
}

func (r *IdentityRepo) RevokeRefreshToken(ctx context.Context, hashedToken string) (error) {
	result := r.DB.WithContext(ctx).
		Model(&entity.RefreshToken{}).
		Where("token_hash = ?", hashedToken).
		Where("revoked_at IS NULL").
		Update("revoked_at", time.Now())

	if result.Error != nil {
        return result.Error
    }

    if result.RowsAffected == 0 {
        return errors.New("refresh token not found or already revoked")
    }

    return nil
}

