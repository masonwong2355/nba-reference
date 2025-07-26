package scraper

import (
	"fmt"
	"nba-predictor/internal/models"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func ScrapeTeamData(db *gorm.DB) {
	log.Info().Msg("Start scraping team data")

	url := "https://www.espn.com/nba/teams"
	doc := getPageDoc(url)

	doc.Find("section.TeamLinks").Each(func(i int, s *goquery.Selection) {
		teamName := s.Find(".di.clr-gray-01.h5").Text()

		teamID := ""
		href, ok := s.Find("a.AnchorLink").Attr("href")
		if ok {
			parts := strings.Split(href, "/")
			teamID = parts[len(parts)-2]
		}

		team := models.Team{
			Name:   teamName,
			TeamID: teamID,
		}

		result := db.Create(&team)
		if result.Error != nil {
			log.Error().Err(result.Error).Str("team", teamName).Msg("Failed to insert")
		} else {
			log.Info().Str("team", teamName).Msg("Inserted team")
		}
	})

	log.Info().Msg("End scraping team data")
}
