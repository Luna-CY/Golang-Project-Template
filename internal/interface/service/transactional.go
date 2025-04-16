package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
)

type Transactional interface {
	// WithTransaction 开始一个事务并调用提供的函数，如果任何错误发生，返回 errors.Error
	WithTransaction(ctx context.Context, call func(ctx context.Context) errors.Error) errors.Error
}
