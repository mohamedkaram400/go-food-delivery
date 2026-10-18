package handler

import (
	"context"
	"log"

	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/requests"
	"github.com/mohamed-karam/go-food-delivery/identity-service/errors"
	"github.com/mohamed-karam/go-food-delivery/identity-service/service"

)


type AddressHandler struct { 
	AddressService *service.AddressService
	pb.UnimplementedAddressServiceServer
}

func NewAddressHandler(AddressService *service.AddressService) *AddressHandler {
	return &AddressHandler{
		AddressService: AddressService,
	}
}

func (s *AddressHandler) CreateAddresses(ctx context.Context, req *pb.CreateAddressesRequest) (*pb.CreateAddressesResponse, error) {

	log.Println("🔥 IDENTITY: Store Address gRPC handler reached")

	// Prepare the register request
	addressesRequest := &requests.CreateAddressesRequest{
		Addresses:     req.Addresses,
	}
	log.Printf("📥 IDENTITY: Request received: %+v", addressesRequest)

	// Call the register endpoint to get the token and user
	addresses, err := s.AddressService.Store(ctx, addressesRequest)
	if err != nil {
		log.Printf("❌ IDENTITY: Service error: %v", err)

		return nil, errors.ToGRPCError(err)
	}

	log.Printf("✅ IDENTITY: Addresses created: %+v", addresses)

	response := &pb.CreateAddressesResponse{
		Addresses: make([]*pb.Address, 0, len(addresses)),
	}

	for _, address := range addresses {
		response.Addresses = append(response.Addresses, &pb.Address{
			Id: int64(address.ID),
			Address: address.Address,
			UserId: int64(address.UserID),
		})
	}
	
	// Return the address response 
	return response, nil
}
