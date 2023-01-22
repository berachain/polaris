// Copyright (C) 2023, Berachain Foundation. All rights reserved.
// See the file LICENSE for licensing terms.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package core

import (
	"math"
	"math/big"

	"cosmossdk.io/errors"
	"github.com/berachain/stargazer/core/state"
	"github.com/berachain/stargazer/core/types"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

// `EVMFactory` is used to build new Stargazer `EVM`s.
type EVMFactory struct {
	// `pr` is the precompile registry is responsible for keeping track of the stateful
	// precompiles, events, and errors that are available to the EVM.
	pr *vm.PrecompileRegistry

	// `sk` represents the entity responsible for providing historical block information to the
	// EVM. It is used in the GetHashFn which is ultimately used in the `BLOCKHASH` opcode.
	sk types.StakingKeeper
}

// `NewEVMFactory` creates and returns a new `EVMFactory` with a new `vm.PrecompileRegistry`.
func NewEVMFactory() *EVMFactory {
	return &EVMFactory{
		pr: vm.NewPrecompileRegistry(),
	}
}

// `Build` creates and returns a new `vm.StargazerEVM` with a new `vm.PrecompileHost`.
func (ef *EVMFactory) Build(
	ssdb state.StargazerStateDB,
	blockCtx vm.BlockContext,
	txCtx vm.TxContext,
	chainConfig *params.EthChainConfig,
	noBaseFee bool,
) *vm.StargazerEVM {
	return vm.NewStargazerEVM(
		blockCtx, txCtx, ssdb, chainConfig,
		ef.BuildVMConfig(nil, noBaseFee),
		vm.NewPrecompileHost(
			ef.pr,
			ssdb,
		),
	)
}

// BuildVMConfig returns a new VMConfig to be used in the VM.
func (ef *EVMFactory) BuildVMConfig(tracer vm.EVMLogger, noBaseFee bool) vm.Config {
	// extraEIPs := ef.extraEIPs.Get(ctx) // silence linter

	// // TODO: this is so bad need to get rid of this garbage proto type
	// if extraEIPs == nil {
	// 	extraEIPs = &params.ExtraEIPs{}
	// }

	// TODO: this is so bad like holy fuck on god fr fr
	// eips := make([]int, len(extraEIPs.EIPs))

	eips := make([]int, 0)
	// for i, eip := range extraEIPs.EIPs {
	// 	eips[i] = int(eip)
	// }

	return vm.Config{
		Debug:     tracer != nil,
		Tracer:    tracer,
		NoBaseFee: noBaseFee,
		ExtraEips: eips,
	}
}

func (ef *EVMFactory) BuildBlockContext(
	ctx sdk.Context,
	gasLimit uint64,
	basefee *big.Int,
) vm.BlockContext {
	coinbase, err := ef.getCoinbaseFromContext(ctx)
	if err != nil {
		ctx.Logger().Error("failed to retrieve coinbase from context", "error", err)
	}

	return vm.BlockContext{
		CanTransfer: CanTransfer,
		Transfer:    Transfer,
		GetHash:     ef.getHashFnFromCosmosContext(ctx),
		Coinbase:    coinbase,
		BlockNumber: big.NewInt(ctx.BlockHeader().Height),
		GasLimit:    gasLimit,
		Time:        big.NewInt(ctx.BlockHeader().Time.Unix()),
		Difficulty:  new(big.Int), // makes no sense outside of PoW.
		BaseFee:     basefee,
	}
}

func (ef *EVMFactory) getHashFnFromCosmosContext(
	ctx sdk.Context,
) func(uint64) common.Hash {
	return func(h uint64) common.Hash {
		if blockHeight = h 
			return // Case 1: The requested height matches the one from the context so we can
			// retrieve the header hash directly from the context.
			// Note: The headerHash is only set at begin block, it will be nil in case of a
			// query context
			headerHash := ctx.HeaderHash()
			if len(headerHash) != 0 {
				return common.BytesToHash(headerHash)
			}

			// only recompute the hash if not set (eg: checkTxState)
			contextBlockHeader := ctx.BlockHeader()
			header, err := tmtypes.HeaderFromProto(&contextBlockHeader)
			if err != nil {
				ctx.Logger().Error("failed to cast tendermint header from proto", "error", err)
				return common.Hash{}
			}

			return common.BytesToHash(header.Hash())

		else:
		return thing.Get(h)
	}
		blockHeight := uint64(ctx.BlockHeight())
		switch {
		case blockHeight == h:
			

		case blockHeight - h:
			// Case 2: if the chain is not the current height we need to retrieve the hash from
			// the store for the current chain epoch. This only applies if the current height is
			// greater than the requested height.

			// If the requested height is greater than the max uint64 value,
			// we return an empty hash.
			if h > uint64(math.MaxInt64) {
				return common.Hash{}
			}

			histInfo, found := ef.sk.GetHistoricalInfo(ctx, int64(h))
			if !found {
				ctx.Logger().Debug("historical info not found", "height", h)
				return common.Hash{}
			}

			header, err := tmtypes.HeaderFromProto(&histInfo.Header)
			if err != nil {
				ctx.Logger().Error("failed to cast tendermint header from proto", "error", err)
				return common.Hash{}
			}

			return common.BytesToHash(header.Hash())
		default:
			// Case 3: heights greater than the current one returns an empty hash.
			return common.Hash{}
		}
	}
}

// getCoinbaseFromContext returns the block proposer's validator operator address.
func (ef *EVMFactory) getCoinbaseFromContext(
	ctx sdk.Context,
) (common.Address, error) {
	// todo: add redundancy here, incrase BlockHeader().ProposerAddress is not found, we want
	// to make sure that life is gucci as sometimes it doesn't matter.
	validator, found := ef.sk.GetValidatorByConsAddr(ctx, ctx.BlockHeader().ProposerAddress)
	if !found {
		return common.Address{}, errors.Wrapf(
			stakingtypes.ErrNoValidatorFound,
			"failed to retrieve validator operator from block proposer address %s",
			ctx.BlockHeader().ProposerAddress,
		)
	}

	return common.BytesToAddress(validator.GetOperator()), nil
}
