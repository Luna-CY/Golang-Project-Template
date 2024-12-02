package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
	"time"
)

func (cls *Example) SaveExample(ctx context.Context, example *model.Example) error {
	if nil == example {
		logger.SugarLogger(ctx).Errorf("I.D.Example.SaveExample example is nil stack %s", string(debug.Stack()))

		return errors.ErrorServerInternalError
	}

	example.UpdateTime = pointer.New(time.Now().Unix())
	if 0 == example.Id {
		example.CreateTime = pointer.New(time.Now().Unix())

		if err := cls.GetDb(ctx).Model(new(model.Example)).Create(&example).Error; nil != err {
			logger.SugarLogger(ctx).Errorf("I.D.Example.SaveExample create example failed, err %v, stack %s", err, string(debug.Stack()))

			return errors.ErrorServerInternalError
		}

		return nil
	}

	if err := cls.GetDb(ctx).Model(new(model.Example)).Where("id = ?", example.Id).Updates(&example).Error; nil != err {
		logger.SugarLogger(ctx).Errorf("I.D.Example.SaveExample save example failed, err %v, stack %s", err, string(debug.Stack()))

		return errors.ErrorServerInternalError
	}

	return nil
}
