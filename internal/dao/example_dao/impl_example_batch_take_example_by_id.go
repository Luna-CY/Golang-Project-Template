package example_dao

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao/option"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *Example) BatchTakeExampleById(ctx context.Context, values []uint64, options ...option.ExampleOption) ([]*model.Example, errors.Error) {
	if 0 == len(values) {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I.D.Example.BatchTakeById: values is empty stack %s", string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("IDE_LE.E_LE.BTEBI_ID.162225")
	}

	var session = cls.GetDb(ctx).Model(new(model.Example))

	var joinTables = make(map[string]struct{})
	for _, option := range options {
		session = option(session, joinTables)
	}

	var data []*model.Example
	if err := session.Where("id in (?)", values).Find(&data).Error; nil != err {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I.D.Example.BatchTakeById: batch take by id failed, err %v, stack %s", err, string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("IDE_LE.E_LE.BTEBI_ID.302225")
	}

	return data, nil
}
