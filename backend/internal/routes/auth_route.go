package routes

import "github.com/gin-gonic/gin"

func AuthRoute(r *gin.RouterGroup) {
	_ = r.Group("/auth")
}
