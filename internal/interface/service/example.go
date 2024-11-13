package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

type Example interface {
	Transactional

	// CreateExample create example
	CreateExample(ctx icontext.Context, field1 string, field2 uint64, field3 bool, field4 model.ExampleEnumFieldType) (*model.Example, error)

	// UpdateExample update example
	// if field value is nil, it will not update this field
	UpdateExample(ctx icontext.Context, example *model.Example, field1 *string, field2 *uint64, field3 *bool, field4 *model.ExampleEnumFieldType) error

	// GetExampleById get example by id
	// if example not found, return ierror.ErrorRecordNotFound error
	GetExampleById(ctx icontext.Context, id uint64, lock bool) (*model.Example, error)

	// ListBySimpleCondition get examples by simple condition
	// if field4 is nil, it will not filter by this field in the query.
	ListBySimpleCondition(ctx icontext.Context, field4 *model.ExampleEnumFieldType, page int, size int) (int64, []*model.Example, error)
}
