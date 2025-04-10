package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

func (cls *Example) FindExampleBySimpleCondition(ctx context.Context, field4 *model.ExampleEnumFieldType, page int, size int) (int64, []*model.Example, errors.Error) {
	var session = cls.GetDb(ctx).Model(new(model.Example))
	session = dao.GormWhereEqualWithNotNil(session, "field4", field4)

	var total int64
	if err := session.Count(&total).Error; nil != err {
		logger.SugarLogger(ctx, logger.WithStack()).Errorf("I.D.Example.FindExampleBySimpleCondition count failed, err %v", err)

		return 0, nil, errors.ErrorServerInternalError("IDE_LE.E_LE.FEBSC.19")
	}

	if 0 == total || 0 == size || int64((page-1)*size) >= total {
		return total, nil, nil
	}

	var data []*model.Example
	if err := session.Offset((page - 1) * size).Limit(size).Order("id desc").Find(&data).Error; nil != err {
		logger.SugarLogger(ctx, logger.WithStack()).Errorf("I.D.Example.FindExampleBySimpleCondition find failed, err %v", err)

		return 0, nil, errors.ErrorServerInternalError("IDE_LE.E_LE.FEBSC.30")
	}

	return total, data, nil
}
