package client

import (

	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewIdentityClient(address string) (pb.IdentityServiceClient, *grpc.ClientConn, error) {

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client := pb.NewIdentityServiceClient(conn)

	return client, nil, nil
}