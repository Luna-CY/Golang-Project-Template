package transactional

import "github.com/Luna-CY/Golang-Project-Template/internal/context"

func (cls *Transactional) AddCommitHooks(hooks ...func(ctx context.Context)) {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	cls.commitHooks = append(cls.commitHooks, hooks...)
}
