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

package utils

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// `KVStoreReader` is a subset of the `KVStore` interface that only exposes read
// methods.
type KVStoreReader interface {
	// Get returns nil if key doesn't exist. Panics on nil key.
	Get(key []byte) []byte

	// Has checks if a key exists. Panics on nil key.
	Has(key []byte) bool
}

// `KVStoreReaderAtBlockHeight` returns a KVStoreReader at a given height. If the height is greater
// than or equal to the current height, the reader will be at the latest height. We return the store
// with the modified height as a `KVStoreReader` since it does not make any sense to return a `KVStore`
// since we cannot update historical versions of the tree.
func KVStoreReaderAtBlockHeight(ctx sdk.Context, storeKey storetypes.StoreKey, height int64) KVStoreReader {
	if height >= ctx.BlockHeight() {
		return ctx.KVStore(storeKey)
	}

	// `version` is 1-indexed, so we need to increment the height by 1.
	cms, err := ctx.MultiStore().CacheMultiStoreWithVersion(height + 1)
	if err != nil {
		panic(err)
	}
	return ctx.WithMultiStore(cms).KVStore(storeKey)
}
