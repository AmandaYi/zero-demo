package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/timex"
	"time"
)

func main() {
	ticker := timex.NewTicker(time.Second)
	//10 代表一圈的刻度
	timingWheel, err := collection.NewTimingWheelWithTicker(time.Second, 10, execTest, ticker)
	if err != nil {
		return
	}
	//代表指针指向
	timingWheel.SetTimer("name", "xiaoming", time.Second) // 延迟1秒执行name
	timingWheel.SetTimer("age", "18", time.Second*2)      // 延迟2秒执行age
	timingWheel.SetTimer("source", "100", time.Second*3)  // 延迟3秒执行source
	timingWheel.SetTimer("desc", "well", time.Second*20)  // 一圈时钟后执行desc

	// 这个Drain代表一瞬间清除所有的执行器
	//timingWheel.Drain(func(key, value any) {
	//	fmt.Println("drain", key, "=", value)
	//})
	//
	defer func() {
		ticker.Stop()
		timingWheel.Stop()
	}()

	for {
		time.Sleep(time.Second)
	}
}
func execTest(key, value interface{}) {
	k := key.(string)
	v := value.(string)
	fmt.Println(k, "=", v)
}
