package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

//这个包自带了连接池，不管它
var redis_client *redis.Client

//连接redis
func CreateRedisConnection(ip, port string, db int) {

	redis_client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ip, port),
		Password: "",
		DB:       db,
	})
}

//获得redis客户端
func GetRedisClient() *redis.Client {
	if redis_client == nil {
		//log.Fataln函数会执行os.Exit让程序退出
		log.Fatalln("大哥，你没有链接redis啊")
	}
	return redis_client
}
