package api

import (
	_ "api-test/api/docs"
	casbinC "api-test/api/handlers/middleware/casbin"
	v1 "api-test/api/handlers/v1"
	token "api-test/api/tokens"
	"api-test/config"
	"api-test/pkg/logger"
	"api-test/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	Enforcer       casbin.Enforcer
	CasbinEnforcer *casbin.Enforcer
	ServiceManager service.IServiceManager
}

// @title Bobur Erkinjonov
// @version 1.7
// @host localhost:1212

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()

	jwtHandler := token.JWTHandler{
		SignKey: option.Conf.SignInKey,
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Jwthandler:     jwtHandler,
	})

	router.Use(casbinC.NewAuthorizer())
	api := router.Group("/v1")

	// Super Admin
	api.DELETE("/superAdmin/delete/admin", handlerV1.SuperAdminDeleteAdmin)

	// Admin
	api.PUT("/admin/update/post", handlerV1.AdminUpdatePost)
	api.PUT("/admin/update/user", handlerV1.AdminUpdateUser)
	api.DELETE("/admin/delete/user", handlerV1.AdminDeleteUser)
	api.DELETE("/admin/delete/comment", handlerV1.AdminDeleteComment)
	api.DELETE("/admin/delete/post", handlerV1.AdminDeletePost)
	api.POST("/admin/create/user", handlerV1.Create)

	// User
	api.GET("/get/user", handlerV1.GetUser)
	api.GET("/get/users", handlerV1.GetUsers)
	api.GET("/users/all", handlerV1.GetAllUsers)
	api.PUT("/update/user", handlerV1.UpdateUser)

	// Post
	api.POST("/create/post", handlerV1.CreatePost)
	api.GET("/get/post", handlerV1.GetPost)
	api.GET("/posts/all", handlerV1.GetAllPosts)
	api.PUT("/update/post", handlerV1.UpdatePost)
	api.DELETE("/delete/post", handlerV1.DeletePost)
	api.PUT("/like/post", handlerV1.Like)

	// Comment
	api.POST("/create/comment", handlerV1.CreateComment)
	api.GET("/get/comment", handlerV1.GetComment)
	api.GET("/comments/all", handlerV1.GetAllComment)
	api.PUT("/update/comment", handlerV1.UpdateComment)
	api.DELETE("/delete/comment", handlerV1.DeleteComment)

	// Authorization
	api.POST("/Verification", handlerV1.Verification)
	api.POST("/register", handlerV1.Register)
	api.POST("/login", handlerV1.LogIn)

	url := ginSwagger.URL("swagger/doc.json")

	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
