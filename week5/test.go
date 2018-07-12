package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	// goroutine1
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	// goroutine2
	go func() {
		for {
			a, ok := <-ch

			if !ok {
				fmt.Println("close")
				close(ch)
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	// 此處關閉了channl，但後面goroutine1卻仍繼續塞資料給已經關閉的channl
	// close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}
