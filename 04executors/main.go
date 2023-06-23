package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/executors"
	"time"
)

type DailyTask struct {
	insertExecutor *executors.BulkExecutor
}

func (dts *DailyTask) initBulkExecutor() {

	//初始化
	//这是 NewBulkExecutor函数选项模式
	dts.insertExecutor = executors.NewBulkExecutor(
		func(tasks []any) {
			fmt.Println(tasks)
			for _, v := range tasks {
				dts.deal(v)
			}
		},
		//这俩条件只要满足一个就开始拿出所有的任务进行执行了
		executors.WithBulkInterval(time.Second*3), // 3s会自动刷一次container中task去执行
		executors.WithBulkTasks(4),                // container最大task数。一般设为2的幂次
	)
}
func (dts *DailyTask) deal(num interface{}) {
	fmt.Println("拿到值num", num.(int))
}

// 外部追加任务 者
func (dts *DailyTask) addTask() {
	var i int = 0
	for {
		time.Sleep(time.Second)
		if i > 1000 {
			break
		}
		err := dts.insertExecutor.Add(i)
		if err != nil {
			fmt.Println("无法增加任务了", err)
		}
		i++
	}
	dts.insertExecutor.Flush()
	dts.insertExecutor.Wait()
}

func main() {
	var dts DailyTask = DailyTask{}
	dts.initBulkExecutor() // 初始化执行器
	dts.addTask()
}
