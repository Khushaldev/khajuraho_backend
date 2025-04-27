package model

import (
	"time"

	"github.com/google/uuid"
)

type RefreshSession struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;index;not null"`
	Token     string    `gorm:"type:varchar(255);uniqueIndex;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	Device    string    `gorm:"type:varchar(100);not null"`
	IP        string    `gorm:"type:varchar(45);not null"`
	UserAgent string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Revoked   bool `gorm:"default:false"`
}
