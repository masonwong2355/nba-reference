package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nba-reference/internal/api/config"
	"nba-reference/internal/logger"
	"nba-reference/internal/team/repository"
	teamservice "nba-reference/internal/team/service"
	teamtransport "nba-reference/internal/team/transport/rest/internalfacing"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 1. Load config first
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "base"
	}
	cfg, err := config.LoadConfig("config/" + env + ".yaml")
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}

	// 2. Setup logger
	logger.InitLogger(cfg.Env == "debug")
	log.Info().Str("env", cfg.Env).Msg("Starting NBA predictor server!")

	// 3. Optionally set Gin mode from config
	// gin.SetMode(cfg.GinMode) // if you want to use this

	// 4. Set up Gin and HTTP server
	router := gin.Default()

	// Initialize database connection
	dsn := "host=localhost user=postgres password=password dbname=nba_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect database")
	}
	// Create repository and service then register routes
	repo := repository.New(db)
	teamSvc := teamservice.New(repo)
	teamtransport.AddRoutes(router, teamSvc)

	srv := &http.Server{
		Addr:    ":8080", // or cfg.Port
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("server error")
		}
	}()

	// 5. Graceful shutdown handler
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("forced to shutdown")
	}
	log.Info().Msg("Server exited")
}
