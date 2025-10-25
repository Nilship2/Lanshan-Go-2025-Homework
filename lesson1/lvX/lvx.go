package main

import (
	"fmt"
)

func average(sum int, count int) float64 {
	var res float64 = float64(sum)
	res /= float64(count)
	return res
}

func main() {
	var sum = 0
	var count = 0
	for {
		fmt.Println("请输入一个整数(输入0结束):")
		var a int
		fmt.Scanln(&a)
		if a == 0 {
			break
		}
		sum += a
		count += 1
	}
	var ave float64 = average(sum, count)
	if ave >= 60 {
		fmt.Println("平均成绩为", ave, "，成绩合格")
	} else {
		fmt.Println("平均成绩为", ave, "，成绩不合格")
	}
}
