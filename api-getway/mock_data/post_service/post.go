package post_service

import (
	pb "api-test/genproto/post-service"
	"context"
)

type MockPostServiceClient interface {
	Create(ctx context.Context, in *pb.Post) (*pb.Post, error)
	GetPost(ctx context.Context, in *pb.GetRequests) (*pb.PostResponse, error)
	GetAllPost(ctx context.Context, in *pb.GetAllPostRequest) (*pb.GetAllPostResponse, error)
	UpdatePost(ctx context.Context, in *pb.Post) (*pb.Post, error)
	DeletePost(ctx context.Context, in *pb.GetRequests) (*pb.Tf, error)
	GetPostByOwnerId(ctx context.Context, in *pb.GetByOwnerIdRequest) (*pb.GetByOwnerIdResponse, error)
	DeletePostByOwnerId(ctx context.Context, in *pb.GetByOwnerIdRequest) (*pb.Tf, error)
}

type mockPostServiceClient struct {
}

func NewMockPostServiceClient() MockPostServiceClient {
	return &mockPostServiceClient{}
}

func (c *mockPostServiceClient) DeletePostByOwnerId(ctx context.Context, in *pb.GetByOwnerIdRequest) (*pb.Tf, error) {
	return &pb.Tf{Tf: true}, nil
}

func (c *mockPostServiceClient) Create(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	return &pb.Post{
		Title:     "Mock title",
		Content:   "Mock content",
		ImageUrl:  "Mock imageurl",
		Id:        "Mock id",
		OwnerId:   "Mock owner id",
		Likes:     0,
		Views:     0,
		Category:  "Mock category",
		CreatedAt: "Mock cteated at",
		UpdatedAt: "Mock updated at",
		Comment:   nil,
		Owner:     nil,
	}, nil
}

func (c *mockPostServiceClient) GetPost(ctx context.Context, in *pb.GetRequests) (*pb.PostResponse, error) {
	return &pb.PostResponse{
		Title:     "Mock title",
		Content:   "Mock content",
		ImageUrl:  "Mock imageurl",
		Id:        "Mock id",
		Likes:     0,
		Views:     0,
		Category:  "Mock category",
		CreatedAt: "Mock cteated at",
		UpdatedAt: "Mock updated at",
		Comment:   nil,
		Owner:     nil,
	}, nil
}
func (c *mockPostServiceClient) GetAllPost(ctx context.Context, in *pb.GetAllPostRequest) (*pb.GetAllPostResponse, error) {
	posts := []*pb.Post{
		{
			Title:     "Mock title",
			Content:   "Mock content",
			ImageUrl:  "Mock imageurl",
			Id:        "Mock id",
			Likes:     0,
			Views:     0,
			Category:  "Mock category",
			CreatedAt: "Mock cteated at",
			UpdatedAt: "Mock updated at",
			Comment:   nil,
			Owner:     nil,
		},
		{
			Title:     "Mock1 title",
			Content:   "Mock1 content",
			ImageUrl:  "Mock1 imageurl",
			Id:        "Mock1 id",
			Likes:     0,
			Views:     0,
			Category:  "Mock1 category",
			CreatedAt: "Mock1 cteated at",
			UpdatedAt: "Mock1 updated at",
			Comment:   nil,
			Owner:     nil,
		},
		{
			Title:     "Mock2 title",
			Content:   "Mock2 content",
			ImageUrl:  "Mock2 imageurl",
			Id:        "Mock2 id",
			Likes:     0,
			Views:     0,
			Category:  "Mock2 category",
			CreatedAt: "Mock2 cteated at",
			UpdatedAt: "Mock2 updated at",
			Comment:   nil,
			Owner:     nil,
		},
	}
	return &pb.GetAllPostResponse{
		Posts: posts,
	}, nil
}
func (c *mockPostServiceClient) UpdatePost(ctx context.Context, in *pb.Post) (*pb.Post, error) {
	return &pb.Post{
		Title:     "Mock title",
		Content:   "Mock content",
		ImageUrl:  "Mock imageurl",
		Id:        "Mock id",
		OwnerId:   "Mock owner id",
		Likes:     0,
		Views:     0,
		Category:  "Mock category",
		CreatedAt: "Mock cteated at",
		UpdatedAt: "Mock updated at",
		Comment:   nil,
		Owner:     nil,
	}, nil
}
func (c *mockPostServiceClient) DeletePost(ctx context.Context, in *pb.GetRequests) (*pb.Tf, error) {
	return &pb.Tf{
		Tf: true,
	}, nil
}
func (c *mockPostServiceClient) GetPostByOwnerId(ctx context.Context, in *pb.GetByOwnerIdRequest) (*pb.GetByOwnerIdResponse, error) {
	posts := []*pb.Post{
		{
			Title:     "Mock title",
			Content:   "Mock content",
			ImageUrl:  "Mock imageurl",
			Id:        "Mock id",
			Likes:     0,
			Views:     0,
			Category:  "Mock category",
			CreatedAt: "Mock cteated at",
			UpdatedAt: "Mock updated at",
			Comment:   nil,
			Owner:     nil,
		},
		{
			Title:     "Mock1 title",
			Content:   "Mock1 content",
			ImageUrl:  "Mock1 imageurl",
			Id:        "Mock1 id",
			Likes:     0,
			Views:     0,
			Category:  "Mock1 category",
			CreatedAt: "Mock1 cteated at",
			UpdatedAt: "Mock1 updated at",
			Comment:   nil,
			Owner:     nil,
		},
		{
			Title:     "Mock2 title",
			Content:   "Mock2 content",
			ImageUrl:  "Mock2 imageurl",
			Id:        "Mock2 id",
			Likes:     0,
			Views:     0,
			Category:  "Mock2 category",
			CreatedAt: "Mock2 cteated at",
			UpdatedAt: "Mock2 updated at",
			Comment:   nil,
			Owner:     nil,
		},
	}
	return &pb.GetByOwnerIdResponse{
		Posts: posts,
	}, nil
}
