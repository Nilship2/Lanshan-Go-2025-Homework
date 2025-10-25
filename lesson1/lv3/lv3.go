package main

import (
	"fmt"
)

func f(n int) int {
	var res int = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	n := 6
	fmt.Scanln(&n)
	fmt.Println(f(n))
}
