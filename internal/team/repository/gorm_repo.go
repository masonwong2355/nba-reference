package repository

import "context"

type Repositroy interface {
	getTeam(ctx context.Context, param *GetTeamPramas) error
}

type GetTeamPramas struct {
	teamId string
}
