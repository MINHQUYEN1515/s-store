package routes

import (
	"s-store/internal/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup, authHandler *handler.AuthHandler) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}
}
