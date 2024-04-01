package mongodb

import (
	"context"
	"strconv"
	"time"

	pbc "comment_service/genproto/comment-service"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type commentRepo struct {
	collection *mongo.Collection
}

func (c *commentRepo) Create(comment *pbc.Comment) (*pbc.Comment, error) {
	comment.Id = uuid.New().String()
	comment.CreatedAt = strconv.FormatInt(time.Now().Unix(), 10)
	comment.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)

	_, err := c.collection.InsertOne(context.Background(), comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (c *commentRepo) GetComment(id string) (*pbc.Comment, error) {
	var comment pbc.Comment
	err := c.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&comment)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (c *commentRepo) GetAllComment(page, limit int64) ([]*pbc.Comment, error) {
	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(limit * (page - 1))

	cursor, err := c.collection.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var comments []*pbc.Comment
	for cursor.Next(context.Background()) {
		var comment pbc.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (c *commentRepo) UpdateComment(req *pbc.Comment) (*pbc.Comment, error) {
	req.UpdatedAt = strconv.FormatInt(time.Now().Unix(), 10)

	filter := bson.M{"id": req.Id}
	update := bson.M{"$set": req}

	_, err := c.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *commentRepo) DeleteComment(id string) (bool, error) {
	filter := bson.M{"id": id}

	_, err := c.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *commentRepo) GetCommentByPostId(id string) ([]*pbc.Comment, error) {
	cursor, err := c.collection.Find(context.Background(), bson.M{"postid": id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var comments []*pbc.Comment
	for cursor.Next(context.Background()) {
		var comment pbc.Comment
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func NewCommentRepo(database *mongo.Collection) *commentRepo {
	return &commentRepo{
		collection: database,
	}
}
