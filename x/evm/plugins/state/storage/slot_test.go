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

package storage_test

import (
	"math/rand"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/x/evm/plugins/state/storage"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("x/evm/plugins/state/storage", func() {
	var slot *storage.Slot
	key := common.Hash{}.Bytes()
	value := common.Hash{}.Bytes()

	BeforeEach(func() {
		rand.Read(key)
		rand.Read(value)
		slot = storage.NewSlot(common.BytesToHash(key), common.BytesToHash(value))
	})

	It("should return the correct key", func() {
		Expect(slot.Key).To(Equal(common.BytesToHash(key).Hex()))
	})

	It("should return the correct value", func() {
		Expect(slot.Value).To(Equal(common.BytesToHash(value).Hex()))
	})

	It("should have valid slot", func() {
		Expect(slot.ValidateBasic()).To(BeNil())
	})

	When("slot key is empty", func() {
		BeforeEach(func() {
			slot.Key = ""
		})

		It("should return an error", func() {
			Expect(slot.ValidateBasic()).NotTo(BeNil())
		})
	})

	When("slot key has leading or trailing spaces", func() {
		When("slot key is not empty", func() {
			BeforeEach(func() {
				slot.Key = " bingbong "
			})

			It("should not return an error", func() {
				Expect(slot.ValidateBasic()).To(BeNil())
			})
		})

		When("slot key is empty", func() {
			BeforeEach(func() {
				slot.Key = "       "
			})

			It("should return an error", func() {
				Expect(slot.ValidateBasic()).NotTo(BeNil())
			})
		})
	})

	It("is cloneable", func() {
		clone := slot.Clone()
		Expect(clone).To(Equal(slot))
		Expect(&clone).NotTo(BeIdenticalTo(&slot))
	})
})
