package main

import (
	"./worker"
	"fmt"
	"log"
	"os"
	"time"
)

const timeout = 100 * time.Second

func main()  {
	fmt.Println("开始运行...")

	// 初始化作业运行器
	r := worker.New(timeout)

	// 调度三个后台处理任务
	for i:=0;i<10;i++{
		r.Add(createTask())

	}

	// 启动作业运行器
	if err := r.Start(); err != nil {
		switch err {
		case worker.ErrTimeout:
			log.Println("作业程序因运行超时而终止")
			os.Exit(1)
		case worker.ErrInterrupt:
			log.Println("作业程序因系统发生中断事件而终止")
			os.Exit(2)
		}
	}
}

// 编写一个模拟后台处理任务
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}