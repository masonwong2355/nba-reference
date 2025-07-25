package scraper

import (
	"fmt"
	"log"
	"math/rand"
	"nba-predictor/internal/models"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// to scraper
// https://www.espn.com/nba/schedule/_/date/20220111 and https://www.espn.com/nba/game/_/gameId/401360426/bucks-hornets
// for the basic game data, e.g. scroe, date, q1-q4 score, referees
func ScrapeGameData(db *gorm.DB) {
	fmt.Println("Start scraping game data...")

	// 1396, 1406 -> 16 - 17

	// configs
	maxWorkers := 6
	start := time.Date(2022, 1, 11, 0, 0, 0, 0, time.UTC)
	end := time.Date(2022, 1, 13, 0, 0, 0, 0, time.UTC)
	// start := time.Date(2017, 4, 14, 0, 0, 0, 0, time.UTC)
	// end := time.Date(2017, 6, 20, 0, 0, 0, 0, time.UTC)
	season := fmt.Sprintf("%d-%d", start.Year(), end.Year())

	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	var count int64

	failedScraperGameID := []string{}

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		d := d
		sem <- struct{}{}
		wg.Add(1)
		go func(day time.Time) {
			defer wg.Done()
			defer func() { <-sem }()
			scrapeGamesForDate(db, season, day, &failedScraperGameID, &count)
			time.Sleep(time.Millisecond * time.Duration(200+rand.Intn(600)))
		}(d)
	}
	wg.Wait()

	fmt.Println("failedScraperGameID:", failedScraperGameID)
	fmt.Println("End scraping game data... total data: ", count)
}

func scrapeGamesForDate(db *gorm.DB, season string, d time.Time, failedScraperGameID *[]string, count *int64) {
	dcount := 0

	datePath := fmt.Sprintf("https://www.espn.com/nba/schedule/_/date/%s", d.Format("20060102"))
	dateDoc := getPageDoc(datePath)

	dateDoc.Find("table.Table").First().Find("tbody.Table__TBODY").Find("tr.Table__TR").Each(func(i int, s *goquery.Selection) {
		// s := doc.Find("table.Table").First().Find("tbody.Table__TBODY").Find("tr.Table__TR").First()

		// get game ID
		aTag := s.Find("td.teams__col").Find("a")

		if aTag.Text() == "Postponed" {
			*failedScraperGameID = append(*failedScraperGameID, datePath)
			return
		}
		gamePath, _ := aTag.Attr("href")
		paths := strings.Split(gamePath, "/")
		if len(paths) == 1 {
			*failedScraperGameID = append(*failedScraperGameID, gamePath)
			return
		}
		gameID := paths[5]

		gamePath = fmt.Sprintf("https://www.espn.com%s", gamePath)
		doc := getPageDoc(gamePath)

		// get game type
		gtype := "regular"
		gNote := doc.Find("div.ScoreCell__GameNote").Text()
		if gNote != "" && gNote == "Preseason" {
			gtype = "preseason"
		}
		if gNote != "" && gNote != "Preseason" {
			gtype = "postseason"
		}

		// get team ID
		ah, _ := doc.Find(".Gamestrip__Team").First().Find("a").Attr("href")
		hh, _ := doc.Find(".Gamestrip__Team").Last().Find("a").Attr("href")
		if ah == "" || hh == "" {
			*failedScraperGameID = append(*failedScraperGameID, gameID)
			return
		}

		awayID := strings.Split(ah, "/")[5]
		homeID := strings.Split(hh, "/")[5]

		// get game score
		as := doc.Find(".Gamestrip__Overview").Find(".Table__TBODY").Find(".Table__TR").First()
		aQ1S := as.Find("td").Eq(1).Text()
		aQ2S := as.Find("td").Eq(2).Text()
		aQ3S := as.Find("td").Eq(3).Text()
		aQ4S := as.Find("td").Eq(4).Text()
		aFinanS := as.Find("td").Eq(5).Text()

		hs := doc.Find(".Gamestrip__Overview").Find(".Table__TBODY").Find(".Table__TR").Last()
		hQ1S := hs.Find("td").Eq(1).Text()
		hQ2S := hs.Find("td").Eq(2).Text()
		hQ3S := hs.Find("td").Eq(3).Text()
		hQ4S := hs.Find("td").Eq(4).Text()
		hFinanS := hs.Find("td").Eq(5).Text()

		// get game info, e.g. time, referee, location
		t := doc.Find(".GameInfo__Meta").Find("span").First().Text()
		dateTime, _ := time.Parse("3:04 PM, January 2, 2006", t)

		area := doc.Find(".Location__Text").Text()

		referees := ""
		doc.Find(".GameInfo__List__Wrapper").Find("li").Each(func(i int, s *goquery.Selection) {
			r := s.Text()
			if r == "Referees:" {
				return
			}

			if referees == "" {
				referees += r
			} else {
				referees += ", " + r
			}
		})

		// build model
		g := models.Game{
			ESPNID:       gameID,
			StartTime:    dateTime,
			SeasonYear:   season,
			Type:         gtype,
			HomeTeamID:   homeID,
			AwayTeamID:   awayID,
			HomeScore:    stringToInt(hFinanS),
			HomeQ1Score:  stringToInt(hQ1S),
			HomeQ2Score:  stringToInt(hQ2S),
			HomeQ3Score:  stringToInt(hQ3S),
			HomeQ4Score:  stringToInt(hQ4S),
			AwayScore:    stringToInt(aFinanS),
			AwayQ1Score:  stringToInt(aQ1S),
			AwayQ2Score:  stringToInt(aQ2S),
			AwayQ3Score:  stringToInt(aQ3S),
			AwayQ4Score:  stringToInt(aQ4S),
			Arena:        area,
			Referees:     referees,
			WinnerTeamID: homeID,
		}

		result := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&g)
		if result.Error != nil {
			log.Printf("Failed to insert %s: %v\n", gameID, result.Error)
		} else {
			fmt.Println("Inserted:", gameID)
			if result.RowsAffected > 0 {
				atomic.AddInt64(count, 1)
			}
		}

		dcount += 1
	})

	fmt.Println("data count of date: ", d.Format("20060102"), " count: ", dcount)
	fmt.Println("----------------------------------------")
}
