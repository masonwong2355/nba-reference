package repository

import (
	"context"
	"nba-reference/internal/team"

	"gorm.io/gorm"
)

type Repository interface {
	// GetTeams(ctx context.Context, param *GetTeamParams) ([]team.Team, error)
	GetTeams(ctx context.Context) ([]team.Team, error)
}

type GetTeamParams struct {
	teamID string
}

type GormRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}

func (r *GormRepository) GetTeams(ctx context.Context) ([]team.Team, error) {
	var teams []team.Team
	if err := r.db.WithContext(ctx).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
