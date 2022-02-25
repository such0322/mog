package test_task

import (
	"github.com/liangdas/armyant/task"
	"io"
	"os"
)

type Manager struct {
	// Writer is where results will be written. If nil, results are written to stdout.
	Writer io.Writer
}

func (this *Manager) writer() io.Writer {
	if this.Writer == nil {
		return os.Stdout
	}
	return this.Writer
}
func (this *Manager) Finish(task task.Task) {
	//total := time.Now().Sub(task.Start)
}
func (this *Manager) CreateWork() task.Work {
	return NewWork(this)
}

// Run makes all the requests, prints the summary. It blocks until
// all work is done.
func NewManager(t task.Task) task.WorkManager {
	// append hey's user agent
	this := new(Manager)
	return this
}
