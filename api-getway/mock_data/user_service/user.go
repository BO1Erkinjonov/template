package user_service

import (
	ps "api-test/genproto/post-service"
	pbu "api-test/genproto/user-service"
	"context"
)

type MockUserServiceClient interface {
	Create(ctx context.Context, in *pbu.User) (*pbu.User, error)
	GetUser(ctx context.Context, in *pbu.GetRequest) (*pbu.User, error)
	GetAllUsers(ctx context.Context, in *pbu.GetAllUsersRequest) (*pbu.GetAllUsersResponse, error)
	Update(ctx context.Context, in *pbu.User) (*pbu.User, error)
	Delete(ctx context.Context, in *pbu.GetRequest) (*pbu.Tf, error)
	CheckUniquess(ctx context.Context, in *pbu.CheckUniqReq) (*pbu.CheckUniqResp, error)
	Exists(ctx context.Context, in *pbu.Req) (*pbu.User, error)
}

type mockUserServiceClient struct {
}

func NewMockUserServiceClient() MockUserServiceClient {
	return &mockUserServiceClient{}
}

func (c *mockUserServiceClient) Create(ctx context.Context, in *pbu.User) (*pbu.User, error) {
	return &pbu.User{
		FirstName:    "Mock first_name",
		LastName:     "Mock last_name",
		Username:     "Mock username",
		Password:     "Mock password",
		Email:        "Mock email",
		Id:           "9566058e-1426-48f1-be48-821276227934",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhMWQxZDg2NC1lMGY3LTExZWUtYThiNS0wNDdjMTZhMTcyMDYifQ.6zRopXFolv69RPiDohlUpLhDDpsd13GyQmHf1YSkrYo",
		Post:         nil,
	}, nil
}

func (c *mockUserServiceClient) GetUser(ctx context.Context, in *pbu.GetRequest) (*pbu.User, error) {
	return &pbu.User{
		FirstName:    "Mock first_name",
		LastName:     "Mock last_name",
		Username:     "Mock username",
		Password:     "Mock password",
		Email:        "Mock email",
		Id:           "Mock id",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhMWQxZDg2NC1lMGY3LTExZWUtYThiNS0wNDdjMTZhMTcyMDYifQ.6zRopXFolv69RPiDohlUpLhDDpsd13GyQmHf1YSkrYo",
		Post:         nil,
	}, nil
}

func (c *mockUserServiceClient) Update(ctx context.Context, in *pbu.User) (*pbu.User, error) {
	return &pbu.User{
		FirstName:    "Mock first_name",
		LastName:     "Mock last_name",
		Username:     "Mock username",
		Password:     "Mock password",
		Email:        "Mock email",
		Id:           "Mock id",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhMWQxZDg2NC1lMGY3LTExZWUtYThiNS0wNDdjMTZhMTcyMDYifQ.6zRopXFolv69RPiDohlUpLhDDpsd13GyQmHf1YSkrYo",
		Post:         nil,
	}, nil
}

func (c *mockUserServiceClient) Delete(ctx context.Context, in *pbu.GetRequest) (*pbu.Tf, error) {
	return &pbu.Tf{Tf: true}, nil
}

func (c *mockUserServiceClient) CheckUniquess(ctx context.Context, in *pbu.CheckUniqReq) (*pbu.CheckUniqResp, error) {
	return &pbu.CheckUniqResp{
		Code: 1,
	}, nil
}

func (c *mockUserServiceClient) Exists(ctx context.Context, in *pbu.Req) (*pbu.User, error) {
	return &pbu.User{
		FirstName:    "Mock first_name",
		LastName:     "Mock last_name",
		Username:     "Mock username",
		Password:     "Mock password",
		Email:        "Mock email",
		Id:           "9566058e-1426-48f1-be48-821276227934",
		RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhMWQxZDg2NC1lMGY3LTExZWUtYThiNS0wNDdjMTZhMTcyMDYifQ.6zRopXFolv69RPiDohlUpLhDDpsd13GyQmHf1YSkrYo",
		Post:         nil,
	}, nil
}

func (c *mockUserServiceClient) CreatePost(ctx context.Context, in *ps.Post) (*ps.Post, error) {
	return &ps.Post{
		Title:    "Mock title",
		ImageUrl: "Mock url",
		Id:       "Mock id",
		OwnerId:  "Mock owner id",
		Comment:  nil,
	}, nil
}

func (c *mockUserServiceClient) GetAllUsers(ctx context.Context, in *pbu.GetAllUsersRequest) (*pbu.GetAllUsersResponse, error) {
	users := []*pbu.User{
		{
			FirstName:    "Mock first_name",
			LastName:     "Mock last_name",
			Username:     "Mock username",
			Password:     "Mock password",
			Email:        "Mock email",
			Id:           "9566058e-1426-48f1-be48-821276227934",
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhMWQxZDg2NC1lMGY3LTExZWUtYThiNS0wNDdjMTZhMTcyMDYifQ.6zRopXFolv69RPiDohlUpLhDDpsd13GyQmHf1YSkrYo",
			Post:         nil,
		},
		{
			FirstName:    "Mock first_name 1",
			LastName:     "Mock last_name 1",
			Username:     "Mock username 1",
			Password:     "Mock password 1",
			Email:        "Mock email 1",
			Id:           "9566058e-1426-48f1-be48-821276227934",
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhMWQxZDg2NC1lMGY3LTExZWUtYThiNS0wNDdjMTZhMTcyMDYifQ.6zRopXFolv69RPiDohlUpLhDDpsd13GyQmHf1YSkrYo",
			Post:         nil,
		},
		{
			FirstName:    "Mock first_name 2",
			LastName:     "Mock last_name 2",
			Username:     "Mock username 2",
			Password:     "Mock password 2",
			Email:        "Mock email 2",
			Id:           "9566058e-1426-48f1-be48-821276227934",
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhMWQxZDg2NC1lMGY3LTExZWUtYThiNS0wNDdjMTZhMTcyMDYifQ.6zRopXFolv69RPiDohlUpLhDDpsd13GyQmHf1YSkrYo",
			Post:         nil,
		},
		{
			FirstName:    "Mock first_name 3",
			LastName:     "Mock last_name 3",
			Username:     "Mock username 3",
			Password:     "Mock password 3",
			Email:        "Mock email 3",
			Id:           "9566058e-1426-48f1-be48-821276227934",
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhMWQxZDg2NC1lMGY3LTExZWUtYThiNS0wNDdjMTZhMTcyMDYifQ.6zRopXFolv69RPiDohlUpLhDDpsd13GyQmHf1YSkrYo",
			Post:         nil,
		},
	}

	return &pbu.GetAllUsersResponse{
		Users: users,
	}, nil
}
