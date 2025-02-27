package task

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"sync"
	"time"
)

type oneTimeTask struct {
	ctx         context.Context
	tag         string
	values      map[string]any
	caller      func(ctx context.Context, values map[string]any, progress func(int64)) errors.Error
	timeout     time.Duration
	initialized int32
	processing  int32
	progress    int64
	error       error
}

var to sync.Once
var ti *Task

func New() *Task {
	to.Do(func() {
		ti = &Task{
			mutex: sync.Mutex{},
			tasks: make(map[string]*oneTimeTask),
		}
	})

	return ti
}

type Task struct {
	mutex sync.Mutex
	tasks map[string]*oneTimeTask
}
