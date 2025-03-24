package main

import (
	"log"

	"api-gateway/config"
	"api-gateway/internal/middleware"
	"api-gateway/internal/routes"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Cannot load config:", err)
	}

	middleware.InitCircuitBreaker()

	r := routes.SetupRouter()
	cfg := config.GetConfig()
	log.Printf("Starting API Gateway on port %s", cfg.Server.Port)
	log.Fatal(r.Run(":" + cfg.Server.Port))
}
