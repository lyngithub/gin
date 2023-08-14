package router

import "github.com/gin-gonic/gin"
import "xx/app/api/notify"

func InitNotifyRouter(Route *gin.RouterGroup) {
	notifyRouter := Route.Group("notify")
	{
		notifyRouter.POST("", notify.Status)
	}
}
