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
package types

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	EthLog       = types.Log
	Bloom        = types.Bloom
	AccessList   = types.AccessList
	AccessListTx = types.AccessListTx
	AccessTuple  = types.AccessTuple
	LegacyTx     = types.LegacyTx
	DynamicFeeTx = types.DynamicFeeTx
	TxData       = types.TxData
	Signer       = types.Signer
	Transaction  = types.Transaction
	Receipt      = types.Receipt
	Log          = types.Log
)

var (
	LatestSignerForChainID  = types.LatestSignerForChainID
	ReceiptStatusSuccessful = types.ReceiptStatusSuccessful
	BytesToBloom            = types.BytesToBloom
	MakeSigner              = types.MakeSigner
	LogsBloom               = types.LogsBloom
	NewTx                   = types.NewTx
	NewMessage              = types.NewMessage
)

const (
	LegacyTxType     = types.LegacyTxType
	AccessListTxType = types.AccessListTxType
	DynamicFeeTxType = types.DynamicFeeTxType
)
