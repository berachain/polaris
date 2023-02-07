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
	AccessList   = types.AccessList
	Block        = types.Block
	Bloom        = types.Bloom
	Log          = types.Log
	Receipt      = types.Receipt
	Receipts     = types.Receipts
	Transaction  = types.Transaction
	Transactions = types.Transactions
	Header       = types.Header
	BlockNonce   = types.BlockNonce
)

var (
	BytesToBloom   = types.BytesToBloom
	CreateBloom    = types.CreateBloom
	MakeSigner     = types.MakeSigner
	LogsBloom      = types.LogsBloom
	DeriveSha      = types.DeriveSha
	NewBlock       = types.NewBlock
	EmptyRootHash  = types.EmptyRootHash
	EmptyUncleHash = types.EmptyUncleHash
)
var (
	ReceiptStatusFailed     = types.ReceiptStatusFailed
	ReceiptStatusSuccessful = types.ReceiptStatusSuccessful
)
