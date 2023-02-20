//nolint:lll
//go:generate abigen --pkg generated staking ../contracts/solidity/out/staking.sol/IStakingModule.json --out ./generated/staking.abigen.go --type StakingModule

package staking

import (
	"context"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/precompile"
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/precompile/contracts/solidity/generated"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var (
	cosmosEventTypes = []string{
		stakingtypes.EventTypeDelegate,
		stakingtypes.EventTypeRedelegate,
		stakingtypes.EventTypeCreateValidator,
		stakingtypes.EventTypeUnbond,
		stakingtypes.EventTypeCancelUnbondingDelegation,
	}

	baseGas = uint64(500)
)

var (
	_ precompile.StatefulPrecompileImpl = (*Contract)(nil)
)

type Contract struct {
	vm.PrecompileContainer

	msgServer stakingtypes.MsgServer
	querier   stakingtypes.QueryServer

	contractAbi abi.ABI
}

// `NewContract` is the constructor of the staking contract.
func NewContract(sk *stakingkeeper.Keeper) *Contract {
	var contractAbi abi.ABI
	if err := contractAbi.UnmarshalJSON([]byte(generated.StakingModuleMetaData.ABI)); err != nil {
		panic(err)
	}
	return &Contract{
		msgServer:   stakingkeeper.NewMsgServerImpl(sk),
		querier:     stakingkeeper.Querier{Keeper: sk},
		contractAbi: contractAbi,
	}
}

// `ABIMethods` implements StatefulPrecompileImpl.
func (c *Contract) ABIMethods() map[string]abi.Method {
	return c.contractAbi.Methods
}

// `PrecompileMethods` implements StatefulPrecompileImpl.
func (c *Contract) PrecompileMethods() precompile.Methods {
	return precompile.Methods{
		&precompile.Method{
			AbiSig:  "getDelegation(address)",
			Execute: c.GetDelegationAddrInput,
		},
		&precompile.Method{
			AbiSig:  "getDelegation(string)",
			Execute: c.GetDelegationStringInput,
		},
		&precompile.Method{
			AbiSig:  "getUnbondingDelegation(address)",
			Execute: c.GetUnbondingDelegationAddrInput,
		},
		&precompile.Method{
			AbiSig:  "getUnbondingDelegation(string)",
			Execute: c.GetUnbondingDelegationStringInput,
		},
		&precompile.Method{
			AbiSig:  "getRedelegations(address,address)",
			Execute: c.GetRedelegationsAddrInput,
		},
		&precompile.Method{
			AbiSig:  "getRedelegations(string,string)",
			Execute: c.GetRedelegationsStringInput,
		},
		&precompile.Method{
			AbiSig:  "delegate(address,uint256)",
			Execute: c.DelegateAddrInput,
		},
		&precompile.Method{
			AbiSig:  "delegate(string,uint256)",
			Execute: c.DelegateStringInput,
		},
		&precompile.Method{
			AbiSig:  "undelegate(address,uint256)",
			Execute: c.UndelegateAddrInput,
		},
		&precompile.Method{
			AbiSig:  "undelegate(string,uint256)",
			Execute: c.UndelegateStringInput,
		},
		&precompile.Method{
			AbiSig:  "beginRedelegate(address,address,uint256)",
			Execute: c.BeginRedelegateAddrInput,
		},
		&precompile.Method{
			AbiSig:  "beginRedelegate(string,string,uint256)",
			Execute: c.BeginRedelegateStringInput,
		},
		&precompile.Method{
			AbiSig:  "cancelUnbondingDelegation(address,uint256,int64)",
			Execute: c.CancelUnbondingDelegationAddrInput,
		},
		&precompile.Method{
			AbiSig:  "cancelUnbondingDelegation(string,uint256,int64)",
			Execute: c.CancelUnbondingDelegationStringInput,
		},
	}
}

// `ABIEvents` implements StatefulPrecompileImpl.
func (c *Contract) ABIEvents() map[string]abi.Event {
	return c.contractAbi.Events
}

// `CustomValueDecoders` implements StatefulPrecompileImpl.
func (c *Contract) CustomValueDecoders() precompile.ValueDecoders {
	return nil
}

// `GetDelegationAddrInput` implements `getDelegation(address)` method.
func (c *Contract) GetDelegationAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}

	return c.delegationHelper(ctx, caller, sdk.ValAddress(val.Bytes()))
}

// `GetDelegationStringInput` implements `getDelegation(string)` method.
func (c *Contract) GetDelegationStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return c.delegationHelper(ctx, caller, val)
}

// `GetUnbondingDelegationAddrInput` implements the `getUnbondingDelegation(address)` method.
func (c *Contract) GetUnbondingDelegationAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}

	return c.getUnbondingDelegationHelper(ctx, caller, sdk.ValAddress(val.Bytes()))
}

func (c *Contract) GetUnbondingDelegationStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return c.getUnbondingDelegationHelper(ctx, caller, val)
}

func (c *Contract) GetRedelegationsAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	dstVal, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}

	return c.GetRedelegationsHelper(ctx, caller, sdk.ValAddress(srcVal.Bytes()), sdk.ValAddress(dstVal.Bytes()))
}

func (c *Contract) GetRedelegationsStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	dstVal, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, ErrInvalidString
	}

	src, err := sdk.ValAddressFromBech32(srcVal)
	if err != nil {
		return nil, err
	}
	dst, err := sdk.ValAddressFromBech32(dstVal)
	if err != nil {
		return nil, err
	}

	return c.GetRedelegationsHelper(ctx, caller, src, dst)
}

func (c *Contract) DelegateAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	return nil, c.delegateHelper(ctx, caller, amount, sdk.ValAddress(val.Bytes()))
}

func (c *Contract) DelegateStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return nil, c.delegateHelper(ctx, caller, amount, val)
}

func (c *Contract) UndelegateAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	return nil, c.undelegateHelper(ctx, caller, amount, sdk.ValAddress(val.Bytes()))
}

func (c *Contract) UndelegateStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return nil, c.undelegateHelper(ctx, caller, amount, val)
}

func (c *Contract) BeginRedelegateAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	dstVal, ok := utils.GetAs[common.Address](args[1])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	return nil, c.beginRedelegateHelper(ctx, caller, amount, sdk.ValAddress(srcVal.Bytes()), sdk.ValAddress(dstVal.Bytes()))
}

func (c *Contract) BeginRedelegateStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	srcVal, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	dstVal, ok := utils.GetAs[string](args[1])
	if !ok {
		return nil, ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[2])
	if !ok {
		return nil, ErrInvalidBigInt
	}

	src, err := sdk.ValAddressFromBech32(srcVal)
	if err != nil {
		return nil, err
	}
	dst, err := sdk.ValAddressFromBech32(dstVal)
	if err != nil {
		return nil, err
	}

	return nil, c.beginRedelegateHelper(ctx, caller, amount, src, dst)
}

func (c *Contract) CancelUnbondingDelegationAddrInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	val, ok := utils.GetAs[common.Address](args[0])
	if !ok {
		return nil, ErrInvalidValidatorAddr
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}
	creationHeight, ok := utils.GetAs[int64](args[2])
	if !ok {
		return nil, ErrInvalidInt64
	}

	return nil, c.cancelUnbondingDelegationHelper(ctx, caller, amount, sdk.ValAddress(val.Bytes()), creationHeight)
}

func (c *Contract) CancelUnbondingDelegationStringInput(
	ctx context.Context,
	caller common.Address,
	value *big.Int,
	readonly bool,
	args ...any,
) ([]any, error) {
	bech32Addr, ok := utils.GetAs[string](args[0])
	if !ok {
		return nil, ErrInvalidString
	}
	amount, ok := utils.GetAs[*big.Int](args[1])
	if !ok {
		return nil, ErrInvalidBigInt
	}
	creationHeight, ok := utils.GetAs[int64](args[2])
	if !ok {
		return nil, ErrInvalidInt64
	}

	val, err := sdk.ValAddressFromBech32(bech32Addr)
	if err != nil {
		return nil, err
	}

	return nil, c.cancelUnbondingDelegationHelper(ctx, caller, amount, val, creationHeight)
}
