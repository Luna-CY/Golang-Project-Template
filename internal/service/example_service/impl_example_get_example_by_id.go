package example_service

import (
	"runtime/debug"

	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao/option"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

func (cls *Example) GetExampleById(ctx context.Context, id uint64, options ...option.ExampleOption) (*model.Example, errors.Error) {
	if 0 == id {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I.S.Example.GetExampleById: id is 0 stack %s", string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("ISE_LE.E_LE.GEBI_ID.173201")
	}

	example, err := cls.example.TakeExampleById(ctx, id, options...)
	if nil != err {
		return nil, err.Relation(errors.ErrorServerInternalError("ISE_LE.E_LE.GEBI_ID.223205"))
	}

	return example, nil
}
