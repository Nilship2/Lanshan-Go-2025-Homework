package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()

	os.OpenFile("in.txt", os.O_RDONLY, 0666)
	writer, _ := os.OpenFile("out.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	//file, _ := os.ReadFile("in.txt")

	//file, err := os.Create("out.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer file.Close()

	//writer := bufio.NewWriter(file)
	flg := 1
	for i := 0; i < 2100000; i++ {
		s := strconv.Itoa(i) + "\n"
		nn, err := writer.WriteString(s)
		if flg == 1 && err != nil {
			fmt.Println(nn, err)
			flg = 0
		}

	}
	t2 := time.Since(t1)
	//fmt.Println(">>>", t1, " ", t2)
	fmt.Printf("Time taken: %v\n", t2)
	//err = writer.Flush()

}
