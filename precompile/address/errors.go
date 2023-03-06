package address

import "errors"

var (
	ErrInvalidHexAddress = errors.New("invalid hex address")
	ErrInvalidString     = errors.New("invalid string address")
)
