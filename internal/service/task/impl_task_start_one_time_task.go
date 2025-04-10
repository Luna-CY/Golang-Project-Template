package task

import (
	"sync/atomic"
	"time"

	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	gonanoid "github.com/matoous/go-nanoid"
)

func (cls *Task) StartOneTimeTask(ctx context.Context, tag string, values map[string]any, caller func(ctx context.Context, values map[string]any, progress func(int64)) errors.Error, timeout time.Duration, unique bool) (string, errors.Error) {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	if unique {
		for id, task := range cls.tasks {
			if task.tag == tag && 1 == atomic.LoadInt32(&task.processing) {
				return "", errors.New(errors.ErrorTypeServerInternalError, "IST_SK.T_SK.SOTT_SK.20", "type %s only one can be executed at the same time, current running task id: %s", tag, id)
			}
		}
	}

	var id = gonanoid.MustID(21)

	cls.tasks[id] = &oneTimeTask{
		ctx:         ctx,
		tag:         tag,
		values:      values,
		caller:      caller,
		timeout:     timeout,
		initialized: 0,
		processing:  0,
		progress:    0,
		error:       nil,
	}

	go cls.run(cls.tasks[id])

	return id, nil
}

func (cls *Task) run(task *oneTimeTask) {
	defer func() {
		var err = recover()
		if nil != err {
			logger.SugarLogger(logger.WithRequestId(task.ctx), logger.WithStack()).Errorf("I.S.Task.run task run faillback: %s", err)
		}
	}()

	if !atomic.CompareAndSwapInt32(&task.initialized, 0, 1) {
		return
	}

	atomic.StoreInt32(&task.processing, 1)
	defer func() {
		atomic.StoreInt32(&task.processing, 0)
	}()

	var ctx, cancel = contextutil.NewContextWithTimeout(task.ctx, task.timeout)
	defer cancel()

	task.error = task.caller(ctx, task.values, func(progress int64) {
		atomic.StoreInt64(&task.progress, progress)
	})
}
