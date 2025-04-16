package dao

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao/option"
	"github.com/Luna-CY/Golang-Project-Template/model"
)

type Example interface {
	Transactional

	// SaveExample 保存 example 到 db
	// 如果 example id 为 0，则创建新记录，否则更新记录
	SaveExample(ctx context.Context, example *model.Example) errors.Error

	// TakeExampleById 从 db 获取 example 通过 id
	// 如果 example 不存在，返回错误类型 with errors.ErrorTypeRecordNotFound
	TakeExampleById(ctx context.Context, id uint64, options ...option.ExampleOption) (*model.Example, errors.Error)

	// BatchTakeExampleById 从 db 获取 example 通过 id 列表
	// 如果 example 不存在，返回错误类型 with errors.ErrorTypeRecordNotFound
	BatchTakeExampleById(ctx context.Context, values []uint64, options ...option.ExampleOption) ([]*model.Example, errors.Error)

	// FindExampleBySimpleCondition 通过简单条件从 db 获取 example
	FindExampleBySimpleCondition(ctx context.Context, page int, size int, options ...option.ExampleOption) (int64, []*model.Example, errors.Error)
}
