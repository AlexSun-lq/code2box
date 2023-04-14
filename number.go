package code2box

import "math"

// 速度慢 不用 ,3倍
func getDigitsNumv2(num int) uint {
	return uint(math.Log10(float64(num))) + 1
}
func GetDigitsNum(num int) uint {
	var b uint = 0
	for num > 0 {
		num = num / 10
		b++
	}
	return b
}
