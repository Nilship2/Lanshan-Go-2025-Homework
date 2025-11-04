package main

import (
	"fmt"
)

func trans(a []int) map[int]int {
	res := make(map[int]int)
	for i := range a {
		fmt.Print(a[i], "\n")
		res[a[i]]++
	}
	return res
}
func main() {

	fmt.Print(trans([]int{1, 2, 3, 4, 4, 5, 5, 5, 5, 5, 5}))

}
