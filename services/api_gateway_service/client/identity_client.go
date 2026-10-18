package client

import (

	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewIdentityClient(address string) (*grpc.ClientConn, error) {

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func NewAuthClient(conn *grpc.ClientConn) pb.AuthServiceClient {
	return pb.NewAuthServiceClient(conn)
}

func NewAddressClient(conn *grpc.ClientConn) pb.AddressServiceClient {
	return pb.NewAddressServiceClient(conn)
} 