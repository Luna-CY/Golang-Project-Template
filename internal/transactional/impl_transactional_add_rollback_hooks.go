package transactional

import "github.com/Luna-CY/Golang-Project-Template/internal/icontext"

func (cls *Transactional) AddRollbackHooks(hooks ...func(ctx icontext.Context)) {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	cls.rollbackHooks = append(cls.rollbackHooks, hooks...)
}
