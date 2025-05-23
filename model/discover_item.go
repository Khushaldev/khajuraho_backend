package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Attribute struct {
	Id    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DiscoverItem struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name         string         `gorm:"type:varchar(100);not null" json:"name"`
	Slug         string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
	CategoryID   uuid.UUID      `gorm:"type:uuid;not null" json:"category_id"`
	Category     Category       `gorm:"foreignKey:CategoryID" json:"category"`
	AddressID    uuid.UUID      `gorm:"type:uuid;not null" json:"address_id"`
	Address      Address        `gorm:"foreignKey:AddressID" json:"address"`
	Description  string         `gorm:"type:text" json:"description,omitempty"`
	Images       pq.StringArray `gorm:"type:text[]" json:"images"`
	Tags         pq.StringArray `gorm:"type:text[]" json:"tags,omitempty"`
	Rating       float64        `gorm:"type:decimal(2,1);default:0" json:"rating,omitempty"`
	ReviewsCount int            `gorm:"default:0" json:"reviewsCount,omitempty"`
	Metadata     map[string]any `gorm:"type:jsonb" json:"metadata"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
}
