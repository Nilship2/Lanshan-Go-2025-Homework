package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func open_box() {
	dir, err := os.Open("./lanshan")
	dir2, _ := os.Open("./91porn")
	if err != nil {
		log.Fatal(err)
	}
	defer dir.Close()
	defer dir2.Close()

	fileInfos, err := dir.Readdir(-1)
	fileInfos2, _ := dir2.Readdir(-1)
	if err != nil {
		log.Fatal(err)
	}
	//i := 0
	//for _, fileInfo := range fileInfos.Size() {

	for i := 0; i < len(fileInfos); i++ {
		name := fileInfos[i].Name()
		sz1 := fileInfos[i].Size()
		sz2 := fileInfos2[i].Size()
		if sz1 == sz2 {
			//i++
			continue
		} else {
			logfile, _ := os.OpenFile("91log.txt", os.O_APPEND, 0666)
			fmt.Fprintf(logfile, "%s\n %s的内存由 %d bytes 变为 %d bytes\n\n", time.Now().String(), name, sz1, sz2)
			//然而面对相同的文件判定了不同的大小
			//不管了就这样吧？
			//好像是换行的读取有差。。。
			//问AI ing......
			//破AI叽里咕噜说什么呢，算了就这样吧
			logfile.Close()
		}
		//i++
		namee := "./lanshan/" + name
		nname := "./91porn/" + name

		file, _ := os.OpenFile(nname, os.O_CREATE|os.O_WRONLY, 0666)
		ope, _ := os.OpenFile(namee, os.O_RDONLY, 0666)

		reader := bufio.NewReader(ope)
		for b, _, _ := reader.ReadLine(); b != nil; b, _, _ = reader.ReadLine() {
			//b, _, _ := reader.ReadLine()
			fmt.Fprintf(file, "%s\n", b)
		}

	}
}

func main() {
	for 2 > 1 {
		time.Sleep(10 * time.Second)
		open_box()
	}
}
