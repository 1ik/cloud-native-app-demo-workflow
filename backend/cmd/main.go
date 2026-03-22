package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// logGmailEnv logs Gmail-related env vars at startup for verifying SecretKeyRef wiring.
// It intentionally does not log the app password value (logs are often copied/shared).
func logGmailEnv() {
	user := os.Getenv("GMAIL_USERNAME")
	pass, passOK := os.LookupEnv("GMAIL_APP_PASSWORD")

	log.Printf("gmail env: GMAIL_USERNAME=%q", user)
	switch {
	case !passOK:
		log.Print("gmail env: GMAIL_APP_PASSWORD not set (missing env var)")
	case pass == "":
		log.Print("gmail env: GMAIL_APP_PASSWORD set but empty")
	default:
		log.Printf("gmail env: GMAIL_APP_PASSWORD is set (length=%d)", len(pass))
	}
}

func main() {
	logGmailEnv()

	// Create Gin router
	r := gin.Default()

	// Redis client (shared for handlers)
	rdb := newRedisClient()

	// Enable CORS for local frontend
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"*"},
	}))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "cloudnativeapp-api",
		})
	})

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "CloudNativeApp API",
			"version": "0.1.0",
		})
	})

	// Hit counter endpoint
	r.GET("/hit", HitHandler(rdb))
	r.GET("/hit2", HitHandler(rdb))

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
