package service

import (
	"time"

	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
)

type Task interface {
	// StartOneTimeTask 启动一次性任务
	// 调用者将在一个 go 协程中执行，不应引用任何外部临时变量或指针，所有参数都通过值传递并使用
	StartOneTimeTask(ctx context.Context, tag string, values map[string]any, caller func(ctx context.Context, values map[string]any, progress func(int64)) errors.Error, timeout time.Duration, unique bool) (string, errors.Error)

	// GetOneTimeTaskState 获取任务状态
	// 如果任务不存在，返回 errors.ErrorRecordNotFound 错误
	GetOneTimeTaskState(ctx context.Context, taskId string) (processing bool, progress int64, err errors.Error)
}
