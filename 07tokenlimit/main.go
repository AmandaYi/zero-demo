package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
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
	fmt.Println(redisClient.Ping()) // true

	//新建一个令牌桶
	const (
		burst   = 10 // 令牌桶最大值
		rate    = 1  // 每秒生成几个令牌
		seconds = 5
	)
	//从整体上令牌桶生产token逻辑如下：
	//
	//用户配置的平均发送速率为r，则每隔1/r秒一个令牌被加入到桶中；
	//假设桶中最多可以存放b个令牌。如果令牌到达时令牌桶已经满了，那么这个令牌会被丢弃；
	//当流量以速率v进入，从桶中以速率v取令牌，拿到令牌的流量通过，拿不到令牌流量不通过，执行熔断逻辑；
	//go-zero 在两类限流器下都采取 lua script 的方式，依赖redis可以做到分布式限流，lua script同时可以做到对 token 生产读取操作的原子性。
	limiter := limit.NewTokenLimiter(rate, burst, redisClient, "test-token-limit-key")

	for {
		if limiter.Allow() {
			fmt.Println("允许通过")
		} else {
			fmt.Println("不允许通过")
			// 这里可以做一些操作，比如把这些请求上下文给缓存起来，做缓存队列，延迟执行，等
		}
		time.Sleep(time.Millisecond * 100)
	}

	return
	//下面都是测试逻辑，闹着玩的
	timer := time.NewTimer(time.Second * seconds)

	quit := make(chan struct{})

	defer timer.Stop()

	go func() {
		<-timer.C
		close(quit)
	}()

	//允许的，禁止的
	var allowed, denied int32
	var wait sync.WaitGroup

	for i := 0; i < runtime.NumCPU(); i++ {
		wait.Add(1)

		go func() {
			for {
				select {
				case <-quit:
					{
						wait.Done()
						return
					}
				default:
					if limiter.Allow() {
						atomic.AddInt32(&allowed, 1)
					} else {
						atomic.AddInt32(&denied, 1)
					}
				}
			}
		}()
	}

	wait.Wait()

	fmt.Printf("allowed: %d, denied: %d, qps: %d\n", allowed, denied, (allowed+denied)/seconds)

}
