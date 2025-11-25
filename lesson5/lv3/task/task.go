package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func DoTask(fileCh <-chan string, resultCh chan<- string, keyword string) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for filePath := range fileCh {
				//fmt.Printf("正在搜索文件: %s\n", filePath)
				searchInFile(filePath, keyword, resultCh)
			}
		}()
	}
	//封装！！！
	wg.Wait()
}

func searchInFile(filePath string, keyword string, resultChan chan<- string) {
	/*file, err := os.Open(filePath)
	if err != nil {
		return
	}*/
	file, _ := os.OpenFile(filePath, os.O_RDONLY, 0666)
	defer file.Close()
	scanner := bufio.NewReader(file)
	lineNum := 1
	for b, _, _ := scanner.ReadLine(); b != nil; b, _, _ = scanner.ReadLine() {
		line := string(b)
		//fmt.Println(line)
		if strings.Contains(line, keyword) {
			resultChan <- fmt.Sprintf("%s (第%d行)\n%s", filePath, lineNum, line)
			continue
		}
		lineNum++
	}
}
