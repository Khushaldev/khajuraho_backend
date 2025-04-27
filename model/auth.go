package model

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTToken struct {
	Token     string
	ExpiresAt time.Time
}

type AuthTokens struct {
	AccessToken  JWTToken
	RefreshToken JWTToken
}

type AuthSession struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID         uuid.UUID `gorm:"type:uuid;index"`
	Provider       string    `gorm:"type:varchar(50)"`
	ProviderUserId string    `gorm:"type:varchar(100)"`
	LoginType      string    `gorm:"type:varchar(50)"`
	Latitude       float64   `gorm:"type:decimal(9,6)"`
	Longitude      float64   `gorm:"type:decimal(9,6)"`
	IP             string    `gorm:"type:varchar(45)"`
	UserAgent      string    `gorm:"type:varchar(255)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
