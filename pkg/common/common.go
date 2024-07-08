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
