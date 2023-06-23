package main

import (
	"github.com/zeromicro/go-zero/core/discov"
	"log"
	"time"
)

// 发布etcd
func main() {
	var person string = "zzy"
	client := discov.NewPublisher([]string{"http://127.0.0.1:2379"}, "person", person)
	if err := client.KeepAlive(); err != nil {
		log.Fatal(err)
	}

	defer client.Stop()
	for {
		time.Sleep(time.Second)
	}
}
