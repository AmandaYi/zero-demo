package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	redisClient, err := redis.NewRedis(redis.RedisConf{
		Host:        "127.0.0.1:6379",
		Type:        "node",
		Pass:        "123456",
		Tls:         false,
		NonBlock:    false,
		PingTimeout: 0,
	})
	if err != nil {
		fmt.Println(err)
	}
	redisLook := redis.NewRedisLock(redisClient, "user_zzy")

	redisLook.SetExpire(3600)
	//Acquire尝试获取一个值，具体逻辑，可以参考源码，使用redis执行一端lua脚本，表示原子性操作
	if ok, err := redisLook.Acquire(); !ok || err != nil {
		fmt.Println("当前有其他用户正在进行操作，请稍后重试")
	}
	defer func() {
		//不能释放别人的锁，不能释放别人的锁，不能释放别人的锁
		//get(key) == value「key」
		redisLook.Release() // 释放redis锁
	}()

}
