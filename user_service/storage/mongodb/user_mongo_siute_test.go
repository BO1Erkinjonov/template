package mongodb

import (
	"github.com/stretchr/testify/suite"
	"testing"
	"user_service/config"
	pb "user_service/genproto/user-service"
	"user_service/pkg/db"
	"user_service/storage/repo"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  repo.UserStorageI
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	pgPool, cleanUp := db.ConnectToMongoDBForSuite(config.Load())
	s.Repository = NewUserServiceMongo(pgPool)
	s.CleanUpFunc = cleanUp
}

func (s *UserRepositoryTestSuite) TestUserCRUD() {
	user := &pb.User{
		FirstName:    "Test first_name",
		LastName:     "Test last_name",
		Username:     "Test user_name",
		Password:     "Test password",
		Email:        "Test email",
		Id:           "7fbdba15-ee94-11ee-956a-047c16a17206",
		RefreshToken: "Test tocekr",
		CreatedAt:    "",
		UpdatedAt:    "",
	}
	createdUser, err := s.Repository.Create(user)
	s.Suite.NotNil(createdUser)
	s.Suite.NoError(err)
	s.Suite.Equal(user.FirstName, createdUser.FirstName)
	s.Suite.Equal(user.LastName, createdUser.LastName)
	s.Suite.Equal(user.Username, createdUser.Username)
	s.Suite.Equal(user.Email, createdUser.Email)
	s.Suite.Equal(user.Id, createdUser.Id)
	s.Suite.Equal(user.RefreshToken, createdUser.RefreshToken)

	getUser, err := s.Repository.GetUser(createdUser.GetId())
	s.Suite.NotNil(getUser)
	s.Suite.NoError(err)

	createdUser.FirstName = "Update name"
	createdUser.LastName = "Update lsat name"

	update, err := s.Repository.Update(createdUser)
	s.Suite.NotNil(update)
	s.Suite.NoError(err)
	s.Suite.Equal(update.FirstName, createdUser.FirstName)
	s.Suite.Equal(update.LastName, createdUser.LastName)
	s.Suite.Equal(update.Username, createdUser.Username)
	s.Suite.Equal(update.Email, createdUser.Email)
	s.Suite.Equal(user.Id, createdUser.Id)
	s.Suite.Equal(user.RefreshToken, createdUser.RefreshToken)

	allUsers, err := s.Repository.GetAll(1, 10)
	s.Suite.NotNil(allUsers)
	s.Suite.NoError(err)

	check := &pb.CheckUniqReq{
		Field: "email",
		Value: createdUser.Email,
	}

	n, err := s.Repository.CheckUniquess(check)
	s.Suite.NoError(err)
	s.Suite.NotNil(n)

	exists, err := s.Repository.Exists(createdUser.Email)
	s.Suite.NoError(err)
	s.Suite.NotNil(exists)
	s.Suite.Equal(update.FirstName, createdUser.FirstName)
	s.Suite.Equal(update.LastName, createdUser.LastName)
	s.Suite.Equal(update.Username, createdUser.Username)
	s.Suite.Equal(update.Email, createdUser.Email)
	s.Suite.Equal(user.Id, createdUser.Id)
	s.Suite.Equal(user.RefreshToken, createdUser.RefreshToken)

	err = s.Repository.Delete(getUser.Id)
	s.Suite.NoError(err)

}

func (s *UserRepositoryTestSuite) TearDownSuity() {
	s.CleanUpFunc()
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
