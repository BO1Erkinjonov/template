package repo

import (
	pb "post_service/genproto/post-service"
)

type PostStorageI interface {
	Create(user *pb.Post) (*pb.Post, error)
	GetPost(id string) (*pb.Post, error)
	GetAllPosts(page, limit int32) (posts []*pb.Post, err error)
	UpdatePost(post *pb.Post) (*pb.Post, error)
	DeletePost(id string) (bool, error)
	GetPostByOwnerId(id string) ([]*pb.Post, error)
	DeletePostByOwnerId(id string) (bool, error)
}
