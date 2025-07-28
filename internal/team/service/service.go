package service

import (
	"context"
	"nba-reference/internal/team"
	"nba-reference/internal/team/repository"
)

type Svc struct {
	repository repository.Repository
}

func New(repository repository.Repository) *Svc {
	return &Svc{
		repository: repository,
	}
}

func (s *Svc) GetTeams(ctx context.Context, params *team.GetTeamParams) ([]team.Team, error) {
	return s.repository.GetTeams(ctx, params)
}
