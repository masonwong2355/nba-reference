package repository

import (
	"context"
	"nba-reference/internal/team"

	"gorm.io/gorm"
)

type Repository interface {
	GetTeams(ctx context.Context, params *team.GetTeamParams) ([]team.Team, error)
}

type GormRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db: db,
	}
}

func (r *GormRepository) GetTeams(ctx context.Context, params *team.GetTeamParams) ([]team.Team, error) {
	var teams []team.Team

	query := r.db.WithContext(ctx)
	if params != nil {
		if params.TeamID != "" {
			query = query.Where("team_id = ?", params.TeamID)
		}
		if params.Name != "" {
			query = query.Where("name = ?", params.Name)
		}
	}

	if err := query.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
