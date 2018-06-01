package main

import (
	"./abopackage"
	"fmt"
	"strconv"
)

var a int = 1
var b int32 = 2
var c int64 = 3
var d string = "999"
var e float32 = 88.8
var f float64 = 99.9
var x string = "I Love Golang_"

func main() {
	res := a + int(b)
	fmt.Println("a + b = ", res)

	res2 := a + int(b) + int(c)
	fmt.Println("a + b + c = ", res2)

	res3 := f / float64(e)
	fmt.Println("f / e = ", res3)

	d2 := strToInt(d)
	res4 := a + d2
	fmt.Println(res4)

	a1 := abopackage.IntToStr(a)
	res5 := x + a1
	fmt.Println(res5)

}

/**
 * String 轉 int
 */
func strToInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}

	return res
}
