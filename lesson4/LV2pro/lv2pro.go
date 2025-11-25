package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, _ := os.OpenFile("log.txt", os.O_APPEND, 0666)
	ope, _ := os.OpenFile("act.txt", os.O_RDONLY, 0666)
	reader := bufio.NewReader(ope)
	//file.WriteString(file,"\n")
	for b, _, _ := reader.ReadLine(); b != nil; b, _, _ = reader.ReadLine() {
		//b, _, _ := reader.ReadLine()
		fmt.Fprintf(file, "%s %s\n", time.Now().String(), b)

		//fmt.Println("%s %s", time.Now().String(), string(b))
		//不是我修得来的爆黄，直接注释
		//更暴力的同步输出（）

	}

}
