package storage

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"post_service/storage/mongodb"
	"post_service/storage/postgres"
	"post_service/storage/repo"
)

type IStorage interface {
	Post() repo.PostStorageI
}

type storagePg struct {
	db       *sqlx.DB
	userRepo repo.PostStorageI
}

type storageMongo struct {
	db       *mongo.Collection
	userRepo repo.PostStorageI
}

func (s storagePg) Post() repo.PostStorageI {
	return s.userRepo
}
func (s storageMongo) Post() repo.PostStorageI {
	return s.userRepo
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{db, postgres.NewPostRepo(db)}
}

func NewStorageMongo(db *mongo.Collection) *storageMongo {
	return &storageMongo{db, mongodb.NewPostRepoMongo(db)}
}
