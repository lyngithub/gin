package notify

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"xx/app/models"
	"xx/global"
)

func GetChanForm() {
	for i := 0; i < 42; i++ {
		go func() {
			GetData()
		}()
	}
}

func GetData() {

	for form := range global.StatusUpdates {
		str := fmt.Sprintf("%s —— %s —— %s", form.User, form.Status, strconv.Itoa(form.Timestamp))
		global.SugarLogger.Info(str)
		uidString := form.User[strings.Index(form.User, "_")+1 : 6]

		uid, err := strconv.Atoi(uidString)
		if err != nil {
			fmt.Println(err)
			return
		}

		if form.Status == "online" {
			//online = 3
			//time.Sleep(time.Millisecond)
			//
			global.RedisClient.SAdd(context.Background(), models.User_online_status, uid)
		} else {
			global.RedisClient.SRem(context.Background(), models.User_online_status, uid)
		}

		//timestampString := strconv.Itoa(form.Timestamp)
		//timestamp, _ := strconv.Atoi(timestampString[0:10])
		//fmt.Printf("%#v", form)
		//fmt.Println(timestamp)
		//user := model.User{
		//	Id:            uid,
		//	Online:        online,
		//	LastLoginTime: timestamp,
		//}
		//re := user.Update()
		//if re != nil {
		//	initialize.SimpleHttpGet(re.Error())
		//}

		//a := global.RedisClient.SetBit(context.Background(), model.User_online_status, int64(uid), 0)
	}
}
