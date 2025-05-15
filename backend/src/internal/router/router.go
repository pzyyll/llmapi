package router

import (
	"fmt"

	"llmapi/src/internal/config"
	"llmapi/src/internal/repository"
	"llmapi/src/internal/router/dashboard"
	"llmapi/src/internal/service"
	"llmapi/src/internal/utils"
	"llmapi/src/internal/utils/log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type (
	Router  struct{}
	Options struct {
		Engine *gin.Engine
		DB     *gorm.DB
		Cfg    *config.Config
	}
)

func SetupRouter(opts *Options) {
	// Set up the router for the application
	engine := opts.Engine
	cfg := opts.Cfg

	engine.GET("/ping", Ping)

	uidGenerator, err := utils.NewUidGenerator(cfg.WorkerID)
	if err != nil {
		panic(fmt.Errorf("Get uid generator fail %v", err.Error()))
	}

	userRepo := repository.NewUserRepo(opts.DB)
	userSvc := service.NewUserService(userRepo, opts.Cfg, uidGenerator)

	if err := userSvc.InitAdminUser(); err != nil {
		log.Sys().Error("Init admin user fail.", "error", err.Error())
	}

	authSvc := service.NewAuthService(userSvc, cfg, uidGenerator)

	dashboard.SetupRouter(&dashboard.Options{
		Engine:  engine,
		UserSvc: userSvc,
		AuthSvc: authSvc,
		Cfg:     opts.Cfg,
	})
}
