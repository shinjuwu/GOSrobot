package main

import (
	"GOSrobot/robot"
	"fmt"
	"os"
	"os/signal"

	"github.com/liangdas/armyant/task"
)

func main() {
	task := task.LoopTask{
		C: 5,
	}
	manager := robot.NewManager(task)
	fmt.Println("機器人出動中....")
	task.Run(manager)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	os.Exit(1)
}
