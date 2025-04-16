package task_service

import (
	"sync/atomic"

	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
)

func (cls *Task) GetOneTimeTaskState(ctx context.Context, taskId string) (processing bool, progress int64, err errors.Error) {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	task, ok := cls.tasks[taskId]
	if !ok {
		return false, 0, errors.ErrorRecordNotFound("IST_SK.T_SK.GOTTS_TE.163122")
	}

	processing = 1 == atomic.LoadInt32(&task.processing)
	progress = atomic.LoadInt64(&task.progress)

	return
}
