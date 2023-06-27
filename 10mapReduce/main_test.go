package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/timex"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 基本使用,如果第三个不把内容写出去，那么就会死锁，所以可以使用MapReduceVoid，死锁的愿意是信道阻塞
func TestUseMapReduce(t *testing.T) {

	var userIds []string = []string{"1", "2", "3", "4", "5", "6"}

	reduce, err := mr.MapReduce(func(source chan<- string) {
		//MapReduce中首先通过buildSource方法通过执行generate(参数为无缓冲channel)产生数据，并返回无缓冲的channel，mapper会从该channel中读取数据
		for _, uid := range userIds {
			source <- uid
		}
	}, func(item string, writer mr.Writer[uint64], cancel func(error)) {
		fmt.Println(item)
		num, err := strconv.Atoi(item)
		if err != nil {
			cancel(err)
		}
		writer.Write(uint64(num))
	}, func(pipe <-chan uint64, writer mr.Writer[[]uint64], cancel func(error)) {
		var newUids []uint64
		for id := range pipe {
			newUids = append(newUids, id)
		}
		writer.Write(newUids)
	})
	if err != nil {
		return
	}
	fmt.Println(reduce)
}

// 与mapReduce类似，只是没有返回值
func TestUseMapReduceVoid(t *testing.T) {
	var userIds = []string{"1", "2", "3", "4", "5", "6"}
	err := mr.MapReduceVoid(func(source chan<- string) {
		//MapReduce中首先通过buildSource方法通过执行generate(参数为无缓冲channel)产生数据，并返回无缓冲的channel，mapper会从该channel中读取数据
		for _, uid := range userIds {
			source <- uid
		}
	}, func(item string, writer mr.Writer[uint64], cancel func(error)) {
		fmt.Println(item)
		num, err := strconv.Atoi(item)
		if err != nil {
			cancel(err)
		}
		writer.Write(uint64(num))
	}, func(pipe <-chan uint64, cancel func(error)) {
		var newUids []uint64
		for id := range pipe {
			newUids = append(newUids, id)
		}
		fmt.Println(newUids)
	})
	if err != nil {
		return
	}
}
func TestUseFlatMap(t *testing.T) {
	var userIds []string = []string{"1", "2", "3", "4", "5", "6"}
	var lock sync.Mutex
	var newUids []uint64
	mr.ForEach(func(source chan<- string) {
		for _, uid := range userIds {
			source <- uid
		}
	}, func(item string) {
		lock.Lock()
		defer lock.Unlock()

		num, _ := strconv.Atoi(item)
		newUids = append(newUids, uint64(num))
	})
	fmt.Println(newUids)
}
func TestUseFinishVoid(t *testing.T) {
	start := timex.Now()
	mr.FinishVoid(func() {
		time.Sleep(time.Second)
	}, func() {
		time.Sleep(time.Second)
	}, func() {
		time.Sleep(time.Second)
	})
	fmt.Println(timex.Since(start))
}
