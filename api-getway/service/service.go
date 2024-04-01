package service

import (
	"api-test/config"
	commentpb "api-test/genproto/comment-service"
	postpb "api-test/genproto/post-service"
	pb "api-test/genproto/user-service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pb.UserServiceClient
	PostService() postpb.PostServiceClient
	CommentService() commentpb.CommentServiceClient
}

type serviceManager struct {
	userService    pb.UserServiceClient
	postService    postpb.PostServiceClient
	commentService commentpb.CommentServiceClient
}

func (s *serviceManager) UserService() pb.UserServiceClient {
	return s.userService
}

func (s *serviceManager) PostService() postpb.PostServiceClient {
	return s.postService
}

func (s *serviceManager) CommentService() commentpb.CommentServiceClient {
	return s.commentService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	connPost, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PostServiceHost, conf.PostServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	connComment, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CommentServiceHost, conf.CommentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		userService:    pb.NewUserServiceClient(connUser),
		postService:    postpb.NewPostServiceClient(connPost),
		commentService: commentpb.NewCommentServiceClient(connComment),
	}
	return serviceManager, nil
}
