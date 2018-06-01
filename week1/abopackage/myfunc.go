package abopackage

import (
	"strconv"
)

/**
 * int 轉 String
 */
func IntToStr(num int) string {
	res := strconv.Itoa(num)

	return res
}
