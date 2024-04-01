package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
	pbc "post_service/genproto/comment-service"
	pb "post_service/genproto/post-service"
	pbu "post_service/genproto/user-service"
	l "post_service/pkg/logger"
	"post_service/service/grpc_client"
	"post_service/storage"
)

type PostService struct {
	storage    storage.IStorage
	logger     l.Logger
	grpcClient grpc_client.IServiceManager
}

func (u *PostService) Create(ctx context.Context, req *pb.Post) (*pb.Post, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	req.Id = id.String()
	post, err := u.storage.Post().Create(req)
	if err != nil {
		u.logger.Error(err.Error())
		return nil, err
	}
	return post, nil
}

func (u *PostService) GetPost(ctx context.Context, id *pb.GetRequests) (*pb.PostResponse, error) {
	post, err := u.storage.Post().GetPost(id.PostId)

	if err != nil {
		return nil, err
	}
	user, err := u.grpcClient.UserService().GetUser(ctx, &pbu.GetRequest{
		UserId: post.OwnerId,
	})
	if err != nil {
		return nil, err
	}

	comment, err := u.grpcClient.CommentService().GetCommentByPostId(ctx, &pbc.GetCommentByPostIdRequest{
		PostId: post.Id,
	})
	if err != nil {
		return nil, err
	}
	for _, com := range comment.Comments {
		var pComment pb.Comment
		pComment.CreatedAt = com.CreatedAt
		pComment.UpdatedAt = com.UpdatedAt
		pComment.Id = com.Id
		pComment.Description = com.Description
		pComment.PostId = com.PostId
		pComment.OwnerId = com.OwnerId
		post.Comment = append(post.Comment, &pComment)
	}

	return &pb.PostResponse{
		Title:     post.Title,
		Content:   post.Content,
		ImageUrl:  post.ImageUrl,
		Id:        post.Id,
		Likes:     post.Likes,
		Views:     post.Views,
		Category:  post.Category,
		CreatedAt: post.CreatedAt,
		Owner: &pb.Owner{
			Id:       user.Id,
			Name:     user.FirstName,
			LastName: user.LastName,
		},
		Comment: post.Comment,
	}, nil
}

func (u *PostService) GetAllPost(ctx context.Context, req *pb.GetAllPostRequest) (*pb.GetAllPostResponse, error) {
	posts, err := u.storage.Post().GetAllPosts(int32(req.Page), int32(req.Limit))
	for _, post := range posts {
		user, err := u.grpcClient.UserService().GetUser(ctx, &pbu.GetRequest{
			UserId: post.OwnerId,
		})
		if err != nil {

			return nil, err
		}

		comment, err := u.grpcClient.CommentService().GetCommentByPostId(ctx, &pbc.GetCommentByPostIdRequest{
			PostId: post.Id,
		})
		if err != nil {
			return nil, err
		}
		for _, com := range comment.Comments {
			var pComment pb.Comment
			pComment.CreatedAt = com.CreatedAt
			pComment.UpdatedAt = com.UpdatedAt
			pComment.Id = com.Id
			pComment.Description = com.Description
			pComment.PostId = com.PostId
			pComment.OwnerId = com.OwnerId
			post.Comment = append(post.Comment, &pComment)
		}
		var pos pb.Owner
		pos.Id = user.Id
		pos.Name = user.FirstName
		pos.LastName = user.LastName
		post.Owner = &pos
	}
	if err != nil {
		return nil, err
	}
	return &pb.GetAllPostResponse{
		Posts: posts,
	}, nil
}

func (u *PostService) UpdatePost(ctx context.Context, post *pb.Post) (*pb.Post, error) {
	uppost, err := u.storage.Post().UpdatePost(post)
	if err != nil {
		return nil, err
	}
	return uppost, nil
}

func (u *PostService) DeletePost(ctx context.Context, id *pb.GetRequests) (*pb.Tf, error) {
	resp, err := u.storage.Post().DeletePost(id.PostId)
	if err != nil {
		return &pb.Tf{
			Tf: false,
		}, err
	}
	return &pb.Tf{
		Tf: resp,
	}, nil
}

func (u *PostService) GetPostByOwnerId(ctx context.Context, id *pb.GetByOwnerIdRequest) (*pb.GetByOwnerIdResponse, error) {
	posts, err := u.storage.Post().GetPostByOwnerId(id.OwnerId)
	if err != nil {
		return nil, err
	}
	for _, post := range posts {
		comment, err := u.grpcClient.CommentService().GetCommentByPostId(ctx, &pbc.GetCommentByPostIdRequest{
			PostId: post.Id,
		})
		if err != nil {
			return nil, err
		}

		for _, comment := range comment.Comments {
			var pComment pb.Comment
			pComment.CreatedAt = comment.CreatedAt
			pComment.UpdatedAt = comment.UpdatedAt
			pComment.Id = comment.Id
			pComment.PostId = comment.PostId
			pComment.OwnerId = comment.OwnerId
			post.Comment = append(post.Comment, &pComment)
		}
	}
	return &pb.GetByOwnerIdResponse{
		Posts: posts,
	}, nil

}

func (u *PostService) DeletePostByOwnerId(ctx context.Context, id *pb.GetByOwnerIdRequest) (*pb.Tf, error) {
	status, err := u.storage.Post().DeletePostByOwnerId(id.OwnerId)
	if err != nil {
		return nil, err
	}
	return &pb.Tf{Tf: status}, nil
}

func NewPostService(db *sqlx.DB, log l.Logger, grpcClient grpc_client.IServiceManager) *PostService {
	return &PostService{
		storage:    storage.NewStoragePg(db),
		logger:     log,
		grpcClient: grpcClient,
	}
}

func NewPostServiceMongo(db *mongo.Collection, log l.Logger, grpcClient grpc_client.IServiceManager) *PostService {
	return &PostService{
		storage:    storage.NewStorageMongo(db),
		logger:     log,
		grpcClient: grpcClient,
	}
}
