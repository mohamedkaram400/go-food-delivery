package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/client"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/config"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/handler"
	"github.com/mohamed-karam/go-food-delivery/api-gateway-service/routes"
)



func main() {
	// Load config file 
	cfg := config.Load()

    // ------------------------------------
	// Connect to the DB
    // ------------------------------------
	

    // ------------------------------------
	// Connect to Identity Service
    // ------------------------------------
	identityClient, identityConn, err :=
		client.NewIdentityClient(cfg.IdentityService)

	if err != nil {
		log.Fatal(err)
	}

	defer identityConn.Close()
    // ------------------------------------
	// Register REST endpoint
	// ------------------------------------

	authHandler := handler.NewAuthHandler(identityClient)


	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger(), gin.Recovery())

	// 7. Versioned API group
	v1 := router.Group("/api/v1")


	routes.AuthRoutes(v1, authHandler)

	// ------------------------------------
	// Start REST server
	// ------------------------------------

	log.Println("API Gateway running on ", cfg.Port)

	if err := router.Run(cfg.Port); err != nil {
		log.Fatal(err)
	}
}