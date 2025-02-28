package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *Example) CreateExample(ctx context.Context, field1 string, field2 uint64, field3 bool, field4 model.ExampleEnumFieldType) (*model.Example, errors.Error) {
	if "" == field1 || 0 == field2 {
		logger.SugarLogger(ctx).Errorf("I.S.Example.CreateExample: field1 is %s field2 is %d stack %s", field1, field2, string(debug.Stack()))

		return nil, errors.ErrorServerInternalError("ISE_LE.E_LE.CE_LE.16")
	}

	var example = new(model.Example)
	example.Field1 = pointer.New(field1)
	example.Field2 = pointer.New(field2)
	example.Field3 = pointer.New(field3)
	example.Field4 = pointer.New(field4)

	if err := cls.example.SaveExample(ctx, example); nil != err {
		return nil, err.Relation(errors.ErrorServerInternalError("ISE_LE.E_LE.CE_LE.26"))
	}

	return example, nil
}
