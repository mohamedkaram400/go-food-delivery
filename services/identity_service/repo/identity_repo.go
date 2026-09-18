package repo

import (
	"context"

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
	var user *entity.User

	if err := r.DB.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}