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
	"github.com/berachain/stargazer/lib/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Header", func() {
	var h *types.Header
	var ha common.Hash
	var sh *types.StargazerHeader

	BeforeEach(func() {
		ha = common.BytesToHash([]byte{1})
		h = &types.Header{
			Coinbase: common.BytesToAddress([]byte{2}),
			Bloom:    types.BytesToBloom([]byte{3}),
		}
		sh = types.NewStargazerHeader(h, ha)
	})

	It("should return the correct values", func() {
		Expect(sh.Author()).To(Equal(common.BytesToAddress([]byte{2})))
		Expect(sh.Hash()).To(Equal(common.BytesToHash([]byte{1})))

		sh.SetHash(common.Hash{})
		Expect(sh.Hash()).To(Equal(h.Hash()))

		data, err := sh.MarshalBinary()
		Expect(err).To(BeNil())
		sh2 := types.NewEmptyStargazerHeader()
		err = sh2.UnmarshalBinary(data)
		Expect(err).To(BeNil())

		Expect(sh2.Author()).To(Equal(sh.Author()))
		Expect(sh2.Hash()).To(Equal(sh.Hash()))
		Expect(sh2.Header.Bloom).To(Equal(sh.Header.Bloom))
		Expect(sh2.Header.ReceiptHash).To(Equal(sh.Header.ReceiptHash))
	})
})
