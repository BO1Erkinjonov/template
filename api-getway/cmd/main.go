package main

import (
	"api-test/api"
	"api-test/config"
	"api-test/pkg/logger"
	"api-test/service"
	"github.com/casbin/casbin/v2"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := service.NewServiceManager(&cfg)
	if err != nil {
		log.Fatal("gRPC dial error", logger.Error(err))
	}

	enforcer, err := casbin.NewEnforcer("auth.conf")

	if err != nil {
		log.Error("NewEnforcer error", logger.Error(err))
		return
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		CasbinEnforcer: enforcer,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		return
	}

}
