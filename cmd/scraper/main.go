package main

import (
	"fmt"
	"log"
	"nba-predictor/internal/scraper"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// go run cmd/scraper/main.go game
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run cmd/scraper/main.go [team|player|schedule|...]")
	}
	cmd := os.Args[1]

	dsn := "host=localhost user=postgres password=password dbname=nba_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	switch cmd {
	case "team":
		scraper.ScrapeTeamData(db)
	case "player":
		scraper.ScrapePlayerData(db)
	case "game":
		scraper.ScrapeGameData(db)
	default:
		fmt.Println("No function matching")
	}
}
