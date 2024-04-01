package v1

import (
	models "api-test/api/handlers/models"
	jwt "api-test/api/tokens"
	"api-test/config"
	pb "api-test/genproto/user-service"
	l "api-test/pkg/logger"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"strconv"
	"time"
)

// GetUser bu funksiyada user faqat o'ziga tegishli bolgan malumotlarni koroladi
// @Summary GetUser
// @Security ApiKeyAuth
// @Description Api for getting user by id
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/get/user [get]
func (h *handlerV1) GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	tok := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(tok, []byte(config.Load().SignInKey))

	response, err := h.serviceManager.UserService().GetUser(
		ctx, &pb.GetRequest{
			UserId: cast.ToString(claims["sub"]),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetUsers bu funksiyada id orqali xoxlagan userni malumotlarni olsa boladi
// @Summary GetUsers
// @Security ApiKeyAuth
// @Description Api for getting user by id
// @Tags user
// @Accept json
// @Produce json
// @Param id query string true "ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/get/users [get]
func (h *handlerV1) GetUsers(c *gin.Context) {
	id := c.Query("id")
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetUser(
		ctx, &pb.GetRequest{
			UserId: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.log.Error("failed to get user", l.Error(err))
		return
	}
	c.JSON(http.StatusOK, response)
}

// GetAllUsers
// @Summary GetAllUsers
// @Security ApiKeyAuth
// @Description Api for getting users
// @Tags user
// @Accept json
// @Produce json
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/users/all [get]
func (h *handlerV1) GetAllUsers(c *gin.Context) {
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
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	response, err := h.serviceManager.UserService().GetAllUsers(
		ctx, &pb.GetAllUsersRequest{
			Page:  int64(reqPage),
			Limit: int64(reqLimit),
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		h.log.Error("failed to list users", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser bu funksiyada user faqat o'zini update qiloladi
// @Summary UpdateUser
// @Security ApiKeyAuth
// @Description Api for updating user
// @Tags user
// @Accept json
// @Produce json
// @Param User body models.User true "createUserModel"
// @Success 200 {object} models.User
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/update/user [put]
func (h *handlerV1) UpdateUser(c *gin.Context) {
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

	tok := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(tok, []byte(config.Load().SignInKey))

	user, err := h.serviceManager.UserService().GetUser(ctx, &pb.GetRequest{UserId: cast.ToString(claims["sub"])})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		h.log.Error("failed to parse query params json", l.Error(err))
		return
	}

	_, refresh, err := h.jwthandler.GenerateAuthJWT()

	response, err := h.serviceManager.UserService().Update(ctx, &pb.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Username:     body.UserName,
		Role:         user.Role,
		Password:     body.Password,
		Email:        user.Email,
		Id:           user.Id,
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
