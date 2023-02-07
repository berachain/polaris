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

package storage

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	blockKey   = []byte("block")
	receiptKey = []byte("receipt")
	hashKey    = []byte("hash")
)

// `BlockKey` returns the key for the block bloom.
func BlockKey() []byte {
	return blockKey
}

// `HashToTxIndexKey` returns the key for a receipt lookup.
func HashToTxIndexKey(h []byte) []byte {
	return append(hashKey, h...)
}

// `TxIndexToRecieptKey` returns the key for the receipt lookup for a given block.
func TxIndexToRecieptKey(txIndex uint64) []byte {
	return append(receiptKey, sdk.Uint64ToBigEndian(txIndex)...)
}
