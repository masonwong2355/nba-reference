package repository

import "context"

type Repository interface {
	getTeam(ctx context.Context, param *GetTeamParams) error
}

type GetTeamParams struct {
	teamID string
}
