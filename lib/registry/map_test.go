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

package registry_test

import (
	"testing"

	"github.com/berachain/stargazer/lib/registry"
	"github.com/berachain/stargazer/lib/registry/mock"
	libtypes "github.com/berachain/stargazer/lib/types"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRegister(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "lib/registry")
}

var _ = Describe("Registry", func() {
	var r libtypes.Registry[string, libtypes.Registrable[string]]

	BeforeEach(func() {
		r = registry.NewMap[string, libtypes.Registrable[string]]()
	})

	When("adding an item", func() {
		BeforeEach(func() {
			// Register an item.
			item := mock.NewMockRegistrable("foo", "bar")
			err := r.Register(item)
			Expect(err).ToNot(HaveOccurred())
		})

		It("should return an error if the item already exists", func() {
			// Register the same item again.
			err := r.Register(mock.NewMockRegistrable("foo", "bar"))
			Expect(err).To(HaveOccurred())
		})

		It("should be able to get the item", func() {
			// Get the item.
			item := r.Get("foo")
			Expect(item.RegistryKey()).To(Equal("foo"))
		})

		It("should be able to remove the item", func() {
			// Remove the item.
			r.Remove("foo")

			// Get the item.
			item := r.Get("foo")
			Expect(item).To(BeNil())
		})

		It("should be able to check if the item exists", func() {
			// Check if the item exists.
			exists := r.Exists("foo")
			Expect(exists).To(BeTrue())

			// Remove the item.
			r.Remove("foo")

			// Check if the item exists.
			exists = r.Exists("foo")
			Expect(exists).To(BeFalse())
		})

		It("should be able to check if an item does not exist", func() {
			// Check an item that does not exist.
			exists := r.Exists("bar")
			Expect(exists).To(BeFalse())
		})

		It("should no-op when removing an item that does not exist", func() {
			// Remove an item that does not exist.
			r.Remove("bar")
		})
	})
})
