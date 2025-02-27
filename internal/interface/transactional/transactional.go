package transactional

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"gorm.io/gorm"
)

type Transactional interface {
	// Session get db object of gorm
	Session() *gorm.DB

	// Rollback rollback transaction, it should always be safe and called as a defer call
	Rollback(ctx context.Context) errors.Error

	// Commit commit transaction
	Commit(ctx context.Context) errors.Error

	// AddCommitHooks add hooks when commit transaction
	AddCommitHooks(hooks ...func(ctx context.Context))

	// AddRollbackHooks add hooks when rollback transaction
	AddRollbackHooks(hooks ...func(ctx context.Context))
}
