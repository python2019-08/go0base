package many2one

import (
	"go0base/goprjs/05producer-consumer/out"
	"sync"
)

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)

const taskNum int64 = 10000

var nums int64 = 100

func producer(wo chan<- Task, startNum int64, nums int64) {
	var i int64
	for i = startNum; i < startNum+nums; i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}

}

func consumer(ro <-chan Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	wg := &sync.WaitGroup{}

	//---- many producers
	pwg := &sync.WaitGroup{} // 用于判断所有的producer是否都结束
	var i int64
	for i = 0; i < taskNum; i += nums {
		if i >= taskNum {
			break
		}
		wg.Add(1)
		pwg.Add(1)

		go func(i int64) {
			defer wg.Done()
			defer pwg.Done()
			producer(taskCh, i, nums)
		}(i)
	}

	// ----one consumer
	wg.Add(1)
	go func() {
		defer wg.Done()
		consumer(taskCh)
	}()

	// ----
	pwg.Wait()
	// 所有的producer是否都结束后，close taskCh
	close(taskCh)

	// 如果有goroutine没有退出，此处wg.Wait()会有死锁
	wg.Wait()
	out.Println("执行成功")
}
