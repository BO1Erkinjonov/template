package repo

import (
	pb "user_service/genproto/user-service"
)

type UserStorageI interface {
	Create(user *pb.User) (*pb.User, error)
	GetUser(id string) (*pb.User, error)
	GetAll(page, limit int64) (users []*pb.User, err error)
	Update(user *pb.User) (*pb.User, error)
	Delete(user_id string) error
	CheckUniquess(req *pb.CheckUniqReq) (int32, error)
	Exists(req string) (*pb.User, error)
}
