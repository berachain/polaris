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

package keeper

import (
	"github.com/berachain/stargazer/eth/core/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/x/evm/key"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetReceipt stores the receipt indexed by the tx index.
func (k *Keeper) SetReceipt(ctx sdk.Context, receipt *types.Receipt) {
	bz, err := receipt.MarshalBinary()
	if err != nil {
		panic(err)
	}

	// We need to store the receipt for the block numer + tx index for efficient iteration., but
	// we also need to allow for a way to lookup a receipt by hash.
	receiptKey := key.TxIndexToReciept(
		uint64(receipt.TransactionIndex),
	)
	// Store the receiptKey in the store with a key of the tx hash.
	ctx.KVStore(k.storeKey).Set(key.HashToTxIndex(receipt.TxHash.Bytes()), receiptKey)

	// Store the receipt indexed by tx index.
	ctx.KVStore(k.storeKey).Set(receiptKey, bz)
}

// `GetReceipt` gets the receipt indexed by the receipt hash.
func (k *Keeper) GetReceipt(ctx sdk.Context, txIndex uint64) *types.Receipt {
	receiptKey := key.TxIndexToReciept(txIndex)
	bz := ctx.KVStore(k.storeKey).Get(receiptKey)
	if bz == nil {
		return nil
	}
	receipt := new(types.Receipt)
	if err := receipt.UnmarshalBinary(bz); err != nil {
		panic(err)
	}
	return receipt
}

// `GetReceiptByTxHash` gets the receipt indexed by the transaction hash.
func (k *Keeper) GetReceiptByTxHash(ctx sdk.Context, txHash common.Hash) *types.Receipt {
	receiptKey := ctx.KVStore(k.storeKey).Get(key.HashToTxIndex(txHash.Bytes()))
	if receiptKey == nil {
		return nil
	}
	bz := ctx.KVStore(k.storeKey).Get(receiptKey)
	if bz == nil {
		return nil
	}
	receipt := new(types.Receipt)
	if err := receipt.UnmarshalBinary(bz); err != nil {
		panic(err)
	}
	return receipt
}
