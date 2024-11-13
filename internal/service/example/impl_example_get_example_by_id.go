package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/ierror"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *Example) GetExampleById(ctx icontext.Context, id uint64, lock bool) (*model.Example, error) {
	if 0 == id {
		logger.SugarLogger(ctx).Errorf("I.S.Example.GetExampleById: id is 0 stack %s", string(debug.Stack()))

		return nil, ierror.ErrorServerInternalError
	}

	return cls.example.TakeExampleById(ctx, id, lock)
}
