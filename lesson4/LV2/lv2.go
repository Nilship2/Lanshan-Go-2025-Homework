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

	for b, _, _ := reader.ReadLine(); b != nil; b, _, _ = reader.ReadLine() {
		//b, _, _ := reader.ReadLine()
		if string(b) == "endl" {
			fmt.Printf("%s %s", time.Now().String(), b)
			break
		}
		fmt.Fprintf(file, "%s %s", time.Now().String(), b)
		file.WriteString("\n")
	}

}
