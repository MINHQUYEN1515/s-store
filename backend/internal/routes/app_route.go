package routes

import "github.com/gin-gonic/gin"

func AppRoute(r *gin.Engine) {
	appGroup := r.Group("api/v1")
	AuthRoute(appGroup)

}
