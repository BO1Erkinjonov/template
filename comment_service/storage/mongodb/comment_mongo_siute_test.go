package mongodb

import (
	"comment_service/config"
	pb "comment_service/genproto/comment-service"
	"comment_service/pkg/db"
	"comment_service/storage/repo"
	"github.com/stretchr/testify/suite"
	"testing"
)

type CommentRepositoryTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.CommentStorageI
}

func (s *CommentRepositoryTestSuite) SetupSuite() {
	pgPool, cleanUp := db.ConnectToMongoDBForSuite(config.Load())
	s.Repository = NewCommentRepo(pgPool)
	s.CleanUpFunc = cleanUp
}

func (s *CommentRepositoryTestSuite) TestCommentCRUD() {
	comment := &pb.Comment{
		Id:          "fd5bb017-0566-471f-9241-bebaa97af860",
		Description: "Test description",
		PostId:      "bfe05df6-eea1-11ee-9dc5-047c16a17206",
		OwnerId:     "dc118cba-ee91-11ee-ac7c-047c16a17206",
		CreatedAt:   "",
		UpdatedAt:   "",
	}
	createdComment, err := s.Repository.Create(comment)
	s.Suite.NotNil(createdComment)
	s.Suite.NoError(err)
	s.Suite.Equal(comment.Id, createdComment.Id)
	s.Suite.Equal(comment.Description, createdComment.Description)
	s.Suite.Equal(comment.PostId, createdComment.PostId)
	s.Suite.Equal(comment.OwnerId, createdComment.OwnerId)

	getComment, err := s.Repository.GetComment(createdComment.GetId())
	s.Suite.NotNil(getComment)
	s.Suite.NoError(err)

	createdComment.Description = "Update title"

	update, err := s.Repository.UpdateComment(createdComment)
	s.Suite.NotNil(update)
	s.Suite.NoError(err)
	s.Suite.Equal(comment.Id, createdComment.Id)
	s.Suite.Equal(comment.Description, createdComment.Description)
	s.Suite.Equal(comment.PostId, createdComment.PostId)
	s.Suite.Equal(comment.OwnerId, createdComment.OwnerId)

	allComments, err := s.Repository.GetAllComment(1, 10)
	s.Suite.NotNil(allComments)
	s.Suite.NoError(err)

	resp, err := s.Repository.GetCommentByPostId(createdComment.PostId)
	s.Suite.NoError(err)
	s.Suite.NotNil(resp)

	status, err := s.Repository.DeleteComment(getComment.Id)
	s.Suite.NoError(err)
	s.Suite.True(status)

}

func (s *CommentRepositoryTestSuite) TearDownSuity() {
	s.CleanUpFunc()
}

func TestCommentRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(CommentRepositoryTestSuite))
}
