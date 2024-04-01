package storage

import (
	"comment_service/storage/mongodb"
	"comment_service/storage/postgres"
	"comment_service/storage/repo"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type IStorage interface {
	Comment() repo.CommentStorageI
}

type storagePg struct {
	db          *sqlx.DB
	commentRepo repo.CommentStorageI
}

type storageMongo struct {
	db          *mongo.Collection
	commentRepo repo.CommentStorageI
}

func (s storageMongo) Comment() repo.CommentStorageI {
	return s.commentRepo
}

func (s storagePg) Comment() repo.CommentStorageI {
	return s.commentRepo
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{db, postgres.NewCommentRepo(db)}
}

func NewStorageMongo(db *mongo.Collection) *storageMongo {
	return &storageMongo{db, mongodb.NewCommentRepo(db)}
}
