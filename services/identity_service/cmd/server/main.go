package main

import (
	"log"
	"net"

	"github.com/mohamed-karam/go-food-delivery/identity-service/config"
	"github.com/mohamed-karam/go-food-delivery/identity-service/conn"
	"github.com/mohamed-karam/go-food-delivery/identity-service/handler"
	pb "github.com/mohamed-karam/go-food-delivery/identity-service/proto/identity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/repo"
	"github.com/mohamed-karam/go-food-delivery/identity-service/service"
	"google.golang.org/grpc"
)



func main() {
	// ------------------------------------
	// Load config file 
	// ------------------------------------
	cfg := config.Load()

	// ------------------------------------
	// Connect to the DB
    // ------------------------------------
	mysql, err := conn.ConnectMySQL(cfg.DSN)
	if err != nil {
		log.Fatal("❌ Failed to connect MySQL:", err)
	}

	sqlDB, err := mysql.DB()
	if err != nil {
		log.Fatal("❌ Failed to get sql.DB:", err)
	}

	defer sqlDB.Close()

	
    if err := migration.Run(mysql); err != nil {
        log.Fatal(err)
    }

	// ------------------------------------
	// Identity Service
	// ------------------------------------
	grpcServer := grpc.NewServer()
    identityRepo := repo.NewIdentityRepo(mysql)
    identityService := service.NewIdentityService(identityRepo)
    grpcHandler := handler.NewIdentityHandler(identityService)

	
	pb.RegisterIdentityServiceServer(grpcServer, grpcHandler)

	// ------------------------------------
	// Start gRPC server
	// ------------------------------------
	listener, err := net.Listen("tcp", cfg.Port)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Identity Service gRPC running on ", cfg.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}