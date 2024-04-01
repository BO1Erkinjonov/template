package handler

import (
	"api-test/api/handlers/models"
	"api-test/config"
	compb "api-test/genproto/comment-service"
	pb "api-test/genproto/post-service"
	"api-test/mock_data/comment_service"
	"api-test/mock_data/post_service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	com, err := comment_service.NewMockCommentServiceClient().Create(ctx, &compb.Comment{
		Description: comment.Description,
		PostId:      comment.PostId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, com)
}
func GetComment(c *gin.Context) {
	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()
	com, err := comment_service.NewMockCommentServiceClient().GetComment(ctx, &compb.Get{
		Id: guid,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, com)
}
func GetAllComment(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()
	rows, err := comment_service.NewMockCommentServiceClient().GetAllComment(ctx, &compb.GetRequest{
		Page:  int64(1),
		Limit: int64(10),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, rows)
}
func UpdateComment(c *gin.Context) {
	var comment models.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	com, err := comment_service.NewMockCommentServiceClient().UpdateComment(ctx, &compb.Comment{
		Id:          comment.Id,
		Description: comment.Description,
		PostId:      comment.PostId,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(com.CreatedAt)
	c.JSON(http.StatusOK, com)
}
func DeleteComment(c *gin.Context) {
	guid := c.Query("id")
	post_id := c.Query("post_id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	_, err := post_service.NewMockPostServiceClient().GetPost(ctx, &pb.GetRequests{
		PostId: post_id,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

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
