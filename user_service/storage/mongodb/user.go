package mongodb

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "user_service/genproto/user-service"
)

type UserRepo struct {
	collection *mongo.Collection
}

func NewUserServiceMongo(collectionName *mongo.Collection) *UserRepo {
	return &UserRepo{
		collection: collectionName,
	}
}

func (u *UserRepo) Create(reg *pb.User) (*pb.User, error) {
	reg.CreatedAt = time.Now().String()
	reg.UpdatedAt = time.Now().String()

	_, err := u.collection.InsertOne(context.Background(), reg)
	if err != nil {
		return nil, err
	}

	return reg, nil
}

func (u *UserRepo) GetUser(id string) (*pb.User, error) {
	var user pb.User
	fmt.Println(id)
	err := u.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *UserRepo) GetAll(page, limit int64) (users []*pb.User, err error) {
	options := options.Find()
	options.SetLimit(limit)
	options.SetSkip(limit * (page - 1))

	cursor, err := c.collection.Find(context.Background(), bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var comment pb.User
		if err := cursor.Decode(&comment); err != nil {
			return nil, err
		}
		users = append(users, &comment)
	}
	return users, nil
}

func (u *UserRepo) Update(req *pb.User) (*pb.User, error) {
	update := bson.M{
		"$set": bson.M{
			"firstname":    req.FirstName,
			"lastname":     req.LastName,
			"username":     req.Username,
			"password":     req.Password,
			"email":        req.Email,
			"refreshtoken": req.RefreshToken,
			"updatedat":    time.Now().String(),
		},
	}
	filter := bson.M{"id": req.Id}
	updateResult, err := u.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	if updateResult.ModifiedCount == 0 {
		return nil, fmt.Errorf("user with id %v not found", req.Id)
	}
	return req, nil
}

func (u *UserRepo) Delete(id string) error {
	_, err := u.collection.DeleteOne(context.Background(), bson.M{"id": id})
	return err
}

func (u *UserRepo) CheckUniquess(reg *pb.CheckUniqReq) (int32, error) {
	count, err := u.collection.CountDocuments(context.Background(), bson.M{reg.Field: reg.Value}, &options.CountOptions{})
	if err != nil {
		return 0, err
	}

	if count != 0 {
		return 0, nil
	}

	n := rand.Int31() % 1000000
	return n, nil
}

func (u *UserRepo) Exists(email string) (*pb.User, error) {
	var user pb.User
	err := u.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
