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

//go:generate rlpgen -type StargazerReceipts -out receipts.rlpgen.go -decoder

const initLen = 64

// `StargazerReceipts` is a slice of *ReceiptForStorage.
type StargazerReceipts struct {
	Receipts []*ReceiptForStorage
}

// `NewStargazerReceipts` creates a new list of receipts.
func NewStargazerReceipts() *StargazerReceipts {
	return &StargazerReceipts{
		Receipts: make([]*ReceiptForStorage, initLen),
	}
}

// `Append` appends a receipt to the list of receipts.
func (sr *StargazerReceipts) Append(r *ReceiptForStorage) {
	sr.Receipts = append(sr.Receipts, r)
}

// `Len` returns the number of receipts in the list.
func (sr *StargazerReceipts) Len() uint {
	return uint(len(sr.Receipts))
}
