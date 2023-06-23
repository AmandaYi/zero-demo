package main

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/bloom"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func main() {
	store := redis.New("127.0.0.1:6379", func(r *redis.Redis) {
		r.Pass = "123456"
		r.Type = redis.NodeType // 单点类型
		//r.Type = redis.ClusterType // 集群类型
		r.Addr = "127.0.0.1:6379"
	})
	filter := bloom.New(store, "testbloom", 1024)
	filter.Add([]byte("123456"))
	filter.Add([]byte("zzy"))

	//查找是否存在当前值，如果存在才能执行业务
	if ok, err := filter.Exists([]byte("123456")); ok && err != nil {
		fmt.Println("正常执行业务 do something")
	} else {
		fmt.Println("不存在的数据，请联系管理员")
	}
}
