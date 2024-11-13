package transactional

import "github.com/Luna-CY/Golang-Project-Template/internal/icontext"

func (cls *Transactional) AddCommitHooks(hooks ...func(ctx icontext.Context)) {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	cls.commitHooks = append(cls.commitHooks, hooks...)
}
