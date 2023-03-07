package sys_init

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"spider-golang-web/global"
)

func InitRedis() {
	global.RedisDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", global.Host),
		Password: "123456", // no password set
		DB:       0,        // use default DB
	})
	res, err := global.RedisDB.Ping(context.Background()).Result()
	if err != nil {
		zap.S().Errorf("error to start redis,detail:%v", err)
		return
	}
	fmt.Println(res)
}
