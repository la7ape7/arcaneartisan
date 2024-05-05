package middlewares

import (
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"time"
	"watIwant/config"
	"watIwant/dao"
	"watIwant/models"
)

func JwtMiddlewareHandler() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:      "WatIWant",
		Key:        []byte(config.GetConfiguration().Environment.ApiSecretKey),
		Timeout:    time.Hour * 1,
		MaxRefresh: time.Hour / 2,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"id": v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals models.UserLogin
			if err := c.BindJSON(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			loggedIn := Authenticate(loginVals)
			if loggedIn {
				return dao.NewUserDAO().GetPublicUser(loginVals.Username), nil
			}
			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(context *gin.Context, code int, mesage string) {
			if mesage == "auth header is empty" {
				context.JSON(code, gin.H{"code": code, "message": "You need to login", "authEndpoint": context.Request.Host + "/auth"})
			} else {
				context.JSON(code, gin.H{"code": code, "message": mesage})
			}
		},
		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

func Authenticate(loginUser models.UserLogin) bool {
	loggedIn := dao.NewAuthDAO().Login(loginUser.Username, loginUser.Password)
	if loggedIn {
		return true
	}
	return false
}
