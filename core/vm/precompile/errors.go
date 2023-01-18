package precompile

import "errors"

var (
	// `ErrNoPrecompileMethodForABIMethod` is returned when no precompile method is provided for a
	// corresponding ABI method.
	ErrNoPrecompileMethodForABIMethod = errors.New("this ABI method does not have a corresponding precompile method")

	// `ErrWrongContainerFactory` is returned when the wrong precompile container factory is used
	// to build a precompile contract.
	ErrWrongContainerFactory = errors.New("this container factory does not support the given precompile contract")

	ErrNotStatelessPrecompile = errors.New("this precompile contract does not implement `PrecompileContainer`")
)
