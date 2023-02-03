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
	"testing"

	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/x/evm/plugin/state/storage"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestStorage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "x/evm/plugin/state/storage")
}

var _ = Describe("StorageTest", func() {
	When("storage is empty", func() {
		It("should not return an error", func() {
			slots := storage.Storage{}
			Expect(slots.ValidateBasic()).To(BeNil())
		})
	})
	When("storage is not empty", func() {
		var slots storage.Storage

		BeforeEach(func() {
			slots = storage.Storage{
				storage.NewSlot(common.BytesToHash([]byte{1, 2, 3}), common.BytesToHash([]byte{1, 2, 3})),
			}
		})

		It("should not return an error", func() {
			Expect(slots.ValidateBasic()).To(BeNil())
		})

		When("a storage key is empty", func() {
			BeforeEach(func() {
				slots[0].Key = ""
			})

			It("should return an error", func() {
				Expect(slots.ValidateBasic()).NotTo(BeNil())
			})
		})

		It("should be Cloneable", func() {
			clone := slots.Clone()
			Expect(clone).To(Equal(slots))
			Expect(clone).NotTo(BeIdenticalTo(slots))
		})

		When("a storage key is duplicated", func() {
			BeforeEach(func() {
				slots = append(slots, storage.NewSlot(
					common.BytesToHash([]byte{1, 2, 3}),
					common.BytesToHash([]byte{1, 2, 3}),
				))
			})

			It("should return an error", func() {
				Expect(slots.ValidateBasic()).NotTo(BeNil())
			})
		})

		It("should be printable", func() {
			Expect(slots.String()).To(ContainSubstring("key:" +
				"\"0x0000000000000000000000000000000000000000000000000000000000010203\" value:" +
				"\"0x0000000000000000000000000000000000000000000000000000000000010203\"",
			))
		})
	})
})
