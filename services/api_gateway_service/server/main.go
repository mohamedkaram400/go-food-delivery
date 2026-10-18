package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/client"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/config"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/handler/identity"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/routes"
)



func main() {
    // -------------------
	// Load config file 
    // -------------------
	cfg := config.Load()

    // ------------------------------------
	// Connect to Auth Service
    // ------------------------------------
	conn, err := client.NewIdentityClient(cfg.IdentityService)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	
	// ------------------------------------
	// Register Client
	// ------------------------------------
	authClient := client.NewAuthClient(conn)
	addressClient := client.NewAddressClient(conn)

    // ------------------------------------
	// Register REST endpoint
	// ------------------------------------
	authHandler := handler.NewAuthHandler(authClient)
	addressHandler := handler.NewAddressHandler(addressClient)

 
	// ------------------------------------
	// SetUp gin routes
	// ------------------------------------
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger(), gin.Recovery())

	// ---------------------------------------
	// Register the routes with their handler
	// ---------------------------------------
	routes.RegisterRoutes(router, routes.Handlers{
		Auth: authHandler,
		Address: addressHandler,
	})

	// ------------------------------------
	// Start REST server
	// ------------------------------------
	log.Println("API Gateway running on ", cfg.Port)

	if err := router.Run(cfg.Port); err != nil {
		log.Fatal(err)
	}
}