package grpc_client

import (
	"fmt"
	"google.golang.org/grpc"
	"post_service/config"
	pbc "post_service/genproto/comment-service"
	pbp "post_service/genproto/user-service"
)

type IServiceManager interface {
	UserService() pbp.UserServiceClient
	CommentService() pbc.CommentServiceClient
}

type serviceManager struct {
	cfg            config.Config
	userService    pbp.UserServiceClient
	commentService pbc.CommentServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	UserConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.UserServiceHost, cfg.UserServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d", cfg.UserServiceHost, cfg.UserServicePort)
	}
	CommentConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d", cfg.CommentServiceHost, cfg.CommentServicePort)
	}
	return &serviceManager{
		cfg:            cfg,
		userService:    pbp.NewUserServiceClient(UserConn),
		commentService: pbc.NewCommentServiceClient(CommentConn),
	}, nil
}

func (s *serviceManager) UserService() pbp.UserServiceClient {
	return s.userService
}

func (s *serviceManager) CommentService() pbc.CommentServiceClient {
	return s.commentService
}
