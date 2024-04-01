package grpc_client

import (
	"comment_service/config"
	pbp "comment_service/genproto/post-service"
	pbc "comment_service/genproto/user-service"
	"fmt"
	"google.golang.org/grpc"
)

type IServiceManager interface {
	PostService() pbp.PostServiceClient
	UserService() pbc.UserServiceClient
}

type serviceManager struct {
	cfg                 config.Config
	templatePostService pbp.PostServiceClient
	templateUserService pbc.UserServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	templatePostConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d", cfg.PostServiceHost, cfg.PostServicePort)
	}
	templateUserConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d", cfg.UserServiceHost, cfg.UserServicePort)
	}
	return &serviceManager{
		cfg:                 cfg,
		templatePostService: pbp.NewPostServiceClient(templatePostConn),
		templateUserService: pbc.NewUserServiceClient(templateUserConn),
	}, nil
}

func (s *serviceManager) PostService() pbp.PostServiceClient {
	return s.templatePostService
}

func (s *serviceManager) UserService() pbc.UserServiceClient {
	return s.templateUserService
}
