package transactional

import "github.com/Luna-CY/Golang-Project-Template/internal/context"

func (cls *Transactional) Commit(ctx context.Context) error {
	cls.mutex.Lock()
	defer cls.mutex.Unlock()

	if cls.flag {
		return nil
	}

	cls.flag = true
	if err := cls.db.Commit().Error; nil != err {
		return err
	}

	// if commit successful, call on commit hooks
	for _, hook := range cls.commitHooks {
		hook(ctx)
	}

	return nil
}
