package robot

import (
	"io"
	"os"

	"github.com/liangdas/armyant/task"
)

type Manager struct {
	Writer io.Writer
}

func init() {

}

func (this *Manager) writer() io.Writer {
	if this.Writer == nil {
		return os.Stdout
	}
	return this.Writer
}

func (this *Manager) Finish(task task.Task) {

}

func (this *Manager) CreateWork() task.Work {

	return NewWork(this)
}

func NewManager(t task.Task) task.WorkManager {
	this := new(Manager)
	return this
}
