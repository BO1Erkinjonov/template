package v1

import (
	"api-test/api/handlers/models"
	token "api-test/api/tokens"
	compb "api-test/genproto/comment-service"
	postpb "api-test/genproto/post-service"
	pb "api-test/genproto/user-service"
	l "api-test/pkg/logger"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"time"
)

// AdminUpdatePost bu funksiyada admin xoxlagan userni postdagi malumotlarni o'zgartiroladi
// @Summary AdminUpdatePost
// @Security ApiKeyAuth
// @Description Api for creating a new post
// @Tags admin
// @Accept json
// @Produce json
// @Param Post body models.AdminPost true "update post"
// @Param post_id query string true "id"
// @Success 200 {object} models.AdminPost
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/admin/update/post [put]
func (h *handlerV1) AdminUpdatePost(c *gin.Context) {
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
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	post, err := h.serviceManager.PostService().GetPost(ctx, &postpb.GetRequests{PostId: id})
	fmt.Println(post)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to query", l.Error(err))
		return
	}
	response, err := h.serviceManager.PostService().UpdatePost(ctx, &postpb.Post{
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
		h.log.Error("failed to updated post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// AdminUpdateUser bu funksiyada admin xoxlagan userni malumotlarni o'zgartiroladi
// @Summary AdminUpdateUser
// @Security ApiKeyAuth
// @Description Api for updating user
// @Tags admin
// @Accept json
// @Produce json
// @Param User body models.User true "createUserModel"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/admin/update/user [put]
func (h *handlerV1) AdminUpdateUser(c *gin.Context) {
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
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	_, refresh, err := h.jwthandler.GenerateAuthJWT()

	response, err := h.serviceManager.UserService().Update(ctx, &pb.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Username:     body.UserName,
		Password:     body.Password,
		Email:        body.Email,
		RefreshToken: refresh,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to update user", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// AdminDeleteUser bu funksiyada admin xoxlagan userni delete qiloladi
// @Summary AdminDeleteUser
// @Security ApiKeyAuth
// @Description Api for deleting user by id
// @Tags admin
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/admin/delete/user [delete]
func (h *handlerV1) AdminDeleteUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	user2, err := h.serviceManager.UserService().GetUser(ctx, &pb.GetRequest{UserId: guid})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		h.log.Error("failed to parse query params json", l.Error(err))
		return
	}
	if user2.Role == "superAdmin" || user2.Role == "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "permission denied1",
		})
		h.log.Error("permission denied")
		return
	}

	response, err := h.serviceManager.UserService().Delete(
		ctx, &pb.GetRequest{
			UserId: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to deleted user", l.Error(err))
		return
	}

	_, err = h.serviceManager.PostService().DeletePostByOwnerId(ctx, &postpb.GetByOwnerIdRequest{
		OwnerId: user2.Id,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to deleted post", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// SuperAdminDeleteAdmin bu funksiyada superAdmin xoxlagan user va admin delete qiloladi
// @Summary SuperAdminDeleteAdmin
// @Security ApiKeyAuth
// @Description Api for deleting user by id
// @Tags superAdmin
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/superAdmin/delete/admin [delete]
func (h *handlerV1) SuperAdminDeleteAdmin(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().Delete(
		ctx, &pb.GetRequest{
			UserId: guid,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to delete user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// AdminDeleteComment  bu funksiyada admin xoxlagan userni postdagi commentini delete qiloladi
// @Summary AdminDeleteComment
// @Security ApiKeyAuth
// @Tags admin
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/admin/delete/comment [delete]
func (h *handlerV1) AdminDeleteComment(c *gin.Context) {
	guid := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

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

// AdminDeletePost bu funksiyada admin xoxlagan port ochiroladi
// @Summary AdminDeletePost
// @Security ApiKeyAuth
// @Description Api for deleting post
// @Tags admin
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.Post
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/admin/delete/post [delete]
func (h *handlerV1) AdminDeletePost(c *gin.Context) {
	id := c.Query("id")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.PostService().DeletePost(ctx, &postpb.GetRequests{PostId: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to deleted post", l.Error(err))
		return
	}
	c.JSON(http.StatusCreated, response)
}

// Create bu funksiyada admin xoxlagan roleda foydalanuvchi qosholadi
// @Summary Create
// @Security ApiKeyAuth
// @Description Api for creating a new user
// @Tags admin
// @Accept json
// @Produce json
// @Param User body models.AdminUser true "createUserModel"
// @Success 200 {object} models.AdminUser
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/admin/create/user [post]
func (h *handlerV1) Create(c *gin.Context) {

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
		h.log.Error("failed to bind json", l.Error(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	if body.Role == "superAdmin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "permission denied: an ordinary admin will not be able to add a superAdmin",
		})
		return
	}
	exists_email, err := h.serviceManager.UserService().CheckUniquess(ctx, &pb.CheckUniqReq{
		Field: "email",
		Value: body.Email,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to check email uniques")
		return
	}

	if exists_email.Code == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This email already in use, please use another email address",
		})
		h.log.Error("failed to check email uniques", l.Error(err))
		return
	}

	exists, err := h.serviceManager.UserService().CheckUniquess(ctx, &pb.CheckUniqReq{
		Field: "username",
		Value: body.UserName,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to check username uniques")
		return
	}
	if exists.Code == 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "This username already in use, please use another username address",
		})
		h.log.Error("failed to check username uniques", l.Error(err))
		return
	}

	Id := uuid.NewString()
	h.jwthandler = token.JWTHandler{
		Sub:     Id,
		Role:    "user",
		SignKey: h.cfg.SignInKey,
		Timout:  h.cfg.AccessTokenTimout,
	}

	_, refresh, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		log.Fatalln(err)
	}

	response, err := h.serviceManager.UserService().Create(ctx, &pb.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Username:     body.UserName,
		Role:         body.Role,
		Password:     body.Password,
		Email:        body.Email,
		Id:           Id,
		RefreshToken: refresh,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to create user", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, response)
}
