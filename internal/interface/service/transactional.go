package service

import "github.com/Luna-CY/Golang-Project-Template/internal/icontext"

type Transactional interface {
	// WithTransaction begin a transaction and call the provided function, return error if any error occurred
	WithTransaction(ctx icontext.Context, call func(ctx icontext.Context) error) error
}
