package main

import (
	"os"

	"nba-reference/internal/logger"
	"nba-reference/internal/scraper"

	"github.com/rs/zerolog/log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// go run cmd/scraper/main.go game
func main() {
	logger.InitLogger(true)

	if len(os.Args) < 2 {
		log.Fatal().Msg("Usage: go run cmd/scraper/main.go [team|player|schedule|...]")
	}
	cmd := os.Args[1]

	dsn := "host=localhost user=postgres password=password dbname=nba_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}

	switch cmd {
	case "team":
		scraper.ScrapeTeamData(db)
	case "player":
		scraper.ScrapePlayerData(db)
	case "game":
		scraper.ScrapeGameData(db)
	default:
		log.Info().Msg("No function matching")
	}
}
