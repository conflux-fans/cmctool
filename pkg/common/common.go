package common

import (
	"time"
)

func Retry(count int, interval time.Duration, fn func() error) error {
	for i := 0; i < count; i++ {
		err := fn()
		if err != nil {
			// logrus.WithError(err).WithField("stack", string(debug.Stack())).WithField("retry cnt", i).Info("retry function error")
			if i == count-1 {
				return err
			} else {
				time.Sleep(interval)
				continue
			}
		}
		return nil
	}
	return nil
}

func Must[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

func MustNil(err error) {
	if err != nil {
		panic(err)
	}
}
