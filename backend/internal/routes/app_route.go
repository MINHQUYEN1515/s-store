package routes

import (
	"s-store/internal/handler"

	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	AuthHandler *handler.AuthHandler
}

func AppRoute(r *gin.Engine, routeHandler *RouteHandler) {
	appGroup := r.Group("api/v1")
	AuthRoute(appGroup, routeHandler.AuthHandler)
}
