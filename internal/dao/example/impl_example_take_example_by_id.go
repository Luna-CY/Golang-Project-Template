package example

import (
	"runtime/debug"

	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"gorm.io/gorm"
)

func (cls *Example) TakeExampleById(ctx context.Context, id uint64, lock bool) (*model.Example, errors.Error) {
	if 0 == id {
		logger.SugarLogger(ctx, logger.WithStack()).Errorf("I.D.Example.TakeExampleById id is 0 stack %s", string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("IDE_LE.E_LE.TEBI.17")
	}

	var session = cls.GetDb(ctx).Model(new(model.Example))
	session = dao.Lock(session, lock)

	var example *model.Example
	if err := session.Where("id = ?", id).Take(&example).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrorRecordNotFound("IDE_LE.E_LE.TEBI.26")
		}

		logger.SugarLogger(ctx, logger.WithStack()).Errorf("I.D.Example.TakeExampleById take example by id failed, err %v, stack %s", err, string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("IDE_LE.E_LE.TEBI.31")
	}

	return example, nil
}
