package main

import (
	"fmt"
)

func main() {
	var score [5]int = [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println("====== 練習1 ======")
	fmt.Println("分數:", score)

	sum := getSun(score)
	fmt.Println("總分:", sum)

	divisor := len(score)
	average := getAverage(divisor, sum)
	fmt.Println("平均值:", average)

	fmt.Println("====== 練習2 ======")
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := getMin(x)
	fmt.Println("最小值:", min)
}

func getSun(score [5]int) int {
	var total int
	for i := 0; i < len(score); i++ {
		total += score[i]
	}

	return total
}

func getAverage(divisor int, dividend int) int {
	average := dividend / divisor

	return average
}

func getMin(arr []int) int {
	var min int
	//預設 - 最小值
	min = arr[0]
	for i := 1; i < len(arr); i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}

	return min
}
