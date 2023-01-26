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
	"math/rand"

	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/x/evm/plugins/state/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("x/evm/plugins/state/types", func() {
	var state types.State
	key := common.Hash{}.Bytes()
	value := common.Hash{}.Bytes()

	BeforeEach(func() {
		rand.Read(key)
		rand.Read(value)
		state = types.NewState(common.BytesToHash(key), common.BytesToHash(value))
	})

	It("should return the correct key", func() {
		Expect(state.Key).To(Equal(common.BytesToHash(key).Hex()))
	})

	It("should return the correct value", func() {
		Expect(state.Value).To(Equal(common.BytesToHash(value).Hex()))
	})

	It("should have valid state", func() {
		Expect(state.ValidateBasic()).To(BeNil())
	})

	When("state key is empty", func() {
		BeforeEach(func() {
			state.Key = ""
		})

		It("should return an error", func() {
			Expect(state.ValidateBasic()).NotTo(BeNil())
		})
	})

	When("state key has leading or trailing spaces", func() {
		When("state key is not empty", func() {
			BeforeEach(func() {
				state.Key = " bingbong "
			})

			It("should not return an error", func() {
				Expect(state.ValidateBasic()).To(BeNil())
			})
		})

		When("state key is empty", func() {
			BeforeEach(func() {
				state.Key = "       "
			})

			It("should return an error", func() {
				Expect(state.ValidateBasic()).NotTo(BeNil())
			})
		})
	})

	It("is cloneable", func() {
		clone := state.Clone()
		Expect(clone).To(Equal(state))
		Expect(&clone).NotTo(BeIdenticalTo(&state))
	})
})
