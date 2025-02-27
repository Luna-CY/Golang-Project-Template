package retry

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"time"
)

func MaxRetry(max int, internal time.Duration, f func() errors.Error) errors.Error {
	var err errors.Error

	for i := 0; i < max; i++ {
		err = f()
		if nil == err {
			return nil
		}

		time.Sleep(internal)
	}

	return err
}
