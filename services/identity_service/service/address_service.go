package service

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/repo"
	"github.com/mohamed-karam/go-food-delivery/identity-service/requests"
)

type AddressService struct {
	AddressRepo  *repo.AddressRepo
}

func NewAddressService(addressRepo *repo.AddressRepo) *AddressService {
	return &AddressService{
		AddressRepo:  addressRepo,
	}
} 

func (s *AddressService) Store(ctx context.Context, req *requests.CreateAddressesRequest) ([]*entity.Address, error) {

	// Check validation errors
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

    addresses := make([]*entity.Address, 0, len(req.Addresses))

	for _, name := range req.Addresses {

		// Prepare address object
		addressObj := &entity.Address{
			Address:         name,
			// UserID:          userID, 
		}

		// Pass the user object to repo for creation
		address, err := s.AddressRepo.Store(ctx, addressObj); 
		if err != nil {
			return nil, err
		}

		addresses = append(addresses, address)
		log.Printf("Address created successfully: address_id=%d", address.ID)
	}

	return addresses, nil
}
