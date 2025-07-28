package team

import (
	"context"
	"time"
)

// Hexagonal Architecture

type Service interface {
	GetTeam(ctx context.Context, parmas GetTeamPramas) (*[]Team, error)
}

type Team struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TeamID    string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type GetTeamPramas struct {
	teamId string
}
