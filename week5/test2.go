package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	total := 0
	for i := 1; i <= 20; i++ {
		total = total + factorial(i)
	}
	costTime := time.Since(startTime)
	fmt.Println("1!+2!+....+20! =", total, ", 計算花費時間:", costTime)

}

func factorial(num int) int {
	sum := 1
	for i := 1; i <= num; i++ {
		sum = sum * i
	}

	return sum
}
