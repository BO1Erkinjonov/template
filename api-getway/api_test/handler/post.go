package handler

import (
	"api-test/api/handlers/models"
	jwt "api-test/api/tokens"
	"api-test/config"
	postpb "api-test/genproto/post-service"
	"api-test/mock_data/post_service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

func CreatePost(c *gin.Context) {
	var (
		body        models.Post
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	tok := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(tok, []byte(config.Load().SignInKey))
	response, err := post_service.NewMockPostServiceClient().Create(ctx, &postpb.Post{
		Title:    body.Title,
		Category: body.Category,
		Content:  body.Content,
		ImageUrl: body.ImageUrl,
		OwnerId:  cast.ToString(claims["sub"]),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetPost(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	post, err := post_service.NewMockPostServiceClient().GetPost(ctx, &postpb.GetRequests{
		PostId: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	_, err = post_service.NewMockPostServiceClient().UpdatePost(ctx, &postpb.Post{
		Id:        post.Id,
		Views:     post.Views + 1,
		Likes:     post.Likes,
		Title:     post.Title,
		Category:  post.Category,
		Content:   post.Content,
		ImageUrl:  post.ImageUrl,
		CreatedAt: post.CreatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, post)
}

func GetAllPosts(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()
	resp, err := post_service.NewMockPostServiceClient().GetAllPost(ctx, &postpb.GetAllPostRequest{
		Page:  int64(1),
		Limit: int64(10),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, resp)
}

func UpdatePost(c *gin.Context) {
	id := c.Query("id")
	var (
		body        models.Post
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

	response, err := post_service.NewMockPostServiceClient().UpdatePost(ctx, &postpb.Post{
		Title:    body.Title,
		Category: body.Category,
		Content:  body.Content,
		ImageUrl: body.ImageUrl,
		Id:       id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func DeletePost(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	response, err := post_service.NewMockPostServiceClient().DeletePost(ctx, &postpb.GetRequests{PostId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response)
}

func Like(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()
	post, err := post_service.NewMockPostServiceClient().GetPost(ctx, &postpb.GetRequests{
		PostId: id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	presp, err := post_service.NewMockPostServiceClient().UpdatePost(ctx, &postpb.Post{
		Id:        post.Id,
		Views:     post.Views,
		Likes:     post.Likes + 1,
		Title:     post.Title,
		Category:  post.Category,
		Content:   post.Content,
		ImageUrl:  post.ImageUrl,
		CreatedAt: post.CreatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, presp)
}

func DisLike(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()
	post, err := post_service.NewMockPostServiceClient().GetPost(ctx, &postpb.GetRequests{
		PostId: id,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if post.Likes == 0 {
		c.JSON(http.StatusOK, post)
		return
	}
	presp, err := post_service.NewMockPostServiceClient().UpdatePost(ctx, &postpb.Post{
		Id:        post.Id,
		Views:     post.Views,
		OwnerId:   post.Owner.Id,
		Likes:     post.Likes - 1,
		Title:     post.Title,
		Category:  post.Category,
		Content:   post.Content,
		ImageUrl:  post.ImageUrl,
		CreatedAt: post.CreatedAt,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, presp)
}
