package repo

import pbc "comment_service/genproto/comment-service"

type CommentStorageI interface {
	Create(comment *pbc.Comment) (*pbc.Comment, error)
	GetComment(id string) (*pbc.Comment, error)
	GetAllComment(page, limit int64) ([]*pbc.Comment, error)
	UpdateComment(comment *pbc.Comment) (*pbc.Comment, error)
	DeleteComment(id string) (bool, error)
	GetCommentByPostId(id string) ([]*pbc.Comment, error)
}
