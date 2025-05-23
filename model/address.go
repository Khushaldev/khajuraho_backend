package model

import (
	"time"

	"github.com/google/uuid"
)

type Country struct {
	ID   uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name string    `gorm:"unique;not null"`
	Code string    `gorm:"unique"`
}

type State struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string    `gorm:"not null"`
	Code      string    `gorm:"unique"`
	Type      string    `gorm:"type:varchar(20);default:'state'"`
	CountryID uuid.UUID `gorm:"type:uuid;not null"`
	Country   Country   `gorm:"foreignKey:CountryID"`
}

type District struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name    string    `gorm:"not null"`
	StateID uuid.UUID `gorm:"type:uuid;not null"`
	State   State     `gorm:"foreignKey:StateID"`
}

type City struct {
	ID         uuid.UUID  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name       string     `gorm:"not null"`
	StateID    uuid.UUID  `gorm:"type:uuid;not null"`
	State      State      `gorm:"foreignKey:StateID"`
	DistrictID *uuid.UUID `gorm:"type:uuid"`
	District   *District  `gorm:"foreignKey:DistrictID"`
}

type Area struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name   string    `gorm:"not null"`
	CityID uuid.UUID `gorm:"type:uuid;not null"`
	City   City      `gorm:"foreignKey:CityID"`
}

type Address struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Line1     string    `gorm:"type:varchar(100)"`
	Line2     string    `gorm:"type:varchar(100)"`
	AreaID    uuid.UUID `gorm:"type:uuid;not null"`
	Area      Area      `gorm:"foreignKey:AreaID"`
	GeoPoint  string    `gorm:"type:geography(POINT,4326)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
