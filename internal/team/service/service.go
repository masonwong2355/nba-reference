package service

import (
	"context"
	"nba-predictor/internal/team"
	"nba-predictor/internal/team/repository"
)

type Svc struct {
	repository repository.Repositroy
}

func New(repository repository.Repositroy) *Svc {
	return &Svc{
		repository: repository,
	}
}

func GetTeam(ctx context.Context, parmas team.GetTeamPramas) (*[]team.Team, error) {
	return nil, nil
}
