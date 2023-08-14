package router

import (
	"github.com/gin-gonic/gin"
	"xx/app/api/user"
	"xx/app/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("", middlewares.JWTAuth(), user.GetUserList)
		UserRouter.GET("list", middlewares.JWTAuth(), user.GetUserList)
		UserRouter.POST("pwd_login", user.LoginByPass)
		UserRouter.POST("register", user.Register)
		UserRouter.GET("http", user.Apple)
	}
}
