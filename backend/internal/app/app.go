package app

import (
	"fmt"
	"log"
	backend "s-store"
	"s-store/internal/database"
	"s-store/internal/handler"
	"s-store/internal/middleware"
	"s-store/internal/migration"
	repository "s-store/internal/repositories"
	"s-store/internal/routes"
	"s-store/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup() (*gin.Engine, error) {
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
	server := gin.New()
	if err := server.SetTrustedProxies(nil); err != nil {
		return nil, err
	}

	// Allow CORS for all origins
	server.Use(middleware.AllowCors())

	server.Use(CustomLogger())
	server.Use(gin.Recovery())

	// Set up dependency injection for handlers and services
	setupDenpendencyInjection(server, db, config)

	return server, nil

}

func setupDenpendencyInjection(r *gin.Engine, db *gorm.DB, config backend.Config) {
	// Initialize repositories
	authRepo := repository.NewAuthRepo(db)

	// Initialize services
	authService := service.AuthServiceNew(authRepo, config.JWTKey)

	// Initialize handlers
	authHandler := handler.AuthHandlerNew(authService)

	routes.AppRoute(r, &routes.RouteHandler{
		AuthHandler: authHandler,
	})
}

func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		clientIP := c.ClientIP()

		fmt.Printf(
			"[%s] %3d | %13v | %-7s %s | %s\n",
			time.Now().Format("2006-01-02 15:04:05"),
			status,
			latency,
			method,
			path,
			clientIP,
		)
	}
}
