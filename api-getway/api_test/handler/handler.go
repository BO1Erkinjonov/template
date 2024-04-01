package handler

import (
	"api-test/api/tokens"
	"api-test/config"
	"api-test/pkg/logger"
	"api-test/service"
	"github.com/casbin/casbin/v2"
)

type handlerV1 struct {
	jwthandler     token.JWTHandler
	log            logger.Logger
	serviceManager service.IServiceManager
	cfg            config.Config
}

// HandlerV1Config ...
type HandlerV1Config struct {
	Jwthandler     token.JWTHandler
	Logger         logger.Logger
	ServiceManager service.IServiceManager
	Cfg            config.Config
	Enforcer       casbin.Enforcer
}

// New ...
func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		jwthandler:     c.Jwthandler,
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
	}
}
