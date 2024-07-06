package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"xx/global"
)

func init() {
	fmt.Println("initRedis")
	global.RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Addr, global.ServerConfig.RedisInfo.Port),
		Password: global.ServerConfig.RedisInfo.Password, // no password set
		DB:       global.ServerConfig.RedisInfo.Db,       // use default DB
	})

	//连接redis
	_, err := global.RedisClient.Ping(context.Background()).Result()
	//判断连接是否成功
	if err != nil {
		fmt.Println("redis连接失败")
		println(err)
	} else {
		fmt.Println("redis连接成功")
	}
}
