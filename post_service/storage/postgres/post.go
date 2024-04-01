package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pb "post_service/genproto/post-service"
	"time"
)

type postRepo struct {
	db *sqlx.DB
}

func (u *postRepo) Create(post *pb.Post) (*pb.Post, error) {
	post.Id = uuid.NewString()
	err := u.db.QueryRow(`INSERT INTO posts(title, 
                  content, 
                  image_url,
                  id, 
                  owner_id,
                  likes,
                  views, 
                  category,
                  created_at, 
                  updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING title,
                   content, 
                   image_url, 
                   id,
                   owner_id,
                   likes, 
                   views, 
                   category, 
 				   created_at`,
		post.Title,
		post.Content,
		post.ImageUrl,
		post.Id,
		post.OwnerId,
		post.Likes,
		post.Views,
		post.Category,
		time.Now(),
		time.Now()).Scan(
		&post.Title,
		&post.Content,
		&post.ImageUrl,
		&post.Id,
		&post.OwnerId,
		&post.Likes,
		&post.Views,
		&post.Category,
		&post.CreatedAt,
	)
	fmt.Println(post.OwnerId)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *postRepo) GetPost(id string) (*pb.Post, error) {
	var post pb.Post
	err := u.db.QueryRow(`SELECT title, 
                  content, 
                  image_url,
                  id, 
                  owner_id,
                  likes,
                  views, 
                  category,
                  created_at, 
                  updated_at FROM posts WHERE id = $1`, id).Scan(&post.Title,
		&post.Content,
		&post.ImageUrl,
		&post.Id,
		&post.OwnerId,
		&post.Likes,
		&post.Views,
		&post.Category,
		&post.CreatedAt,
		&post.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (u *postRepo) GetAllPosts(page, limit int32) (posts []*pb.Post, err error) {
	offset := limit * (page - 1)
	rows, err := u.db.Query(`SELECT * FROM posts LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post pb.Post
		err = rows.Scan(&post.Title,
			&post.Content,
			&post.ImageUrl,
			&post.Id,
			&post.OwnerId,
			&post.Likes,
			&post.Views,
			&post.Category,
			&post.CreatedAt,
			&post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (u *postRepo) UpdatePost(post *pb.Post) (*pb.Post, error) {
	_, err := u.db.Exec(`UPDATE posts SET likes = $1, views = $2, category = $3, updated_at = now() WHERE id = $4 `, post.Likes, post.Views, post.Category, post.Id)
	if err != nil {
		return nil, err
	}
	upPost, err := u.GetPost(post.Id)
	if err != nil {
		return nil, err
	}
	return upPost, nil
}

func (u *postRepo) DeletePost(id string) (bool, error) {
	_, err := u.db.Exec(`DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *postRepo) GetPostByOwnerId(id string) ([]*pb.Post, error) {
	rows, err := u.db.Query(`SELECT title, 
                  content, 
                  image_url,
                  id, 
                  owner_id,
                  likes,
                  views, 
                  category,
                  created_at, 
                  updated_at FROM posts WHERE owner_id = $1`, id)
	if err != nil {
		return nil, err
	}
	var posts []*pb.Post
	for rows.Next() {
		var post pb.Post
		if err = rows.Scan(&post.Title,
			&post.Content,
			&post.ImageUrl,
			&post.Id,
			&post.OwnerId,
			&post.Likes,
			&post.Views,
			&post.Category,
			&post.CreatedAt,
			&post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (u *postRepo) DeletePostByOwnerId(id string) (bool, error) {
	_, err := u.db.Exec(`DELETE FROM posts WHERE owner_id = $1`, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func NewPostRepo(db *sqlx.DB) *postRepo {
	return &postRepo{db: db}
}
