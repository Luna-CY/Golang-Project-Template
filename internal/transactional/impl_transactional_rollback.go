package transactional

import "github.com/Luna-CY/Golang-Project-Template/internal/context"

func (cls *Transactional) Rollback(ctx context.Context) error {
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
		return err
	}

	return nil
}
