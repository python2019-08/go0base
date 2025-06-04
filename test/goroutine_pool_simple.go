package test

import (
	"fmt"
	"sync"
)

func task(id int) {

	fmt.Println(id)
}

func Simple_goroutine_pool() {
	numTasks := 100
	concurrency := 5

	var wg sync.WaitGroup
	taskch := make(chan int)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for taskID := range taskch {
				task(taskID)
			}
		}()
	}

	for i := 0; i < numTasks; i++ {
		taskch <- i
	}

	close(taskch)
	wg.Wait()
	fmt.Println("完成100个任务")
}
