package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

func (cls *Example) ListBySimpleCondition(ctx context.Context, field4 *model.ExampleEnumFieldType, page int, size int) (int64, []*model.Example, errors.Error) {
	total, data, err := cls.example.FindExampleBySimpleCondition(ctx, field4, page, size)
	if nil != err {
		return 0, nil, err.Relation(errors.ErrorServerInternalError("IS.E_LE.LBSC_ON.12"))
	}

	return total, data, nil
}
