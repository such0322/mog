package main

import (
	"fmt"
	"github.com/liangdas/armyant/task"
	test_task "mog/robot/test"
	"os"
	"os/signal"
)

func main() {

	task := task.LoopTask{
		C: 1, //并发数
	}
	//manager := table_test.NewManager(task) //房间模型的demo
	manager := test_task.NewManager(task) //gate demo
	fmt.Println("开始压测请等待")
	task.Run(manager)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	task.Stop()
	os.Exit(1)
}
