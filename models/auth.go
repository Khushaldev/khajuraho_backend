package models

import (
	"time"

	"github.com/google/uuid"
)

type AuthProvider struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;index"`
	Provider  string    `gorm:"type:varchar(50)"`
	LoginType string    `gorm:"type:varchar(50)"`
	Latitude  float64   `gorm:"type:decimal(9,6)"`
	Longitude float64   `gorm:"type:decimal(9,6)"`
	IP        string    `gorm:"type:varchar(45)"`
	UserAgent string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
