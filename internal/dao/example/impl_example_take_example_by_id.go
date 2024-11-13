package example

import (
	"errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/ierror"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"gorm.io/gorm"
	"runtime/debug"
)

func (cls *Example) TakeExampleById(ctx icontext.Context, id uint64, lock bool) (*model.Example, error) {
	if 0 == id {
		logger.SugarLogger(ctx).Errorf("I.D.Example.TakeExampleById id is 0 stack %s", string(debug.Stack()))

		return nil, ierror.ErrorServerInternalError
	}

	var session = cls.GetDb(ctx).Model(new(model.Example))
	session = dao.Lock(session, lock)

	var example *model.Example
	if err := session.Where("id = ?", id).Take(&example).Error; nil != err {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ierror.ErrorRecordNotFound
		}

		logger.SugarLogger(ctx).Errorf("I.D.Example.TakeExampleById take example by id failed, err %v, stack %s", err, string(debug.Stack()))

		return nil, ierror.ErrorServerInternalError
	}

	return example, nil
}
