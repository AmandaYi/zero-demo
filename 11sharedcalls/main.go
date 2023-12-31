package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/core/syncx"
	"sync"
	"time"
)

func main() {
	const round = 5
	var wg sync.WaitGroup
	barrier := syncx.NewLockedCalls()

	wg.Add(round)
	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()
			val, err := barrier.Do("once", func() (interface{}, error) {
				time.Sleep(time.Second)
				return stringx.RandId(), nil
			})

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(val)
			}
		}()
	}

	wg.Wait()
}
