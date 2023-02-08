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

package core_test

import (
	. "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega.
)

var _ = Describe("EVM Test Suite", func() {
	// var host *mock.StargazerHostChainMock

	// hash1 := common.Hash{1}
	// hash2 := common.Hash{2}
	// hash3 := common.Hash{3}
	// hash4 := common.Hash{4}

	// currentHeader := &types.StargazerHeader{
	// 	Header: &types.Header{
	// 		Number:     big.NewInt(int64(123)),
	// 		BaseFee:    big.NewInt(69),
	// 		ParentHash: common.Hash{111},
	// 	},
	// 	CachedHash: common.Hash{1},
	// }

	Context("TestGetHashFunc", func() {
		BeforeEach(func() {
			// host = mock.NewMockHost()
		})
		// It("should return the correct hash", func() {
		// 	host.StargazerHeaderAtHeightFunc = func(ctx context.Context, height uint64) *types.StargazerHeader {
		// 		return &types.StargazerHeader{
		// 			Header: &types.Header{
		// 				Number:     big.NewInt(int64(height)),
		// 				BaseFee:    big.NewInt(69),
		// 				ParentHash: common.Hash{123},
		// 			},
		// 			CachedHash: crypto.Keccak256Hash([]byte{byte(height)}),
		// 		}
		// 	}
		// 	hash := core.GetHashFn(context.Background(), currentHeader, host)
		// 	Expect(hash(112)).To(Equal(crypto.Keccak256Hash([]byte{byte(112)})))
		// })
	})
})
