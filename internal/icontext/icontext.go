package icontext

import (
	"context"
	"time"
)

type Context interface {
	context.Context

	notAllowImplement()
}

type IContext struct {
	Context context.Context
}

func (cls *IContext) notAllowImplement() {}

func (cls *IContext) Deadline() (deadline time.Time, ok bool) {
	return cls.Context.Deadline()
}

func (cls *IContext) Done() <-chan struct{} {
	return cls.Context.Done()
}

func (cls *IContext) Err() error {
	return cls.Context.Err()
}

func (cls *IContext) Value(key any) any {
	return cls.Context.Value(key)
}
