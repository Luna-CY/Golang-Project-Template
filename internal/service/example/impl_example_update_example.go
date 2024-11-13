package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/ierror"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *Example) UpdateExample(ctx icontext.Context, example *model.Example, field1 *string, field2 *uint64, field3 *bool, field4 *model.ExampleEnumFieldType) error {
	if nil == example {
		logger.SugarLogger(ctx).Errorf("I.S.Example.UpdateExample: example is nil stack %s", string(debug.Stack()))

		return ierror.ErrorServerInternalError
	}

	example.Field1 = pointer.Or(field1, example.Field1)
	example.Field2 = pointer.Or(field2, example.Field2)
	example.Field3 = pointer.Or(field3, example.Field3)
	example.Field4 = pointer.Or(field4, example.Field4)

	return cls.example.SaveExample(ctx, example)
}
