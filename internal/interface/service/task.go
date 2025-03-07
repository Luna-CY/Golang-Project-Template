package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"time"
)

type Task interface {
	// StartOneTimeTask start task
	// caller will be executed in a go routine, and any external temporary variables or pointers should not be referenced, and all parameters are passed through values and used
	StartOneTimeTask(ctx context.Context, tag string, values map[string]any, caller func(ctx context.Context, values map[string]any, progress func(int64)) errors.Error, timeout time.Duration, unique bool) (string, errors.Error)

	// GetOneTimeTaskState get task state
	// if task not found return errors.ErrorRecordNotFound error
	GetOneTimeTaskState(ctx context.Context, taskId string) (processing bool, progress int64, err errors.Error)
}
