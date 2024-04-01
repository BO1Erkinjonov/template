package comment_service

import (
	pbc "api-test/genproto/comment-service"
	"context"
)

type MockCommentServiceServer interface {
	Create(context.Context, *pbc.Comment) (*pbc.Comment, error)
	GetComment(context.Context, *pbc.Get) (*pbc.Comment, error)
	GetAllComment(context.Context, *pbc.GetRequest) (*pbc.GetResponse, error)
	UpdateComment(context.Context, *pbc.Comment) (*pbc.Comment, error)
	DeleteComment(context.Context, *pbc.Get) (*pbc.Tf, error)
	GetCommentByPostId(context.Context, *pbc.GetCommentByPostIdRequest) (*pbc.GetCommentByPostIdResponse, error)
}

type mockCommentServiceClient struct {
}

func NewMockCommentServiceClient() MockCommentServiceServer {
	return &mockCommentServiceClient{}
}

func (c *mockCommentServiceClient) Create(context.Context, *pbc.Comment) (*pbc.Comment, error) {
	return &pbc.Comment{
		Id:          "Mock id",
		Description: "Mock descrition",
		PostId:      "Mock post id",
		OwnerId:     "Mock owner id",
		CreatedAt:   "Mock created at",
		UpdatedAt:   "Mock updated at",
		User:        nil,
		Post:        nil,
	}, nil
}

func (c *mockCommentServiceClient) GetComment(context.Context, *pbc.Get) (*pbc.Comment, error) {
	return &pbc.Comment{
		Id:          "Mock id",
		Description: "Mock descrition",
		PostId:      "Mock post id",
		OwnerId:     "Mock owner id",
		CreatedAt:   "Mock created at",
		UpdatedAt:   "Mock updated at",
		User:        nil,
		Post:        nil,
	}, nil
}

func (c *mockCommentServiceClient) GetAllComment(context.Context, *pbc.GetRequest) (*pbc.GetResponse, error) {
	comments := []*pbc.Comment{
		{
			Id:          "Mock id",
			Description: "Mock descrition",
			PostId:      "Mock post id",
			OwnerId:     "Mock owner id",
			CreatedAt:   "Mock created at",
			UpdatedAt:   "Mock updated at",
			User:        nil,
			Post:        nil,
		},
		{
			Id:          "Mock 1 id",
			Description: "Mock 1 descrition",
			PostId:      "Mock 1 post id",
			OwnerId:     "Mock 1 owner id",
			CreatedAt:   "Mock 1 created at",
			UpdatedAt:   "Mock 1 updated at",
			User:        nil,
			Post:        nil,
		},
		{
			Id:          "Mock 2 id",
			Description: "Mock 2 descrition",
			PostId:      "Mock 2 post id",
			OwnerId:     "Mock 2 owner id",
			CreatedAt:   "Mock 2 created at",
			UpdatedAt:   "Mock 2 updated at",
			User:        nil,
			Post:        nil,
		},
	}
	return &pbc.GetResponse{
		Comments: comments,
	}, nil
}

func (c *mockCommentServiceClient) UpdateComment(context.Context, *pbc.Comment) (*pbc.Comment, error) {
	return &pbc.Comment{
		Id:          "Mock id",
		Description: "Mock descrition",
		PostId:      "Mock post id",
		OwnerId:     "Mock owner id",
		CreatedAt:   "Mock created at",
		UpdatedAt:   "Mock updated at",
		User:        nil,
		Post:        nil,
	}, nil
}

func (c *mockCommentServiceClient) DeleteComment(context.Context, *pbc.Get) (*pbc.Tf, error) {
	return &pbc.Tf{
		Tf: true,
	}, nil
}

func (c *mockCommentServiceClient) GetCommentByPostId(context.Context, *pbc.GetCommentByPostIdRequest) (*pbc.GetCommentByPostIdResponse, error) {
	comments := []*pbc.Comment{
		{
			Id:          "Mock id",
			Description: "Mock descrition",
			PostId:      "Mock post id",
			OwnerId:     "Mock owner id",
			CreatedAt:   "Mock created at",
			UpdatedAt:   "Mock updated at",
			User:        nil,
			Post:        nil,
		},
		{
			Id:          "Mock 1 id",
			Description: "Mock 1 descrition",
			PostId:      "Mock 1 post id",
			OwnerId:     "Mock 1 owner id",
			CreatedAt:   "Mock 1 created at",
			UpdatedAt:   "Mock 1 updated at",
			User:        nil,
			Post:        nil,
		},
		{
			Id:          "Mock 2 id",
			Description: "Mock 2 descrition",
			PostId:      "Mock 2 post id",
			OwnerId:     "Mock 2 owner id",
			CreatedAt:   "Mock 2 created at",
			UpdatedAt:   "Mock 2 updated at",
			User:        nil,
			Post:        nil,
		},
	}
	return &pbc.GetCommentByPostIdResponse{
		Comments: comments,
	}, nil
}
