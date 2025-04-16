package example_dao

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao/option"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

func (cls *Example) FindExampleBySimpleCondition(ctx context.Context, page int, size int, options ...option.ExampleOption) (int64, []*model.Example, errors.Error) {
	var session = cls.GetDb(ctx).Model(new(model.Example))

	var joinTables = make(map[string]struct{})
	for _, option := range options {
		session = option(session, joinTables)
	}

	var total int64
	if err := session.Count(&total).Error; nil != err {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I.D.Example.FindExampleBySimpleCondition count failed, err %v", err)

		return 0, nil, errors.ErrorServerInternalError("IDE_LE.E_LE.FEBSC_ON.232225")
	}

	if 0 == total || 0 == size || int64((page-1)*size) >= total {
		return total, nil, nil
	}

	var data []*model.Example
	if err := session.Offset((page - 1) * size).Limit(size).Find(&data).Error; nil != err {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I.D.Example.FindExampleBySimpleCondition find failed, err %v", err)

		return 0, nil, errors.ErrorServerInternalError("IDE_LE.E_LE.FEBSC_ON.342225")
	}

	return total, data, nil
}
