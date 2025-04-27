package dto

import (
	"khajuraho/backend/model"
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;index"`
	Token     string    `gorm:"type:varchar(255);index"`
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Request/Response Structs
type GoogleLoginRequest struct {
	IDToken   string  `json:"id_token"`
	Device    string  `json:"device"`
	IP        string  `json:"ip"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UserAgent string  `json:"user_agent"`
}

type AuthTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GoogleLoginResponse struct {
	User   model.User        `json:"user"`
	Tokens AuthTokenResponse `json:"tokens"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token"`
}
