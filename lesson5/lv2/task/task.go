package task

import (
	"sync"
)

func Add_task(a int, b int, c int) int {
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(c)
	d := b / c
	e := b % c
	for i := 0; i < c; i++ {
		go func() {
			defer wg.Done()
			j := 1
			if i < e {
				j = j - 1
			}
			for ; j <= d; j++ {
				lock.Lock()
				a = a + 1
				b = b - 1
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	return a
}
