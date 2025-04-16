package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao/option"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

type Example interface {
	Transactional

	// CreateExample 创建 example
	CreateExample(ctx context.Context, field1 string, field2 uint64, field3 bool, field4 model.ExampleEnumFieldType) (*model.Example, errors.Error)

	// UpdateExample 更新 example
	// 如果字段值为 nil，则不更新此字段
	UpdateExample(ctx context.Context, example *model.Example, field1 *string, field2 *uint64, field3 *bool, field4 *model.ExampleEnumFieldType) errors.Error

	// GetExampleById 从 db 获取 example 通过 id
	// 如果 example 不存在，返回 errors.ErrorRecordNotFound 错误
	GetExampleById(ctx context.Context, id uint64, options ...option.ExampleOption) (*model.Example, errors.Error)

	// ListBySimpleCondition 通过简单条件从 db 获取 example
	// 如果 field4 为 nil，则不在此查询中过滤此字段
	ListBySimpleCondition(ctx context.Context, page int, size int, options ...option.ExampleOption) (int64, []*model.Example, errors.Error)
}
