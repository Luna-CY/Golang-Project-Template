package icontextutil

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/transactional"
)

func CheckOnTransactional(ctx icontext.Context) bool {
	return ctx.Value("transactional") != nil
}

func GetTransactional(ctx icontext.Context) (transactional.Transactional, bool) {
	if !CheckOnTransactional(ctx) {
		return nil, false
	}

	return ctx.Value("transactional").(transactional.Transactional), true
}

func SetTransactional(ctx icontext.Context, transactional transactional.Transactional) icontext.Context {
	return NewContextWithValue(ctx, "transactional", transactional)
}
