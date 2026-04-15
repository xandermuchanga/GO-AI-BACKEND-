package main

import (
	"log"

	"ai-customer-support/config"
	"ai-customer-support/handlers"
	"ai-customer-support/routes"
	"ai-customer-support/services"
)

func main() {
	// 1. Load Configuration
	cfg := config.LoadConfig()

	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is required")
	}

	// 2. Initialize Services
	aiService, err := services.NewAIService(cfg.GeminiAPIKey)
	if err != nil {
		log.Fatalf("Failed to initialize AI Service: %v", err)
	}
	defer aiService.Close()

	// 3. Initialize Handlers
	messageHandler := handlers.NewMessageHandler(aiService)

	// 4. Setup Routes
	router := routes.SetupRouter(messageHandler)

	// 5. Start Server
	log.Printf("Server starting on port %s...", cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
