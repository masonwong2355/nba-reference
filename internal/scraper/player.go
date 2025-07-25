package scraper

import (
	"fmt"
	"log"
	"nba-predictor/internal/models"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
)

func ScrapePlayerData(db *gorm.DB) {
	fmt.Println("Start scraping player data...")

	var teams []models.Team
	err := db.Find(&teams).Error
	if err != nil {
		log.Fatal("Failed to get teams:", err)
	}

	teamIDs := []string{}
	for _, t := range teams {
		teamIDs = append(teamIDs, t.TeamID)
	}

	for _, teamID := range teamIDs {
		url := fmt.Sprintf("https://www.espn.com/nba/team/stats/_/name/%s/season/2025/seasontype/2/split/33", teamID)
		doc := getPageDoc(url)

		// ----------------------------------------------------------------------------------------
		count := 0
		espnID, name, linkName, playerUrl := "", "", "", ""
		doc.Find("tbody.Table__TBODY").First().Find("a.AnchorLink").Each(func(i int, s *goquery.Selection) {
			name = s.Text()

			href, ok := s.Attr("href")
			if ok {
				parts := strings.Split(href, "/")
				espnID = parts[len(parts)-2]
				linkName = parts[len(parts)-1]
			}

			playerUrl = fmt.Sprintf("https://www.espn.com/nba/player/bio/_/id/%s/%s", espnID, linkName)
			playerDoc := getPageDoc(playerUrl)

			// team := ""
			// if h, ok := playerDoc.Find("ul.PlayerHeader__Team_Info").Find("a.AnchorLink").Attr("href"); ok {
			// 	p := strings.Split(h, "/")
			// 	team = p[len(p)-1]
			// }

			bio := map[string]string{}
			playerDoc.Find("section.Bio").Find("div.Bio__Item").Each(func(i int, s *goquery.Selection) {
				label := s.Find(".Bio__Label").Text()
				value := s.Find(".flex-uniform").Text()

				bio[label] = value
			})
			// fmt.Print("bio ", bio)

			// height := playerDoc.Find("section.Bio").Find("div.Bio__Item").Eq(2).Find(".flex-uniform").Text()
			height := bio["HT/WT"]
			parts := strings.Split(height, ",")
			heightVal := ""
			weightVal := ""
			if len(parts) == 2 {
				heightVal = strings.TrimSpace(parts[0])
				weightVal = strings.TrimSpace(strings.TrimSuffix(parts[1], " lbs"))
			}

			weightInt := 0
			if weightVal != "" {
				weightInt, _ = strconv.Atoi(weightVal)
			}

			birthdate := bio["Birthdate"]
			birthdateStr := birthdate // e.g. "3/3/1998 (27)"
			if strings.Contains(birthdateStr, " ") {
				birthdateStr = strings.Split(birthdateStr, " ")[0] // "3/3/1998"
			}
			var birthdateTime time.Time
			if birthdateStr != "" {
				// Parse as MM/DD/YYYY
				birthdateTime, _ = time.Parse("1/2/2006", birthdateStr)
			}

			experience := bio["Experience"]
			// Get Jersey Number and Position from Header
			header := playerDoc.Find(".PlayerHeader__Team_Info").Text()
			re := regexp.MustCompile(`#(\d+)`)
			match := re.FindStringSubmatch(header)
			var JerseyNumber *int
			if len(match) > 1 {
				n, _ := strconv.Atoi(match[1])
				JerseyNumber = &n
			}
			position := playerDoc.Find(".PlayerHeader__Team_Info").Find("li").Last().Text()

			player := models.Player{
				ESPNID:       espnID,
				FullName:     name,
				TeamID:       teamID,
				JerseyNumber: JerseyNumber,
				Position:     position,
				Height:       heightVal,
				Weight:       weightInt,
				Birthdate:    birthdateTime,
				Experience:   experience,
			}

			result := db.Create(&player)
			if result.Error != nil {
				log.Printf("Failed to insert %s: %v\n", name, result.Error)
			} else {
				fmt.Println("Inserted:", name)
			}

			count += 1
		})
		fmt.Println("End scraping player data... total data: ", count)

	}
	// ----------------------------------------------------------------------------------------
}
