package transactional

import "github.com/Luna-CY/Golang-Project-Template/internal/context"

func (cls *Transactional) AddRollbackHooks(hooks ...func(ctx context.Context)) {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	cls.rollbackHooks = append(cls.rollbackHooks, hooks...)
}
