package dto

import (
	"middleware_jwt/pkg"
)

type LoginRequest struct {
	Email    string `json:"email" valo:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token pkg.Token    `json:"token"`
}
