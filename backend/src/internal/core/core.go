package core

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"llmapi/src/internal/config"
	"llmapi/src/internal/middleware"
	"llmapi/src/internal/repository"
	internalRouter "llmapi/src/internal/router"
	"llmapi/src/internal/utils/log"
	"llmapi/src/pkg/logger"
)

var (
	router *gin.Engine
	cfg    *config.Config
)

func InitServer() error {
	logger.InitDefaultLogger()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		// Skip error if .env file is not found
		if !os.IsNotExist(err) {
			return fmt.Errorf("error loading .env file: %w", err)
		}
	}

	// Load configuration
	cfg = config.LoadConfig()
	logger.SetLevelString(cfg.LogLevel)

	// Initialize database connection
	db, err := repository.InitDB(cfg)
	if err != nil {
		return err
	}

	router = gin.New()

	// Set middleware for the engine
	router.Use(middleware.RequestLogger())
	router.Use(gin.Recovery())

	internalRouter.SetupRouter(&internalRouter.Options{
		Engine: router,
		DB:     db,
		Cfg:    cfg,
	})

	return nil
}

func Run() {
	if router == nil {
		panic("router is not initialized")
	}
	address := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Sys().Info("Starting server...", "host", address)
	if err := router.Run(address); err != nil {
		log.Sys().Error("Failed to start server", "error", err)
	}
	log.Sys().Info("Exit...")
}
