package service

import (
	pbc "comment_service/genproto/comment-service"
	pbu "comment_service/genproto/post-service"
	pb "comment_service/genproto/user-service"
	l "comment_service/pkg/logger"
	"comment_service/service/grpc_client"
	"comment_service/storage"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type CommentService struct {
	storage storage.IStorage
	logger  l.Logger
	client  grpc_client.IServiceManager
}

func (u *CommentService) Create(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	req.Id = id.String()
	user, err := u.storage.Comment().Create(req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (u *CommentService) GetComment(ctx context.Context, id *pbc.Get) (*pbc.Comment, error) {
	getReq, err := u.storage.Comment().GetComment(id.Id)
	if err != nil {
		return nil, err
	}

	reqUser, err := u.client.UserService().GetUser(ctx, &pb.GetRequest{UserId: getReq.OwnerId})
	if err != nil {
		return nil, err
	}
	reqPost, err := u.client.PostService().GetPost(ctx, &pbu.GetRequests{PostId: getReq.PostId})
	if err != nil {
		return nil, err
	}

	return &pbc.Comment{
		Id:          getReq.Id,
		Description: getReq.Description,
		PostId:      getReq.PostId,
		OwnerId:     getReq.OwnerId,
		CreatedAt:   getReq.CreatedAt,
		UpdatedAt:   getReq.UpdatedAt,
		User: &pbc.User{
			FirstName:    reqUser.FirstName,
			LastName:     reqUser.LastName,
			Username:     reqUser.Username,
			Role:         reqUser.Role,
			Password:     reqUser.Password,
			Email:        reqUser.Email,
			Id:           reqUser.Id,
			RefreshToken: reqUser.RefreshToken,
			CreatedAt:    reqUser.CreatedAt,
			UpdatedAt:    reqUser.UpdatedAt,
		},
		Post: &pbc.Post{
			Title:     reqPost.Title,
			Content:   reqPost.Content,
			ImageUrl:  reqPost.ImageUrl,
			Id:        reqPost.Id,
			Likes:     reqPost.Likes,
			Views:     reqPost.Views,
			Category:  reqPost.Category,
			CreatedAt: reqPost.CreatedAt,
			UpdatedAt: reqPost.UpdatedAt,
		},
	}, nil
}

func (u *CommentService) GetAllComment(ctx context.Context, req *pbc.GetRequest) (*pbc.GetResponse, error) {
	commants, err := u.storage.Comment().GetAllComment(req.Page, req.Limit)
	if err != nil {
		return nil, err
	}
	return &pbc.GetResponse{
		Comments: commants,
	}, nil
}

func (u *CommentService) UpdateComment(ctx context.Context, req *pbc.Comment) (*pbc.Comment, error) {
	resp, err := u.storage.Comment().UpdateComment(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *CommentService) DeleteComment(ctx context.Context, id *pbc.Get) (*pbc.Tf, error) {
	status, err := u.storage.Comment().DeleteComment(id.Id)
	if err != nil {
		return &pbc.Tf{
			Tf: status,
		}, err
	}
	return &pbc.Tf{
		Tf: status,
	}, err
}

func (u *CommentService) GetCommentByPostId(ctx context.Context, req *pbc.GetCommentByPostIdRequest) (*pbc.GetCommentByPostIdResponse, error) {
	comment, err := u.storage.Comment().GetCommentByPostId(req.PostId)
	if err != nil {
		return nil, err
	}
	return &pbc.GetCommentByPostIdResponse{
		Comments: comment,
	}, nil
}

func NewCommentService(db *sqlx.DB, log l.Logger, client grpc_client.IServiceManager) *CommentService {
	return &CommentService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		client:  client,
	}
}

func NewCommentServiceMongo(db *mongo.Collection, log l.Logger, client grpc_client.IServiceManager) *CommentService {
	return &CommentService{
		storage: storage.NewStorageMongo(db),
		logger:  log,
		client:  client,
	}
}
