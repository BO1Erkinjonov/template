package mongodb

import (
	"github.com/stretchr/testify/suite"
	"post_service/config"
	pbp "post_service/genproto/post-service"
	"post_service/pkg/db"
	"post_service/storage/repo"
	"testing"
)

type PostRepositoryTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.PostStorageI
}

func (s *PostRepositoryTestSuite) SetupSuite() {
	pgPool, cleanUp := db.ConnectToMongoDBForSuite(config.Load())
	s.Repository = NewPostRepoMongo(pgPool)
	s.CleanUpFunc = cleanUp
}

func (s *PostRepositoryTestSuite) TestPostCRUD() {
	post := &pbp.Post{
		Title:     "Test title",
		Content:   "Test Content",
		ImageUrl:  "Test ImageUrl",
		Id:        "bfe05df6-eea1-11ee-9dc5-047c16a17206",
		OwnerId:   "4282f24b-ef6c-11ee-ae93-047c16a17206",
		Likes:     0,
		Views:     0,
		Category:  "Test Category",
		CreatedAt: "",
		UpdatedAt: "",
	}
	createdPost, err := s.Repository.Create(post)
	s.Suite.NotNil(createdPost)
	s.Suite.NoError(err)
	s.Suite.Equal(post.Title, createdPost.Title)
	s.Suite.Equal(post.Content, createdPost.Content)
	s.Suite.Equal(post.ImageUrl, createdPost.ImageUrl)
	s.Suite.Equal(post.Id, createdPost.Id)
	s.Suite.Equal(post.OwnerId, createdPost.OwnerId)
	s.Suite.Equal(post.Likes, createdPost.Likes)
	s.Suite.Equal(post.Views, createdPost.Views)
	s.Suite.Equal(post.Category, createdPost.Category)

	getPost, err := s.Repository.GetPost(createdPost.GetId())
	s.Suite.NotNil(getPost)
	s.Suite.NoError(err)

	createdPost.Title = "Update title"
	createdPost.Content = "Update lsat content"

	update, err := s.Repository.UpdatePost(createdPost)
	s.Suite.NotNil(update)
	s.Suite.NoError(err)
	s.Suite.Equal(post.Title, createdPost.Title)
	s.Suite.Equal(post.Content, createdPost.Content)
	s.Suite.Equal(post.ImageUrl, createdPost.ImageUrl)
	s.Suite.Equal(post.Id, createdPost.Id)
	s.Suite.Equal(post.OwnerId, createdPost.OwnerId)
	s.Suite.Equal(post.Likes, createdPost.Likes)
	s.Suite.Equal(post.Views, createdPost.Views)
	s.Suite.Equal(post.Category, createdPost.Category)

	allPosts, err := s.Repository.GetAllPosts(1, 10)
	s.Suite.NotNil(allPosts)
	s.Suite.NoError(err)

	resp, err := s.Repository.GetPostByOwnerId(createdPost.OwnerId)
	s.Suite.NoError(err)
	s.Suite.NotNil(resp)

	status, err := s.Repository.DeletePost(getPost.Id)
	s.Suite.NoError(err)
	s.Suite.True(status)
}

func (s *PostRepositoryTestSuite) TearDownSuity() {
	s.CleanUpFunc()
}

func TestPostRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(PostRepositoryTestSuite))
}
