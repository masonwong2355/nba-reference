package internalfacing

import (
	"nba-reference/internal/team"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTeamsHandler(svc team.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		teams, err := svc.GetTeams(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, teams)
	}
}
