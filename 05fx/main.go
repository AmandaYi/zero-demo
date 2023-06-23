package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/fx"
	"time"
)

// fx跟Promise基本上一样
func main() {
	fx.From(func(source chan<- any) {
		var i int = 0
		for {
			source <- i
			i++
			time.Sleep(time.Second)
			if i > 10 {
				break
			}
		}
	}).Filter(func(item any) bool {
		//过滤偶数
		return item.(int)%2 == 0
	}).Map(func(item any) any {
		//给每个数乘以2
		return item.(int) * 2
	}).Distinct(func(item any) any {
		//去重
		return item
	}).Reverse().Reduce(func(pipe <-chan any) (any, error) {
		//reduce是一个处理全部数据的函数，必须等待全部的流数据输入完毕后，才会执行
		var result int
		for item := range pipe {
			i := item.(int)
			result += i
		}
		fmt.Println(result)
		return result, nil
	})

	//并发调用RPC
	fx.Parallel(func() {
		fmt.Println("调用RPC")
	}, func() {
		fmt.Println("调用RPC")
	}, func() {
		fmt.Println("调用RPC")
	})

}
