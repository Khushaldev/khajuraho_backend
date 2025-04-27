package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	accessTokenTTL  = 15 * time.Minute    // 15 minutes
	refreshTokenTTL = 30 * 24 * time.Hour // 30 days
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	DisplayName    string    `gorm:"not null" json:"display_name"`
	Email          string    `gorm:"uniqueIndex;not null" json:"email"`
	ProfilePicture string    `gorm:"default:''" json:"profile_picture"`
	CountryCode    string    `gorm:"default:''" json:"country_code"`
	PhoneNumber    string    `gorm:"default:''" json:"phone_number"`
	IsActive       bool      `gorm:"default:true" json:"is_active"`
	IsVerified     bool      `gorm:"default:false" json:"is_verified"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
