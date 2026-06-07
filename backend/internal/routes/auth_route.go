package routes

import "github.com/gin-gonic/gin"

func AuthRoute(r *gin.RouterGroup) {
	auth := r.Group("/auth")

}
