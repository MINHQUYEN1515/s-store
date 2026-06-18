package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

func Success(c *gin.Context, code int, message string, data any) {
	c.JSON(http.StatusOK, APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string, data any) {
	c.JSON(status, APIResponse{
		Code:    status,
		Message: message,
		Data:    data,
	})
}
