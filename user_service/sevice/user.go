package sevice

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	commentPb "user_service/genproto/comment-service"
	postPb "user_service/genproto/post-service"
	userPb "user_service/genproto/user-service"
	l "user_service/pkg/logger"
	"user_service/sevice/grpc_client"
	"user_service/storage"
)

type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpc_client.IServiceManager
}

func (u *UserService) Create(ctx context.Context, req *userPb.User) (*userPb.User, error) {

	user, err := u.storage.User().Create(req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return &userPb.User{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Role:         user.Role,
		Password:     user.Password,
		Email:        user.Email,
		Id:           user.Id,
		RefreshToken: user.RefreshToken,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}, nil
}

func (u *UserService) GetUser(ctx context.Context, req *userPb.GetRequest) (*userPb.User, error) {
	user, err := u.storage.User().GetUser(req.UserId)

	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	posts, err := u.client.PostService().GetPostByOwnerId(ctx, &postPb.GetByOwnerIdRequest{
		OwnerId: user.Id,
	})
	if err != nil {
		return nil, err
	}
	for _, post := range posts.Posts {
		comments, err := u.client.CommentService().GetCommentByPostId(ctx, &commentPb.GetCommentByPostIdRequest{PostId: post.Id})
		if err != nil {
			return nil, err
		}
		var uPost userPb.Post
		for _, com := range comments.Comments {
			var pCom userPb.Comment
			pCom.Id = com.Id
			pCom.Description = com.Description
			uPost.Comment = append(uPost.Comment, &pCom)
		}
		uPost.Title = post.Title
		uPost.Content = post.Content
		uPost.ImageUrl = post.ImageUrl
		uPost.Id = post.Id
		uPost.OwnerId = post.OwnerId
		uPost.Likes = post.Likes
		uPost.Views = post.Views
		user.Post = append(user.Post, &uPost)
	}
	return user, nil
}

func (u *UserService) Update(ctx context.Context, req *userPb.User) (*userPb.User, error) {
	user, err := u.storage.User().Update(req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (u *UserService) GetAllUsers(ctx context.Context, req *userPb.GetAllUsersRequest) (*userPb.GetAllUsersResponse, error) {
	users, err := u.storage.User().GetAll(req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	return &userPb.GetAllUsersResponse{Users: users}, nil
}

func (u *UserService) Delete(ctx context.Context, req *userPb.GetRequest) (*userPb.Tf, error) {
	err := u.storage.User().Delete(req.UserId)
	if err != nil {
		u.logger.Error(err.Error())
		return &userPb.Tf{
			Tf: false,
		}, err
	}
	return &userPb.Tf{
		Tf: true,
	}, nil
}

func (u *UserService) CheckUniquess(ctx context.Context, req *userPb.CheckUniqReq) (*userPb.CheckUniqResp, error) {
	code, err := u.storage.User().CheckUniquess(req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return &userPb.CheckUniqResp{
		Code: code,
	}, nil
}

func (u *UserService) Exists(ctx context.Context, req *userPb.Req) (*userPb.User, error) {
	user, err := u.storage.User().Exists(req.Email)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func NewUserService(db *sqlx.DB, log l.Logger, client grpc_client.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func NewUserServiceMongo(db *mongo.Collection, log l.Logger, client grpc_client.IServiceManager) *UserService {
	return &UserService{
		storage: storage.NewStorageMongo(db),
		logger:  log,
		client:  client,
	}
}
