package main

import (
	"fmt"
	"log"

	"example/web-service-gin"
	"example/web-service-gin/internal/database"
	"example/web-service-gin/internal/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := backend.LoadConfig()

	db, err := database.ConnectDatabase(config)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()
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
