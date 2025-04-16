package transactional

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"gorm.io/gorm"
)

type Transactional interface {
	// Session 获取 gorm 的 db 对象
	Session() *gorm.DB

	// Rollback 回滚事务，应该总是安全的，并作为 defer 调用
	Rollback(ctx context.Context) errors.Error

	// Commit 提交事务
	Commit(ctx context.Context) errors.Error

	// AddCommitHooks 在提交事务时添加钩子
	AddCommitHooks(hooks ...func(ctx context.Context))

	// AddRollbackHooks 在回滚事务时添加钩子
	AddRollbackHooks(hooks ...func(ctx context.Context))
}
