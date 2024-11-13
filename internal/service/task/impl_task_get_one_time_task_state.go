package task

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/ierror"
	"sync/atomic"
)

func (cls *Task) GetOneTimeTaskState(ctx icontext.Context, taskId string) (processing bool, progress int64, err error) {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	task, ok := cls.tasks[taskId]
	if !ok {
		return false, 0, ierror.ErrorRecordNotFound
	}

	processing = 1 == atomic.LoadInt32(&task.processing)
	progress = atomic.LoadInt64(&task.progress)

	return
}
