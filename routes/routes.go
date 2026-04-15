package routes

import (
	"ai-customer-support/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(messageHandler *handlers.MessageHandler) *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Routes
	r.POST("/message", messageHandler.HandleMessage)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	return r
}
