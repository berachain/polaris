package errors

import "fmt"

func Wrap(err error, desc string) error {
	return fmt.Errorf("%w: %s", err, desc)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return fmt.Errorf("%w: %s", err, fmt.Sprintf(format, args...))
}
