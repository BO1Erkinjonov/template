package handler

import (
	models "api-test/api/handlers/models"
	jwt "api-test/api/tokens"
	"api-test/config"
	pb "api-test/genproto/user-service"
	"api-test/mock_data/user_service"
	l "api-test/pkg/logger"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

func GetUser(c *gin.Context) {
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	tok := c.GetHeader("Authorization")
	claims, err := jwt.ExtractClaim(tok, []byte(config.Load().SignInKey))

	response, err := user_service.NewMockUserServiceClient().GetUser(
		ctx, &pb.GetRequest{
			UserId: cast.ToString(claims["sub"]),
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	fmt.Println(response.Post)
	c.JSON(http.StatusOK, response)
}

func GetUsers(c *gin.Context) {
	id := c.Query("id")
	var jspbMarshal protojson.MarshalOptions
	jspbMarshal.UseProtoNames = true

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	response, err := user_service.NewMockUserServiceClient().GetUser(
		ctx, &pb.GetRequest{
			UserId: id,
		})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func GetAllUsers(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	response, err := user_service.NewMockUserServiceClient().GetAllUsers(
		ctx, &pb.GetAllUsersRequest{
			Page:  int64(1),
			Limit: int64(10),
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		l.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateUser(c *gin.Context) {
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
		l.Error(err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	response, err := user_service.NewMockUserServiceClient().Update(ctx, &pb.User{
		FirstName:    body.FirstName,
		LastName:     body.LastName,
		Password:     body.Password,
		Email:        body.Email,
		RefreshToken: "test token",
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		l.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
