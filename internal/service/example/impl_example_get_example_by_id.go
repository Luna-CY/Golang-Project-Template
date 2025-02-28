package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *Example) GetExampleById(ctx context.Context, id uint64, lock bool) (*model.Example, errors.Error) {
	if 0 == id {
		logger.SugarLogger(ctx).Errorf("I.S.Example.GetExampleById: id is 0 stack %s", string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("ISE_LE.E_LE.GEBI_ID.15")
	}

	example, err := cls.example.TakeExampleById(ctx, id, lock)
	if nil != err {
		return nil, err.Relation(errors.ErrorServerInternalError("ISE_LE.E_LE.GEBI_ID.20"))
	}

	return example, nil
}
