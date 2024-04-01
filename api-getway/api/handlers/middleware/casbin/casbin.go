package casbin

import (
	jwt "api-test/api/tokens"
	"api-test/config"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CasbinHandler struct {
	config     config.Config
	enforce    casbin.Enforcer
	jwthandler jwt.JWTHandler
}

func NewAuthorizer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token1 := ctx.GetHeader("Authorization")
		if token1 == "" {
			sub := "unauthorized"
			obj := ctx.Request.URL.Path
			etc := ctx.Request.Method
			fmt.Println(sub, obj, etc)
			e, _ := casbin.NewEnforcer("auth.conf", "auth.csv")
			t, _ := e.Enforce(sub, obj, etc)
			if t {
				ctx.Next()
				return
			}
		}

		claims, err := jwt.ExtractClaim(token1, []byte(config.Load().SignInKey))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		sub := claims["role"]
		obj := ctx.Request.URL.Path
		etc := ctx.Request.Method
		fmt.Println(sub, obj, etc)
		e, _ := casbin.NewEnforcer("auth.conf", "auth.csv")
		t, _ := e.Enforce(sub, obj, etc)
		fmt.Println(t)
		if t {
			ctx.Next()
			return
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "permission denied",
		})
	}
}
