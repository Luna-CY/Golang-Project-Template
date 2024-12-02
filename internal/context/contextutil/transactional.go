package contextutil

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/transactional"
)

func CheckOnTransactional(ctx context.Context) bool {
	return ctx.Value("transactional") != nil
}

func GetTransactional(ctx context.Context) (transactional.Transactional, bool) {
	if !CheckOnTransactional(ctx) {
		return nil, false
	}

	return ctx.Value("transactional").(transactional.Transactional), true
}

func SetTransactional(ctx context.Context, transactional transactional.Transactional) context.Context {
	return NewContextWithValue(ctx, "transactional", transactional)
}
