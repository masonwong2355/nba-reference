package scraper

import (
	"fmt"
	"nba-predictor/internal/models"
	"nba-predictor/internal/team"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func ScrapePlayerData(db *gorm.DB) {
	log.Info().Msg("Start scraping player data")

	var teams []team.Team
	err := db.Find(&teams).Error
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get teams")
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
				log.Error().Err(result.Error).Str("name", name).Msg("Failed to insert player")
			} else {
				log.Info().Str("name", name).Msg("Inserted player")
			}

			count += 1
		})
		log.Info().Int("count", count).Msg("End scraping player data")

	}
	// ----------------------------------------------------------------------------------------
}
