package main

import (
	"fmt"
	"log"

	backend "s-store"
	"s-store/internal/database"
	"s-store/internal/middleware"
	"s-store/internal/migration"
	"s-store/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config := backend.LoadConfig()

	// Connect to the database
	db, err := database.ConnectDataBase(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	// Perform database migration
	if err := migration.Migrate(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	// Get the underlying sql.DB from gorm.DB
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}
	defer sqlDB.Close()
	log.Println("Connected to database successfully")

	// Set Gin mode based on environment
	if config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	// Initialize Gin router and set up routes
	server := gin.Default()
	routes.AppRoute(server)

	// Allow CORS for all origins
	server.Use(middleware.AllowCors())

	addr := fmt.Sprintf(":%s", config.AppPort)
	log.Printf("Server is running on %s", addr)
	if err := server.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
