package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nba-predictor/internal/api/config"
	"nba-predictor/internal/logger"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
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
