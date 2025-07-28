package team

import (
	"context"
	"time"
)

// Hexagonal Architecture

type Service interface {
	GetTeams(ctx context.Context, params *GetTeamParams) ([]Team, error)
}

type Team struct {
	ID        string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TeamID    string    `gorm:"not null"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type GetTeamParams struct {
	TeamID string
	Name   string
}
