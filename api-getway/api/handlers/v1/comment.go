package v1

import (
	"api-test/api/handlers/models"
	jwt "api-test/api/tokens"
	"api-test/config"
	compb "api-test/genproto/comment-service"
	pb "api-test/genproto/post-service"
	l "api-test/pkg/logger"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
	"time"
)

// CreateComment ...
// @Summary CreateComment
// @Security ApiKeyAuth
// @Description Api for creating a new comment
// @Tags comment
// @Accept json
// @Produce json
// @Param User body models.Comment true "create comment"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/create/comment [post]
func (h *handlerV1) CreateComment(c *gin.Context) {
	var comment models.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resp, err := h.serviceManager.PostService().GetPost(ctx, &pb.GetRequests{
		PostId: comment.PostId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to post id", l.Error(err))
		return
	}

	com, err := h.serviceManager.CommentService().Create(ctx, &compb.Comment{
		Description: comment.Description,
		PostId:      comment.PostId,
		OwnerId:     resp.Owner.Id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to created comment", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, com)
}

// GetComment ...
// @Summary GetComment
// @Security ApiKeyAuth
// @Tags comment
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/get/comment [get]
func (h *handlerV1) GetComment(c *gin.Context) {
	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	com, err := h.serviceManager.CommentService().GetComment(ctx, &compb.Get{
		Id: guid,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to query", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, com)
}

// GetAllComment gets user by id
// @Summary GetAllComment
// @Security ApiKeyAuth
// @Description Api for getting comments
// @Tags comment
// @Accept json
// @Produce json
// @Param page query string true "page"
// @Param limit query string true "limit"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/comments/all [get]
func (h *handlerV1) GetAllComment(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")

	reqPage, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("page error", l.Error(err))
		return
	}
	reqLimit, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("limit error", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	rows, err := h.serviceManager.CommentService().GetAllComment(ctx, &compb.GetRequest{
		Page:  int64(reqPage),
		Limit: int64(reqLimit),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to page or limit", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, rows)
}

// UpdateComment ...
// @Summary UpdateComment
// @Security ApiKeyAuth
// @Description Api for updating comment
// @Tags comment
// @Accept json
// @Produce json
// @Param User body models.Comment true "create comment"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/update/comment [put]
func (h *handlerV1) UpdateComment(c *gin.Context) {
	var comment models.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to bind json", l.Error(err))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	resp, err := h.serviceManager.PostService().GetPost(ctx, &pb.GetRequests{
		PostId: comment.PostId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to post id", l.Error(err))
		return
	}
	tok := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(tok, []byte(config.Load().SignInKey))

	if cast.ToString(claims["sub"]) != resp.Owner.Id {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "this is not your comment",
		})
		h.log.Error("this is not your comment", l.Error(err))
		return
	}
	com, err := h.serviceManager.CommentService().UpdateComment(ctx, &compb.Comment{
		Id:          comment.Id,
		Description: comment.Description,
		PostId:      comment.PostId,
		OwnerId:     resp.Owner.Id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to created comment", l.Error(err))
		return
	}
	fmt.Println(com.CreatedAt)
	c.JSON(http.StatusOK, com)
}

// DeleteComment ...
// @Summary DeleteComment
// @Security ApiKeyAuth
// @Tags comment
// @Accept json
// @Produce json
// @Param post_id query string true "PostID"
// @Param id query string true "ID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/delete/comment [delete]
func (h *handlerV1) DeleteComment(c *gin.Context) {
	guid := c.Query("id")
	post_id := c.Query("post_id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	_, err := h.serviceManager.PostService().GetPost(ctx, &pb.GetRequests{
		PostId: post_id,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to query", l.Error(err))
		return
	}

	s, err := h.serviceManager.CommentService().DeleteComment(ctx, &compb.Get{
		Id: guid,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to query", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, s)
}
