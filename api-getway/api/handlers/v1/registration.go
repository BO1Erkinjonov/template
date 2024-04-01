package v1

import (
	"api-test/api/handlers/models"
	token "api-test/api/tokens"
	pbu "api-test/genproto/user-service"
	l "api-test/pkg/logger"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"strings"
	"time"
)

// Register ...
// @Summary Register
// @Description Register - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param Register body models.RegisterModel true "createRegisterModel"
// @Success 200 {object} models.RegisterModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/register/ [post]
func (h *handlerV1) Register(c *gin.Context) {
	var (
		body models.RegisterResponse
	)

	err := c.ShouldBindJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.log.Error("Failed to bind json", l.Error(err))
		return
	}
	body.Email = strings.TrimSpace(body.Email)
	body.Email = strings.ToLower(body.Email)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	exists_email, err := h.serviceManager.UserService().CheckUniquess(ctx, &pbu.CheckUniqReq{
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

	exists, err := h.serviceManager.UserService().CheckUniquess(ctx, &pbu.CheckUniqReq{
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

	body.Id = uuid.New().String()

	rdb := redis.NewClient(&redis.Options{
		Addr: "redisdb:6379",
	})
	byteDate, err := json.Marshal(&body)
	if err != nil {
		log.Fatalln(err)
	}

	err = rdb.Set(context.Background(), "email_"+body.Email, byteDate, 0).Err()
	if err != nil {
		log.Fatalln(err)
	}

	err = rdb.SetEx(context.Background(), body.Email, exists.Code, time.Minute*1).Err()
	if err != nil {
		log.Fatalln(err)
	}
	code := strconv.Itoa(int(exists.Code))

	auth := smtp.PlainAuth("", "boburerkinzonov@gmail.com", "llqmgbilccvhltfd", "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, "boburerkinzonov@gmail.com", []string{body.Email}, []byte(code))
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, true)
}

// Verification ...
// @Summary Verification
// @Description Verification - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param code query int true "Code"
// @Param email query string true "Email"
// @Success 200 {object} models.RegisterModel
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/Verification/ [post]
func (h *handlerV1) Verification(c *gin.Context) {
	codeRegis := c.Query("code")
	emailRegis := c.Query("email")

	rdb := redis.NewClient(&redis.Options{
		Addr: "redisdb:6379",
	})

	respCode, err := rdb.Get(context.Background(), emailRegis).Result()
	if err != nil {
		log.Fatalln(err)
	}
	var code int
	if err := json.Unmarshal([]byte(respCode), &code); err != nil {
		log.Fatalln(err)
	}

	code1, err := strconv.Atoi(codeRegis)
	if err != nil {
		log.Fatalln(err)
	}

	if code != code1 {
		c.JSON(http.StatusBadRequest, false)
	} else {

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
		defer cancel()
		err = rdb.Del(context.Background(), emailRegis).Err()
		if err != nil {
			log.Fatalln(err)
		}
		var regis models.RegisterResponse
		respUser, err := rdb.Get(context.Background(), "email_"+emailRegis).Result()
		if err != nil {
			log.Fatalln(err)
		}
		err = rdb.Del(context.Background(), "email_"+emailRegis).Err()
		if err != nil {
			log.Fatalln(err)
		}
		if err := json.Unmarshal([]byte(respUser), &regis); err != nil {
			log.Fatalln(err)
		}

		h.jwthandler = token.JWTHandler{
			Sub:     regis.Id,
			Role:    "user",
			SignKey: h.cfg.SignInKey,
			Timout:  h.cfg.AccessTokenTimout,
		}

		access, refresh, err := h.jwthandler.GenerateAuthJWT()
		if err != nil {
			log.Fatalln(err)
		}

		_, err = h.serviceManager.UserService().Create(ctx, &pbu.User{
			FirstName:    regis.FirstName,
			LastName:     regis.LastName,
			Username:     regis.UserName,
			Role:         "user",
			Password:     regis.Password,
			Email:        regis.Email,
			Id:           regis.Id,
			RefreshToken: refresh,
		})
		if err != nil {
			log.Fatalln(err)
		}

		c.JSON(http.StatusOK, access)
	}
}

// LogIn ...
// @Summary LogIn
// @Description LogIn - Api for registering users
// @Tags Register
// @Accept json
// @Produce json
// @Param password query string true "Password"
// @Param email query string true "Email"
// @Success 200 {object} models.RegisterResponse
// @Failure 400 {object} models.StandardErrorModel
// @Failure 500 {object} models.StandardErrorModel
// @Router /v1/login/ [post]
func (h *handlerV1) LogIn(c *gin.Context) {
	password := c.Query("password")
	email := c.Query("email")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	user, err := h.serviceManager.UserService().Exists(ctx, &pbu.Req{
		Email: email,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, "email error")
		return
	}

	if password != user.Password {
		c.JSON(http.StatusBadRequest, "password error")
		return
	}

	h.jwthandler = token.JWTHandler{
		Sub:     user.Id,
		Role:    user.Role,
		SignKey: h.cfg.SignInKey,
		Timout:  h.cfg.AccessTokenTimout,
	}
	access, refresh, err := h.jwthandler.GenerateAuthJWT()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = h.serviceManager.UserService().Update(ctx, &pbu.User{
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Username:     user.Username,
		Role:         user.Role,
		Password:     user.Password,
		Email:        user.Email,
		Id:           user.Id,
		RefreshToken: refresh,
	})
	c.JSON(http.StatusOK, access)
}
