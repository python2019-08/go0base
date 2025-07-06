package many2many

import "go0base/goprjs/05producer-consumer/out"

type Task struct {
	ID int64
}

func (t *Task) run() {
	out.Println(t.ID)
}

var taskCh = make(chan Task, 10)
