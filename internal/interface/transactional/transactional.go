package transactional

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"gorm.io/gorm"
)

type Transactional interface {
	// Session get db object of gorm
	Session() *gorm.DB

	// Rollback rollback transaction, it should always be safe and called as a defer call
	Rollback(ctx context.Context) error

	// Commit commit transaction
	Commit(ctx context.Context) error

	// AddCommitHooks add hooks when commit transaction
	AddCommitHooks(hooks ...func(ctx context.Context))

	// AddRollbackHooks add hooks when rollback transaction
	AddRollbackHooks(hooks ...func(ctx context.Context))
}
