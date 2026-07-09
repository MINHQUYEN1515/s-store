package response

import "s-store/internal/model/entity"

type AuthResponse struct {
	Token       string            `json:"token"`
	RefeshToken string            `json:"refresh_token"`
	User        entity.UserEntity `json:"user"`
}
