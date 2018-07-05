package main

import (
	"fmt"
	"sync"
	"time"
)

// 伐木工廠
type woodFactory struct {
	// 產量
	productivity float32
}

// 造紙工廠
type paperFactory struct {
	// 造紙工廠名稱
	name string
	// 每秒可處理紙漿數量
	process float32
	// 產量
	productivity int
}

// 印刷廠
type printingFactory struct {
	// 每秒可印刷出的紙
	productivity int
}

func main() {
	woodFactory := woodFactory{
		productivity: 1,
	}

	paperFac1 := paperFactory{
		name:         "A",
		process:      0.5,
		productivity: 5000,
	}

	paperFac2 := paperFactory{
		name:         "B",
		process:      0.3,
		productivity: 3000,
	}

	printFact := printingFactory{
		productivity: 6000,
	}

	// 當前紙漿數量
	var paperPulpNum float32
	// 當前紙張數量
	var paperNum int
	// 當前印刷紙數量
	var printNum int

	// 建立WaitGroup
	var wg sync.WaitGroup

	// WaitGroup + 1
	wg.Add(1)

	// 伐木囉～
	go cutTree(woodFactory, &paperPulpNum)

	// // 造紙囉～
	go mkPaper(paperFac1, &paperPulpNum, &paperNum)
	go mkPaper(paperFac2, &paperPulpNum, &paperNum)

	// 印刷囉～
	go print(printFact, &wg, &printNum, &paperNum)

	wg.Wait()
	fmt.Println("============ 已完成印刷", printNum, "張紙 ==============")
}

// 伐木場作業
func cutTree(x woodFactory, paperPulpNum *float32) {
	for {
		paperPulpNumTmp := *paperPulpNum + x.productivity

		// 先印在塞資料，顯示的順序看起來比較順眼
		fmt.Println("伐木工廠: 生產了", x.productivity, "頓紙漿, 紙漿剩餘數量:", paperPulpNumTmp, "頓")
		*paperPulpNum = paperPulpNumTmp

		// 休息1秒
		time.Sleep(time.Second)
	}
}

// 造紙廠作業
func mkPaper(x paperFactory, paperPulpNum *float32, paperNum *int) {

	for {
		if *paperPulpNum-x.process >= 0 {
			*paperPulpNum = *paperPulpNum - x.process
			paperNumTmp := *paperNum + x.productivity

			// 先印在塞資料，顯示的順序看起來比較順眼
			fmt.Println("造紙廠", x.name, ": 製成了", x.productivity, "張紙, 目前紙張數量:", paperNumTmp, "張, 紙漿剩餘數量:", *paperPulpNum, "頓")
			*paperNum = paperNumTmp
		}

		// 休息1秒
		time.Sleep(time.Second)
	}
}

// 印刷廠作業
func print(x printingFactory, wg *sync.WaitGroup, printNum *int, paperNum *int) {
	for {
		if *paperNum-x.productivity >= 0 {
			*paperNum = *paperNum - x.productivity
			printNumTmp := *printNum + x.productivity

			// 先印在塞資料，顯示的順序看起來比較順眼
			fmt.Println("印刷廠: 目前印刷了", printNumTmp, "張紙, 剩餘紙張數量:", *paperNum, "張紙")
			*printNum = printNumTmp
		}

		if *printNum >= 60000 {
			//工作完成 回報WaitGroup -1
			wg.Done()
		}

		// 休息1秒
		time.Sleep(time.Second)
	}
}
