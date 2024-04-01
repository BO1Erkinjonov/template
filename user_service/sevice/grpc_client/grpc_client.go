package grpc_client

import (
	"fmt"
	"google.golang.org/grpc"
	"user_service/config"
	comment "user_service/genproto/comment-service"
	pb "user_service/genproto/post-service"
)

type IServiceManager interface {
	PostService() pb.PostServiceClient
	CommentService() comment.CommentServiceClient
}

type serviceManager struct {
	cfg            config.Config
	postService    pb.PostServiceClient
	commentService comment.CommentServiceClient
}

func New(cfg config.Config) (IServiceManager, error) {
	templatePostConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.PostServiceHost, cfg.PostServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d", cfg.PostServiceHost, cfg.PostServicePort)
	}
	commentServiceConn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", cfg.CommentServiceHost, cfg.CommentServicePort),
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, fmt.Errorf("post service dial host: %s port: %d", cfg.CommentServiceHost, cfg.CommentServicePort)
	}

	return &serviceManager{
		cfg:            cfg,
		postService:    pb.NewPostServiceClient(templatePostConn),
		commentService: comment.NewCommentServiceClient(commentServiceConn),
	}, nil
}

func (s *serviceManager) PostService() pb.PostServiceClient {
	return s.postService
}

func (s *serviceManager) CommentService() comment.CommentServiceClient {
	return s.commentService
}
