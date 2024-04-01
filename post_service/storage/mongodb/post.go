package mongodb

import (
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	pb "post_service/genproto/post-service"
)

type postRepo struct {
	collection *mongo.Collection
}

func (p *postRepo) Create(post *pb.Post) (*pb.Post, error) {
	post.Id = uuid.NewString()
	post.CreatedAt = time.Now().String()
	post.UpdatedAt = time.Now().String()
	_, err := p.collection.InsertOne(context.Background(), post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postRepo) GetPost(id string) (*pb.Post, error) {
	var post pb.Post
	err := p.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (p *postRepo) GetAllPosts(page, limit int32) ([]*pb.Post, error) {
	options := options.Find()
	options.SetLimit(int64(limit))
	options.SetSkip(int64(limit * (page - 1)))

	cursor, err := p.collection.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts []*pb.Post
	for cursor.Next(context.Background()) {
		var post pb.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (p *postRepo) UpdatePost(post *pb.Post) (*pb.Post, error) {
	post.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	filter := bson.M{"id": post.Id}
	update := bson.M{
		"$set": bson.M{
			"title":     post.Title,
			"content":   post.Content,
			"imageurl":  post.Content,
			"id":        post.Id,
			"ownerid":   post.OwnerId,
			"likes":     post.Likes,
			"views":     post.Views,
			"category":  post.Category,
			"createdat": post.CreatedAt,
			"updatedat": time.Now().String(),
		},
	}
	_, err := p.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (p *postRepo) DeletePost(id string) (bool, error) {
	filter := bson.M{"id": id}

	_, err := p.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (p *postRepo) GetPostByOwnerId(id string) ([]*pb.Post, error) {
	cursor, err := p.collection.Find(context.Background(), bson.M{"ownerid": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var posts []*pb.Post
	for cursor.Next(context.Background()) {
		var post pb.Post
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (p *postRepo) DeletePostByOwnerId(id string) (bool, error) {
	filter := bson.M{"ownerid": id}

	_, err := p.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return true, nil
}

func NewPostRepoMongo(database *mongo.Collection) *postRepo {
	return &postRepo{
		collection: database,
	}
}
