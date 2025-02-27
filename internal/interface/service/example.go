package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

type Example interface {
	Transactional

	// CreateExample create example
	CreateExample(ctx context.Context, field1 string, field2 uint64, field3 bool, field4 model.ExampleEnumFieldType) (*model.Example, errors.Error)

	// UpdateExample update example
	// if field value is nil, it will not update this field
	UpdateExample(ctx context.Context, example *model.Example, field1 *string, field2 *uint64, field3 *bool, field4 *model.ExampleEnumFieldType) errors.Error

	// GetExampleById get example by id
	// if example not found, return errors.ErrorRecordNotFound error
	GetExampleById(ctx context.Context, id uint64, lock bool) (*model.Example, errors.Error)

	// ListBySimpleCondition get examples by simple condition
	// if field4 is nil, it will not filter by this field in the query.
	ListBySimpleCondition(ctx context.Context, field4 *model.ExampleEnumFieldType, page int, size int) (int64, []*model.Example, errors.Error)
}
