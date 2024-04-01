package postgres

import (
	pbc "comment_service/genproto/comment-service"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type commentRepo struct {
	db *sqlx.DB
}

func (u *commentRepo) Create(comment *pbc.Comment) (*pbc.Comment, error) {
	comment.Id = uuid.NewString()
	query := `INSERT INTO comments(id, description, post_id, owner_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, description, post_id, owner_id, created_at, updated_at`
	err := u.db.QueryRow(query, comment.Id, comment.Description, comment.PostId, comment.OwnerId, time.Now(), time.Now()).Scan(&comment.Id, &comment.Description, &comment.PostId, &comment.OwnerId, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (u *commentRepo) GetComment(id string) (*pbc.Comment, error) {
	query := `SELECT * FROM comments WHERE id = $1`
	var comment pbc.Comment
	err := u.db.QueryRow(query, id).Scan(&comment.Id, &comment.Description, &comment.PostId, &comment.OwnerId, &comment.CreatedAt, &comment.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil

}

func (u *commentRepo) GetAllComment(page, limit int64) ([]*pbc.Comment, error) {
	var comments []*pbc.Comment
	offset := limit * (page - 1)
	rows, err := u.db.Query(`SELECT * FROM comments LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment pbc.Comment
		if err := rows.Scan(&comment.Id, &comment.Description, &comment.PostId, &comment.OwnerId, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (u *commentRepo) UpdateComment(req *pbc.Comment) (*pbc.Comment, error) {
	err := u.db.QueryRow(`UPDATE comments SET description = $1, updated_at = now() WHERE id = $2 RETURNING description, created_at, updated_at`, req.Description, req.Id).Scan(&req.Description, &req.CreatedAt, &req.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (u *commentRepo) DeleteComment(id string) (bool, error) {
	_, err := u.db.Exec(`DELETE FROM comments WHERE id = $1`, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *commentRepo) GetCommentByPostId(id string) ([]*pbc.Comment, error) {
	var comments []*pbc.Comment
	rows, err := u.db.Query(`SELECT * FROM comments WHERE post_id = $1`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment pbc.Comment
		if err := rows.Scan(&comment.Id, &comment.Description, &comment.PostId, &comment.OwnerId, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func NewCommentRepo(db *sqlx.DB) *commentRepo {
	return &commentRepo{db: db}
}
