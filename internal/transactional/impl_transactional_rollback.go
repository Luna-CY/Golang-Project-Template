package transactional

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
)

func (cls *Transactional) Rollback(ctx context.Context) errors.Error {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	if cls.flag {
		return nil
	}

	cls.flag = true
	defer func() {
		// call on defer
		for _, hook := range cls.rollbackHooks {
			hook(ctx)
		}
	}()

	if err := cls.db.Rollback().Error; nil != err {
		return errors.New(errors.ErrorTypeServerInternalError, "IT.R_CK.25", err)
	}

	return nil
}
