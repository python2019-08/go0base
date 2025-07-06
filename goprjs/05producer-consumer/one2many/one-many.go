package one2many

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

func producer(wo chan<- Task) {
	var i int64
	for i = 1; i < taskNum; i++ {
		t := Task{
			ID: i,
		}
		wo <- t
	}

	// 单个生产者 完成生产后，可以关闭通道
	close(wo)
}

func consumer(ro <-chan Task) {
	for t := range ro {
		if t.ID != 0 {
			t.run()
		}
	}
}

func Exec() {
	// A WaitGroup must not be copied after first use.
	wg := &sync.WaitGroup{} // wg 不允许copy，否则会死锁

	// ---- one producer
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		producer(taskCh)

	}(wg)

	// ----many consumers
	var i int64                   //one2many
	for i = 0; i < taskNum; i++ { //one2many
		if i%100 == 0 { //one2many
			wg.Add(1)

			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				consumer(taskCh)
			}(wg)
		} //one2many
	} //one2many

	// ----
	// 如果有goroutine没有退出，此处wg.Wait()会有死锁
	wg.Wait()
	out.Println("执行成功")
}
