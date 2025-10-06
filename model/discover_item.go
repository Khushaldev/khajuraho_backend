package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type DiscoverItem struct {
	ID               uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name             string         `gorm:"type:varchar(100);not null" json:"name"`
	Slug             string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"slug"`
	CategoryID       uuid.UUID      `gorm:"type:uuid;not null;index" json:"category_id"`
	Category         Category       `gorm:"foreignKey:CategoryID" json:"category"`
	AddressID        uuid.UUID      `gorm:"type:uuid;not null;index" json:"address_id"`
	Address          Address        `gorm:"foreignKey:AddressID" json:"address"`
	Description      string         `gorm:"type:text" json:"description,omitempty"`
	AttributeGroupID int64          `gorm:"type:integer;not null;index" json:"attribute_group_id"`
	AttributeGroup   AttributeGroup `gorm:"foreignKey:AttributeGroupID" json:"attribute_group"` // Use attribute_group for consistency
	IsActive         bool           `gorm:"default:true;not null" json:"isActive"`              // Added isActive field
	CreatedAt        time.Time      `json:"createdAt"`
	UpdatedAt        time.Time      `json:"updatedAt"`

	// Relationships (managed via join tables or separate tables)
	Images []*DiscoverItemImage `gorm:"foreignKey:DiscoverItemID" json:"images,omitempty"`
	Tags   []*Tag               `gorm:"many2many:discover_item_tags;" json:"tags,omitempty"`
}

type DiscoverItemImage struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	DiscoverItemID uuid.UUID `gorm:"type:uuid;not null;index" json:"discover_item_id"`
	ImageURL       string    `gorm:"type:varchar(2048);not null" json:"imageUrl"`
	AltText        string    `gorm:"type:varchar(255)" json:"altText,omitempty"`
	SortOrder      int       `gorm:"default:0;not null" json:"sortOrder"`
	IsPrimary      bool      `gorm:"default:false;not null" json:"isPrimary"`
	CreatedAt      time.Time `json:"createdAt"`
}

type Tag struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(50);not null;unique" json:"name"`
	// DiscoverItems []*DiscoverItem `gorm:"many2many:discover_item_tags;" json:"-"` // Optional back-reference
}

// Join table for DiscoverItem and Tag (Many-to-Many)
type DiscoverItemTag struct {
	DiscoverItemID uuid.UUID `gorm:"type:uuid;primaryKey" json:"discover_item_id"`
	TagID          uint      `gorm:"primaryKey" json:"tag_id"`
	CreatedAt      time.Time `json:"createdAt"`
}

// --- Attributes ---

type AttributeType struct {
	Id   int64  `gorm:"type:integer;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(100);not null;unique" json:"name"` // e.g., TEXT, NUMBER, BOOLEAN, DATE, JSON
}

type Attribute struct {
	Id              int64         `gorm:"type:integer;primaryKey;autoIncrement" json:"id"`
	Name            string        `gorm:"type:varchar(100);not null;unique" json:"name"` // e.g., Color, Size, Price, Has Wifi
	Tag             string        `gorm:"type:varchar(100);not null;unique" json:"tag"`  // e.g., color, size, price, has_wifi (for machine use)
	Description     string        `gorm:"type:text" json:"description,omitempty"`
	AttributeTypeID int64         `gorm:"type:integer;not null;index" json:"attribute_type_id"`
	Type            AttributeType `gorm:"foreignKey:AttributeTypeID" json:"type"`
	// Value string `gorm:"-" json:"value"` // This was for temporary holding, actual values are in DiscoverItemAttributeValue
}

type AttributeGroup struct {
	Id         int64        `gorm:"type:integer;primaryKey;autoIncrement" json:"id"`
	Name       string       `gorm:"type:varchar(100);not null;unique" json:"name"`
	Attributes []*Attribute `gorm:"many2many:attribute_group_attributes;" json:"attributes"`
}

// Join table for AttributeGroup and Attribute (Many-to-Many)
type AttributeGroupAttribute struct {
	AttributeGroupID int64 `gorm:"primaryKey" json:"attribute_group_id"`
	AttributeID      int64 `gorm:"primaryKey" json:"attribute_id"`
	// SortOrder int `gorm:"default:0" json:"sortOrder"` // Optional: if order of attributes within a group matters
}

type DiscoverItemAttributeValue struct {
	ID             uint64         `gorm:"primaryKey;autoIncrement" json:"id"`
	DiscoverItemID uuid.UUID      `gorm:"type:uuid;not null;index:idx_item_attr_unique,unique" json:"discover_item_id"`
	AttributeID    int64          `gorm:"not null;index:idx_item_attr_unique,unique" json:"attribute_id"`
	Value          datatypes.JSON `gorm:"type:jsonb" json:"value"` // Stores the actual value, type determined by Attribute.AttributeTypeID
	// DiscoverItem   DiscoverItem `gorm:"foreignKey:DiscoverItemID"` // Optional back-references
	// Attribute      Attribute    `gorm:"foreignKey:AttributeID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
