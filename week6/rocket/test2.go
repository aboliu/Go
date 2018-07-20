package main

import (
	"fmt"
	"rocket"
	"time"
)

type rocketA struct {
	sec int
}

func (r rocketA) Launch() {
	for i := r.sec; i > 0; i-- {
		fmt.Println("火箭A即將升空，倒數", i, "秒")
		// 休息1秒
		time.Sleep(time.Second)
	}
	fmt.Println("火箭A，咻～升空")
}

type rocketB struct {
	sec int
}

func (r rocketB) Launch() {
	for i := r.sec; i > 0; i-- {
		fmt.Println("火箭B即將升空，倒數", i, "秒")
		// 休息1秒
		time.Sleep(time.Second)
	}
	fmt.Println("火箭B，咻～升空")
}

func main() {
	a := rocketA{
		sec: 10,
	}

	b := rocketB{
		sec: 8,
	}

	go rocket.Launch(a)
	go rocket.Launch(b)

	// 休息20秒
	time.Sleep(time.Second * 20)
}
