package handler

import (
	"context"
	"log"
	"strconv"

	"github.com/mohamed-karam/go-food-delivery/identity-service/errors"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/requests"
	"github.com/mohamed-karam/go-food-delivery/identity-service/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

	// Extract metadata from incoming gRPC request.
	md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return nil, status.Error(codes.Unauthenticated, "Missing authentication metadata")
    }

	values := md.Get("x-user-id")
	if len(values) != 1 {
        return nil, status.Error(codes.Unauthenticated, "Missing user ID")
    }

    userID, err := strconv.ParseInt(values[0], 10, 64)
    if err != nil || userID <= 0 {
        return nil, status.Error(codes.Unauthenticated, "Invalid user ID")
    }


	// Prepare the register request
	addressesRequest := &requests.CreateAddressesRequest{
		Addresses:     req.Addresses,
	}
	log.Printf("📥 IDENTITY: Request received: %+v", addressesRequest)

	// Call the register endpoint to get the token and user
	addresses, err := s.AddressService.Store(ctx, userID, addressesRequest)
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
 