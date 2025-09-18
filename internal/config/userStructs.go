package config

import (
	"time"

	"github.com/google/uuid"
)

type userReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userResp struct {
	Id           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	IsRed        bool      `json:"is_chirpy_red"`
}
