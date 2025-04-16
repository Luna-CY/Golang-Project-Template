package example_service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao/option"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

func (cls *Example) ListBySimpleCondition(ctx context.Context, page int, size int, options ...option.ExampleOption) (int64, []*model.Example, errors.Error) {
	total, data, err := cls.example.FindExampleBySimpleCondition(ctx, page, size, options...)
	if nil != err {
		return 0, nil, err.Relation(errors.ErrorServerInternalError("ISE_LE.E_LE.LBSC_ON.133214"))
	}

	return total, data, nil
}
