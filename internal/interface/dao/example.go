package dao

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

type Example interface {
	Transactional

	// SaveExample save example to db
	// if example id is 0, it will create a new record, otherwise, it will update the record
	SaveExample(ctx context.Context, example *model.Example) error

	// TakeExampleById get example by id from db
	// if example not found, return errors.ErrorRecordNotFound error
	TakeExampleById(ctx context.Context, id uint64, lock bool) (*model.Example, error)

	// FindExampleBySimpleCondition find examples by simple condition from db
	// if field4 is nil, it will not filter by this field in the query.
	FindExampleBySimpleCondition(ctx context.Context, field4 *model.ExampleEnumFieldType, page int, size int) (int64, []*model.Example, error)
}