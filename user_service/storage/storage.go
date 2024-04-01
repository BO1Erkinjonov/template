package storage

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"user_service/storage/mongodb"
	"user_service/storage/postgres"
	"user_service/storage/repo"
)

type IStorage interface {
	User() repo.UserStorageI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.UserStorageI
}

type storageMongo struct {
	db       *mongo.Collection
	userRepo repo.UserStorageI
}

func (s storageMongo) User() repo.UserStorageI {
	return s.userRepo
}

func (s storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{db, postgres.NewUserRepo(db)}
}
func NewStorageMongo(db *mongo.Collection) *storageMongo {
	return &storageMongo{db, mongodb.NewUserServiceMongo(db)}
}
