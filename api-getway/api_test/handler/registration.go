package handler

import (
	"api-test/config"
	pbu "api-test/genproto/user-service"
	"api-test/mock_data/user_service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	exists_email, err := user_service.NewMockUserServiceClient().CheckUniquess(ctx, &pbu.CheckUniqReq{
		Field: "email",
		Value: "email",
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

	exists, err := user_service.NewMockUserServiceClient().CheckUniquess(ctx, &pbu.CheckUniqReq{
		Field: "username",
		Value: "username",
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

	c.JSON(http.StatusOK, true)
}

func Verification(c *gin.Context) {
	var regis pbu.User
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()
	_, err := user_service.NewMockUserServiceClient().Create(ctx, &pbu.User{
		FirstName:    regis.FirstName,
		LastName:     regis.LastName,
		Username:     regis.Username,
		Role:         "user",
		Password:     regis.Password,
		Email:        regis.Email,
		Id:           regis.Id,
		RefreshToken: "test token",
	})
	if err != nil {
		log.Fatalln(err)
	}

	c.JSON(http.StatusOK, "test token")

}

func LogIn(c *gin.Context) {
	password := "Mock password"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(config.Load().CtxTimeout))
	defer cancel()

	user, err := user_service.NewMockUserServiceClient().Exists(ctx, &pbu.Req{
		Email: "Mock email",
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, "email error")
		return
	}

	if password != user.Password {
		c.JSON(http.StatusBadRequest, "password error")
		return
	}

	_, err = user_service.NewMockUserServiceClient().Update(ctx, &pbu.User{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Role:         user.Role,
		Password:     user.Password,
		Email:        user.Email,
		Id:           user.Id,
		RefreshToken: "test token",
	})
	c.JSON(http.StatusOK, "test token")
}
