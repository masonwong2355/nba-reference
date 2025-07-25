package scraper

import (
	"fmt"
	"log"
	"nba-predictor/internal/models"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
)

func ScrapeTeamData(db *gorm.DB) {
	fmt.Println("Start scraping team data...")

	url := "https://www.espn.com/nba/teams"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible)")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

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
