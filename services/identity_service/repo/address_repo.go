package repo

import (
	"context"

	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"gorm.io/gorm"
)


type AddressRepo struct {
    DB *gorm.DB
}

func NewAddressRepo(db *gorm.DB) *AddressRepo {
	return &AddressRepo{
		DB: db,
	}
}

func (r *AddressRepo) Store(ctx context.Context, address *entity.Address) (*entity.Address, error) {
	
	if err := r.DB.WithContext(ctx).Create(address).Error; err != nil {
		return nil, err
	}

	return address, nil
}

func (r *AddressRepo) GetUserByID(ctx context.Context, userID uint64) (*entity.User, error) {
    var user entity.User

    if err := r.DB.WithContext(ctx).
        First(&user, userID).
        Error; err != nil {
        return nil, err
    }

    return &user, nil
}