package example_dao

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao/option"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"gorm.io/gorm"
	"runtime/debug"
)

func (cls *Example) TakeExampleById(ctx context.Context, id uint64, options ...option.ExampleOption) (*model.Example, errors.Error) {
	if 0 == id {
		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I.D.Example.TakeById id is %v stack %s", id, string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("IDE_LE.E_LE.TEBI_ID.172225")
	}

	var session = cls.GetDb(ctx).Model(new(model.Example))

	var joinTables = make(map[string]struct{})
	for _, option := range options {
		session = option(session, joinTables)
	}

	var example *model.Example
	if err := session.Where("id = ?", id).Take(&example).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrorRecordNotFound("IDE_LE.E_LE.TEBI_ID.302225")
		}

		logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I.D.Example.TakeExampleById take example by id failed, err %v, stack %s", err, string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("IDE_LE.E_LE.TEBI_ID.352225")
	}

	return example, nil
}
