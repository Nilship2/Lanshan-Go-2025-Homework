package main

import (
	"fmt"
	"nilmod/task"
)

func main() {
	var a int
	var b int
	var c int
	fmt.Scanf("%d%d%d", &a, &b, &c)
	var result int = task.Add_task(a, b, c)
	fmt.Printf("%d\n", result)
}
