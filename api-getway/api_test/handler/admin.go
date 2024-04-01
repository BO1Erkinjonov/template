package handler

import (
	"api-test/api/handlers/models"
	"api-test/config"
	compb "api-test/genproto/comment-service"
	postpb "api-test/genproto/post-service"
	pb "api-test/genproto/user-service"
	"api-test/mock_data/comment_service"
	"api-test/mock_data/post_service"
	"api-test/mock_data/user_service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

func AdminUpdatePost(c *gin.Context) {
	id := c.Query("post_id")
	var (
		body        models.AdminPost
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()
	post, err := post_service.NewMockPostServiceClient().GetPost(ctx, &postpb.GetRequests{PostId: id})
	fmt.Println(post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	response, err := post_service.NewMockPostServiceClient().UpdatePost(ctx, &postpb.Post{
		Title:     post.Title,
		Content:   post.Content,
		ImageUrl:  post.ImageUrl,
		Id:        id,
		Likes:     post.Likes,
		Views:     post.Views,
		Category:  post.Category,
		CreatedAt: post.CreatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func AdminUpdateUser(c *gin.Context) {
	var (
		body        models.User
		jspbMarshal protojson.MarshalOptions
	)
	jspbMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	response, err := user_service.NewMockUserServiceClient().Update(ctx, &pb.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Username:     body.UserName,
		Password:     body.Password,
		Email:        body.Email,
		RefreshToken: "Mock token",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func AdminDeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	user2, err := user_service.NewMockUserServiceClient().GetUser(ctx, &pb.GetRequest{UserId: guid})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	if user2.Role == "superAdmin" || user2.Role == "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "permission denied1",
		})
		return
	}

	response, err := user_service.NewMockUserServiceClient().Delete(
		ctx, &pb.GetRequest{
			UserId: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = post_service.NewMockPostServiceClient().DeletePostByOwnerId(ctx, &postpb.GetByOwnerIdRequest{
		OwnerId: user2.Id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

func SuperAdminDeleteAdmin(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	response, err := user_service.NewMockUserServiceClient().Delete(
		ctx, &pb.GetRequest{
			UserId: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func AdminDeleteComment(c *gin.Context) {
	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	s, err := comment_service.NewMockCommentServiceClient().DeleteComment(ctx, &compb.Get{
		Id: guid,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, s)
}

func Create(c *gin.Context) {

	var (
		body        models.AdminUser
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	exists_email, err := user_service.NewMockUserServiceClient().CheckUniquess(ctx, &pb.CheckUniqReq{
		Field: "email",
		Value: body.Email,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if exists_email.Code == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This email already in use, please use another email address",
		})
		return
	}

	exists, err := user_service.NewMockUserServiceClient().CheckUniquess(ctx, &pb.CheckUniqReq{
		Field: "username",
		Value: body.UserName,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if exists.Code == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This username already in use, please use another username address",
		})
		return
	}

	Id := uuid.NewString()

	response, err := user_service.NewMockUserServiceClient().Create(ctx, &pb.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Username:     body.UserName,
		Role:         body.Role,
		Password:     body.Password,
		Email:        body.Email,
		Id:           Id,
		RefreshToken: "Mock token",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response)
}
