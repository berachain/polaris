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

package rpc

// `EthReaderBackend` is the backend for the `eth` namespace of the JSON-RPC API.
// It is only able to retrieve information about the current state of the chain by
// number. For querying data by hash, one must determine the block number that the
// data is stored at and then query by number.
// type EthReaderBackend struct {
// 	k keeper.Keeper
// }

// ==============================================================================
// EthReaderBackend
// ==============================================================================

// // `BlockNumber` implements the `eth_blockNumber` JSON-RPC method.
// func (eb *EthReaderBackend) BlockNumber(ctx sdk.Context) uint64 {
// 	return uint64(ctx.BlockHeight())
// }

// // `GetBlockByNumber` is used to implement the `eth_getBlockByNumber` JSON-RPCÍ.
// func (eb *EthReaderBackend) GetBlockByNumber(
// 	ctx sdk.Context, number uint64, fullTx bool,
// ) (*types.StargazerBlock, error) {
// 	block, found := eb.k.GetStargazerBlockAtHeight(ctx, number)
// 	if !found {
// 		return nil, errors.New("no block found")
// 	}
// 	return block, nil
// }

// // `GetStargazerBlockTransactionCountByNumber` returns the number of transactions in a block from a block
// // matching the given block number.
// func (eb *EthReaderBackend) BlockTransactionCountByNumber(ctx sdk.Context, number uint64) uint64 {
// 	// store := storeutils.KVStoreReaderAtBlockHeight(ctx, k.storeKey, int64(number))
// 	block, found := eb.k.GetStargazerBlockAtHeight(ctx, number)
// 	if !found {
// 		return 0
// 	}

// 	return uint64(block.TxIndex())
// }
