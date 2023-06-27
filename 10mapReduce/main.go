package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mr"
	"strconv"
)

func main() {

}

func init() {

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
