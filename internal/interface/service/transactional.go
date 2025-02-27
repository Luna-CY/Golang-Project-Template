package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
)

type Transactional interface {
	// WithTransaction begin a transaction and call the provided function, return error if any error occurred
	WithTransaction(ctx context.Context, call func(ctx context.Context) errors.Error) errors.Error
}
