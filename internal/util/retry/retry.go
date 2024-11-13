package retry

import "time"

func MaxRetry(max int, internal time.Duration, f func() error) error {
	var err error

	for i := 0; i < max; i++ {
		err = f()
		if nil == err {
			return nil
		}

		time.Sleep(internal)
	}

	return err
}
