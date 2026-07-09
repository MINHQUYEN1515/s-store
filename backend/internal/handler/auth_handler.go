package handler

import (
	"s-store/internal/model/request"
	"s-store/internal/model/response"
	"s-store/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
}

func AuthHandlerNew(service service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: service,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var request request.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, 400, err.Error(), nil)
		return
	}
	result, err := h.authService.Register(request)
	if err != nil {
		response.Error(c, 400, err.Error(), nil)
		return
	}
	response.Success(c, 200, "Registration successful", result)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request request.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, 400, err.Error(), nil)
		return
	}

	result, err := h.authService.Login(request)
	if err != nil {
		response.Error(c, 400, err.Error(), nil)
		return
	}
	response.Success(c, 200, "Login successful", result)
}
