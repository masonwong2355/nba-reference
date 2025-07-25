package scraper

import (
	"fmt"
	"log"
	"nba-predictor/internal/models"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
)

func ScrapeTeamData(db *gorm.DB) {
	fmt.Println("Start scraping team data...")

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
			log.Printf("Failed to insert %s: %v\n", teamName, result.Error)
		} else {
			fmt.Println("Inserted:", teamName)
		}
	})

	fmt.Println("End scraping team data...")
}
