package team

import "time"

type Service interface {
}

type Team struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TeamID    string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
