package internalfacing

import (
	"nba-reference/internal/team"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, svc team.Service) {
	router.GET("/teams", GetTeamsHandler(svc))
}
