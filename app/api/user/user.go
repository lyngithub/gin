package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
	"xx/app/api"
	"xx/app/api/chaptcha"
	"xx/app/domain/user"
	"xx/app/models"
	"xx/forms"
	"xx/utils"
)

type User struct {
	HttpServer utils.IFastHttp
}

func GetUserList(ctx *gin.Context) {
	zap.S().Info("获取用户列表页")

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	// 生成grpc的client并调用接口
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	fmt.Println(pnInt)
	fmt.Println(pSizeInt)
}

func LoginByPass(c *gin.Context) {
	//表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		api.HandleValidatorError(c, err)
		return
	}

	if !chaptcha.Store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	//登录的逻辑
	//生成token
	uid := 1
	nickname := "jimmy"
	role := 2
	nowTime := time.Now().Unix()

	token, err := user.CreateJwtToken(uint(uid), nickname, uint(role), nowTime+60*60*24*30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         uid,
		"nick_name":  nickname,
		"token":      token,
		"expired_at": (nowTime + 60*60*24*30) * 1000,
	})
}

func Register(c *gin.Context) {
	//用户注册
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		api.HandleValidatorError(c, err)
		return
	}
	// 注册逻辑
	uid := 1
	nickname := "jimmy"
	role := 2
	nowTime := time.Now().Unix()
	token, err := user.CreateJwtToken(uint(uid), nickname, uint(role), nowTime+60*60*24*30)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         uid,
		"nick_name":  nickname,
		"token":      token,
		"expired_at": (nowTime + 60*60*24*30) * 1000,
	})
}

func Apple(c *gin.Context) {

	//registerForm := forms.RegisterForm{}
	//	//if err := c.ShouldBind(&registerForm); err != nil {
	//	//	api.HandleValidatorError(c, err)
	//	//	return
	//	//}

	params := "222"
	httpServer := utils.FastHttp("http://www.baidu.com", "GET", params)
	resp := httpServer.Http()
	fmt.Println(resp)
}
