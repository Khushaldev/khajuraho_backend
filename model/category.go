package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Slug      string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Name      string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Icon      string    `gorm:"type:varchar(255)"`
	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SubCategory struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Category_Id uuid.UUID `gorm:"type:uuid;index;not null"`
	Slug        string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Name        string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Icon        string    `gorm:"type:varchar(255)"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
