package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/discov"
	"time"
)

func main() {
	sub, err := discov.NewSubscriber([]string{"http://127.0.0.1:2379"}, "person", discov.Exclusive())
	fmt.Println(err)
	sub.AddListener(func() {
		fmt.Println("有值")
	})
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			{
				fmt.Println("读取", sub.Values())
			}

		}
	}
}
