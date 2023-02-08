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

package mock

//go:generate moq -out ./refund.mock.go -pkg mock ../../ RefundJournal

// `NewEmptyRefundJournal` returns an empty `RefundJournalMock`.
func NewEmptyRefundJournal() *RefundJournalMock {
	return &RefundJournalMock{
		AddRefundFunc: func(gas uint64) {
			panic("mock out the AddRefund method")
		},
		FinalizeFunc: func() {
			// no-op
		},
		GetRefundFunc: func() uint64 {
			panic("mock out the GetRefund method")
		},
		RegistryKeyFunc: func() string {
			return "emptyrefund"
		},
		RevertToSnapshotFunc: func(n int) {
			// no-op
		},
		SnapshotFunc: func() int {
			// no-op
			return 0
		},
		SubRefundFunc: func(gas uint64) {
			panic("mock out the SubRefund method")
		},
	}
}
