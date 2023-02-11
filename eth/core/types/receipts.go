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
	"unsafe"

	"github.com/ethereum/go-ethereum/core/types"
)

// `initCapacity` is the initial capacity of the list of receipts.
const initCapacity = 64

// `StargazerReceipts` is a slice of `*ReceiptForStorage` receipts.
//
//go:generate rlpgen -type StargazerReceipts -out receipts.rlpgen.go -decoder
type StargazerReceipts struct {
	// `Receipts` is a list of `ReceiptForStorage`s, each of which represent a transaction receipt.
	Receipts []*ReceiptForStorage
}

// `NewStargazerReceipts` creates and returns a `StargazerReceipts` with a list of receipts.
func NewStargazerReceipts() *StargazerReceipts {
	return &StargazerReceipts{
		Receipts: make([]*ReceiptForStorage, 0, initCapacity),
	}
}

// `StargazerReceiptsFromReceipts` converts a list of `Receipt`s to a `StargazerReceipts`.
func StargazerReceiptsFromReceipts(receipts Receipts) *StargazerReceipts {
	//#nosec:G103 unsafe pointer is safe here since `ReceiptForStorage` is an alias of `Receipt`.
	return &StargazerReceipts{
		Receipts: *(*([]*ReceiptForStorage))((unsafe.Pointer(&receipts))),
	}
}

// `Append` appends a receipt to the list of receipts.
func (sr *StargazerReceipts) Append(r *Receipt) {
	//#nosec:G103
	sr.Receipts = append(sr.Receipts, ((*ReceiptForStorage)(unsafe.Pointer(r))))
}

// `Bloom` returns the bloom filter of the list of receipts.
func (sr *StargazerReceipts) Bloom() Bloom {
	//#nosec:G103 unsafe pointer is safe here since `ReceiptForStorage` is an alias of `Receipt`.
	return types.CreateBloom(*(*(Receipts))((unsafe.Pointer(&sr.Receipts))))
}

// `Len` returns the number of receipts in the list.
func (sr *StargazerReceipts) Len() uint {
	return uint(len(sr.Receipts))
}
