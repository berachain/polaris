package mock

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	gethvm "github.com/ethereum/go-ethereum/core/vm"
	"github.com/holiman/uint256"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/vm"
)

type PrecompileEVM interface {
	GetStateDB() gethvm.StateDB

	Call(
		caller vm.ContractRef,
		addr common.Address,
		input []byte,
		gas uint64,
		value *big.Int,
	) (ret []byte, leftOverGas uint64, err error)
	StaticCall(
		caller vm.ContractRef,
		addr common.Address,
		input []byte,
		gas uint64,
	) (ret []byte, leftOverGas uint64, err error)
	Create(
		caller vm.ContractRef,
		code []byte,
		gas uint64,
		value *big.Int,
	) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error)
	Create2(
		caller vm.ContractRef,
		code []byte,
		gas uint64,
		endowment *big.Int,
		salt *uint256.Int,
	) (ret []byte, contractAddr common.Address, leftOverGas uint64, err error)
	GetContext() *vm.BlockContext
}

type MessageRouter interface {
	Handler(msg sdk.Msg) baseapp.MsgServiceHandler
	HandlerByTypeURL(typeURL string) baseapp.MsgServiceHandler
}
