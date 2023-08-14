package token

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"xx/app/models"
	"xx/global"
	"xx/utils"
)

type Resonse struct {
	Application string `json:"application"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetToken(ctx *gin.Context) {

	url := global.ServerConfig.IMInfo.TokenUrl

	params := map[string]interface{}{
		"grant_type":    global.ServerConfig.IMInfo.GrantType,
		"client_id":     global.ServerConfig.IMInfo.ClientId,
		"client_secret": global.ServerConfig.IMInfo.ClientSecret,
	}

	res, err := utils.HttpPost(url, map[string]string{}, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var tokenRes Resonse
	err1 := json.Unmarshal([]byte(res), &tokenRes)
	if err1 != nil {
		fmt.Println(res)
		fmt.Println("解析 JSON 失败:", err)
		return
	}
	global.RedisClient.Set(context.Background(), models.Huanxin_token, tokenRes.AccessToken, 30*24*60*60)

	fmt.Println(tokenRes.AccessToken)
}
