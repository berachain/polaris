package precompile

import (
	coreprecompile "pkg.berachain.dev/polaris/eth/core/precompile"

	"pkg.berachain.dev/polaris/eth/accounts/abi"
	"pkg.berachain.dev/polaris/eth/common"
)

type BaseContract interface {
	coreprecompile.StatefulImpl
}

// `baseContract` is a base implementation of `StatefulImpl`.
type baseContract struct {
	// `contractAbi` stores the ABI of the precompile.
	contractAbi abi.ABI
	// `address stores the` address of the precompile.
	address common.Address
}

// `NewBaseContract` creates a new `BasePrecompile`.
func NewBaseContract(contractAbi abi.ABI, address common.Address) BaseContract {
	return &baseContract{
		contractAbi: contractAbi,
		address:     address,
	}
}

// `RegistryKey` implements StatefulImpl.
func (c *baseContract) RegistryKey() common.Address {
	return c.address
}

// `ABIMethods` implements StatefulImpl.
func (c *baseContract) ABIMethods() map[string]abi.Method {
	return c.contractAbi.Methods
}

// `ABIEvents` implements StatefulImpl.
func (c *baseContract) ABIEvents() map[string]abi.Event {
	return c.contractAbi.Events
}

// `CustomValueDecoders` implements StatefulImpl.
func (c *baseContract) CustomValueDecoders() coreprecompile.ValueDecoders {
	return nil
}

// `PrecompileMethods` implements StatefulImpl.
func (c *baseContract) PrecompileMethods() coreprecompile.Methods {
	return coreprecompile.Methods{}
}
