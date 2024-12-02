package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *Example) GetExampleById(ctx context.Context, id uint64, lock bool) (*model.Example, error) {
	if 0 == id {
		logger.SugarLogger(ctx).Errorf("I.S.Example.GetExampleById: id is 0 stack %s", string(debug.Stack()))

		return nil, errors.ErrorServerInternalError
	}

	return cls.example.TakeExampleById(ctx, id, lock)
}
