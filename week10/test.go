/*
	# Week10 練習

	題目：
	撰寫一程式可對目標API 進行連線，需可自行調整併發數量

	記錄：
	單一連線花費時間
	單一最大花費時間
	單一最小花費時間
	壓測總數花費時間
	平均花費時間

	參考關鍵字：cron, slice, (t Time) Sub, goroutine
*/
package main

import (
	"fmt"
	"github.com/robfig/cron"
	"net/http"
	"sync"
	"time"
)

// 併發數量
var num int = 10

func main() {
	fmt.Println("＝＝＝＝＝Starting＝＝＝＝＝")

	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("Run gogogogo...")
		gogogogo()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}

}

func gogogogo() {
	// 連線花費時間列表
	var time_list []time.Duration

	// 建立WaitGroup
	var wg sync.WaitGroup

	// WaitGroup + 1
	wg.Add(num)

	ch := make(chan time.Duration)

	for i := 1; i <= num; i++ {
		go getAPI(&wg, ch)
	}

	for i := 1; i <= num; i++ {
		time_list = append(time_list, <-ch)
	}

	wg.Wait()

	fmt.Println("併發數量： ", num)
	// 印出「單一連線花費時間」
	fmt.Println("＝＝＝＝＝單一連線花費時間＝＝＝＝＝")
	small := time_list[0]
	big := time_list[0]
	var total time.Duration
	for i := range time_list {
		// 計算「單一最大花費時間」
		if time_list[i] < small {
			small = time_list[i]
		}

		// 計算「單一最小花費時間」
		if time_list[i] > big {
			big = time_list[i]
		}

		// 計算「壓測總數花費時間」
		total += time_list[i]

		fmt.Println(time_list[i])
	}

	// 計算「平均花費時間」
	avg := total.Nanoseconds() / int64(num)

	fmt.Println("＝＝＝＝＝其他訊息＝＝＝＝＝")
	fmt.Println("單一最大花費時間: ", big)
	fmt.Println("單一最小花費時間: ", small)
	fmt.Println("壓測總數花費時間: ", total)
	fmt.Println("平均花費時間: ", time.Duration(avg))
}

func getAPI(wg *sync.WaitGroup, ch chan time.Duration) {
	startTime := time.Now()

	client := &http.Client{}
	//向服务端发送get请求
	request, _ := http.NewRequest("GET", "http://bb.dev.d2/api/payment_gateway", nil)
	//接收服务端返回给客户端的信息
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		// 	str, _ := ioutil.ReadAll(response.Body)
		// 	bodystr := string(str)
		// 	return bodystr
	}

	// return ""

	endTime := time.Now()
	elapsed := endTime.Sub(startTime)
	// fmt.Println(reflect.TypeOf(elapsed)) // ch <- elapsed
	ch <- elapsed
	wg.Done()
}
