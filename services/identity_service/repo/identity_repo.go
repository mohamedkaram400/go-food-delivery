package repo

import (
	"context"

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


func (r *IdentityRepo) GetUserByEmail(ctx context.Context, email string) {

}