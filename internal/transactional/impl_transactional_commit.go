package transactional

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
)

func (cls *Transactional) Commit(ctx context.Context) errors.Error {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	if cls.flag {
		return nil
	}

	cls.flag = true
	if err := cls.db.Commit().Error; nil != err {
		return errors.New(errors.ErrorTypeServerInternalError, "IT.C_IT.18", err)
	}

	// if commit successful, call on commit hooks
	for _, hook := range cls.commitHooks {
		hook(ctx)
	}

	return nil
}
