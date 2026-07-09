package app

import (
	"log"
	backend "s-store"
	"s-store/internal/database"
	"s-store/internal/handler"
	"s-store/internal/middleware"
	"s-store/internal/migration"
	repository "s-store/internal/repositories"
	"s-store/internal/routes"
	"s-store/internal/service"

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
	server := gin.Default()

	// Allow CORS for all origins
	server.Use(middleware.AllowCors())

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

	// Set up routes with handlers
	routes.AuthRoute(&r.RouterGroup, authHandler)

	routes.AppRoute(r, &routes.RouteHandler{
		AuthHandler: authHandler,
	})
}
