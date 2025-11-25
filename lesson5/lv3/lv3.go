package main

import (
	"fmt"
	"nilmod/task"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	aimPath := "./source"
	fileCh := make(chan string, 100)
	resultCh := make(chan string, 100)

	var resultWg sync.WaitGroup

	resultWg.Add(1)
	go func() {
		defer resultWg.Done()
		for result := range resultCh {
			fmt.Println("")
			fmt.Printf("%s\n", result)
			fmt.Println("")
		}
	}()

	go func() {
		defer close(fileCh)
		err := filepath.Walk(aimPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				fileCh <- path
			}
			return nil
		})
		if err != nil {
			fmt.Printf("遍历失败")
		}
	}()
	fmt.Println("要查找什么关键词呢：")
	var keyword string
	fmt.Scanf("%s", &keyword)
	task.DoTask(fileCh, resultCh, keyword)
	close(resultCh)
	resultWg.Wait()
}
