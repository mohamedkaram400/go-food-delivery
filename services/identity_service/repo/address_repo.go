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

func (r *AddressRepo) NumberOfAddress(ctx context.Context, userID uint64) (int64, error) {
    var count int64

    err := r.DB.WithContext(ctx).
        Model(&entity.Address{}).
        Where("user_id = ?", userID).
        Count(&count).Error

    if err != nil {
        return 0, err
    }

    return count, nil
}