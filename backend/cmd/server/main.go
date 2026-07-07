package main

import (
	"fmt"
	"log"

	backend "s-store"
	"s-store/internal/database"
	"s-store/internal/migration"
	"s-store/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config := backend.LoadConfig()

	db, err := database.ConnectDatabase(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	if err := migration.Migrate(db); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database connection: %v", err)
	}
	defer sqlDB.Close()
	log.Println("Connected to database successfully")

	if config.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	server := gin.Default()
	routes.AppRoute(server)

	addr := fmt.Sprintf(":%s", config.AppPort)
	log.Printf("Server is running on %s", addr)
	if err := server.Run(addr); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
