package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	"github.com/Luna-CY/Golang-Project-Template/internal/util/pointer"
	"github.com/Luna-CY/Golang-Project-Template/model"
	"runtime/debug"
)

func (cls *Example) UpdateExample(ctx context.Context, example *model.Example, field1 *string, field2 *uint64, field3 *bool, field4 *model.ExampleEnumFieldType) errors.Error {
	if nil == example {
		logger.SugarLogger(ctx).Errorf("I.S.Example.UpdateExample: example is nil stack %s", string(debug.Stack()))

		return errors.ErrorServerInternalError("IS.E_LE.UE_LE.16")
	}

	example.Field1 = pointer.Or(field1, example.Field1)
	example.Field2 = pointer.Or(field2, example.Field2)
	example.Field3 = pointer.Or(field3, example.Field3)
	example.Field4 = pointer.Or(field4, example.Field4)

	if err := cls.example.SaveExample(ctx, example); nil != err {
		return err.Relation(errors.ErrorServerInternalError("IS.E_LE.UE_LE.25"))
	}

	return nil
}
