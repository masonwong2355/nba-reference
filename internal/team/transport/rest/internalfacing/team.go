package internalfacing

import (
	"nba-reference/internal/team"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeamsHandler(svc team.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := &team.GetTeamParams{
			TeamID: c.Query("teamID"),
			Name:   c.Query("name"),
		}

		teams, err := svc.GetTeams(c.Request.Context(), params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, teams)
	}
}
