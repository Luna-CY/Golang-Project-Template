package example

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

func (cls *Example) ListBySimpleCondition(ctx context.Context, field4 *model.ExampleEnumFieldType, page int, size int) (int64, []*model.Example, error) {
	return cls.example.FindExampleBySimpleCondition(ctx, field4, page, size)
}
