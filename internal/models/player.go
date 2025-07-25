package models

import (
	"time"
)

type Player struct {
	ID       string `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ESPNID   string `gorm:"type:varchar(30);not null;uniqueIndex"`
	FullName string `gorm:"type:varchar(100);not null"`
	TeamID   string `gorm:"type:uuid"` // Foreign key
	// Team         Team      `gorm:"foreignKey:TeamID"` // GORM relation (optional)
	JerseyNumber *int      `gorm:"type:integer"`
	Position     string    `gorm:"type:varchar(40)"`
	Height       string    `gorm:"type:varchar(20)"`
	Weight       int       `gorm:"type:integer"`
	Birthdate    time.Time `gorm:"type:date"`
	Experience   string    `gorm:"type:varchar(20)"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
