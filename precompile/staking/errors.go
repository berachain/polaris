package staking

import "errors"

var (
	ErrInvalidValidatorAddr = errors.New("invalid validator address")
	ErrInvalidString        = errors.New("invalid string")
	ErrInvalidBigInt        = errors.New("invalid big int")
	ErrInvalidUint64        = errors.New("invalid uint64")
	ErrInvalidInt64         = errors.New("invalid int64")
)
