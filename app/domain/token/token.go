package token

import (
	"context"
	"fmt"
	"xx/global"
	"xx/models"
	"xx/utils"
)

func GetOnline() {

	allMember := global.RedisClient.SMembers(context.Background(), models.User_online_status)
	//for member := range allMember.Val() {
	//
	//}

	url := "https://a1.easemob.com/1150230316212388/poyoo/users/1150230316212388/presence"
	params := map[string]interface{}{
		"host":     "a1.easemob.com",
		"org_name": "1150230316212388",
		"app_name": "poyoo",
		"username": allMember.Val(),
	}
	res, err := utils.HttpPost(url, map[string]string{}, params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(res)
}
