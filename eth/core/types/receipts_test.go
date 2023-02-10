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

package types_test

import (
	"github.com/berachain/stargazer/eth/core/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Receipts", func() {
	var sr *types.StargazerReceipts

	BeforeEach(func() {
		sr = types.NewStargazerReceipts()
	})

	It("should append and return the right length", func() {
		Expect(sr.Len()).To(Equal(uint(0)))

		sr.Append(&types.Receipt{Type: uint8(0)})
		Expect(sr.Len()).To(Equal(uint(1)))

		sr.Append(&types.Receipt{Type: uint8(1)})
		Expect(sr.Len()).To(Equal(uint(2)))

		Expect(sr.Receipts[0].Type).To(Equal(uint8(0)))
		Expect(sr.Receipts[1].Type).To(Equal(uint8(1)))

		Expect(func() { _ = sr.Receipts[2] }).To(Panic())
	})

	It("should create new stargazer receipts from receipts", func() {
		receipts := types.Receipts{{Type: uint8(0)}, {Type: uint8(1)}}
		stargazerReceipts := types.StargazerReceiptsFromReceipts(receipts)

		Expect(stargazerReceipts.Len()).To(Equal(uint(2)))
		Expect(stargazerReceipts.Receipts[0].Type).To(Equal(uint8(0)))
		Expect(stargazerReceipts.Receipts[1].Type).To(Equal(uint8(1)))
	})
})
