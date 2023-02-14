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

package state_test

import (
	"math/big"

	"github.com/berachain/stargazer/eth/common"
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/state/journal/mock"
	"github.com/berachain/stargazer/eth/core/vm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	alice = common.Address{1}
	bob   = common.Address{2}
)

var _ = Describe("StateDB", func() {
	var sdb vm.StargazerStateDB

	BeforeEach(func() {
		sdb = state.NewStateDB(mock.NewEmptyStatePlugin())
	})

	It("Should suicide correctly", func() {
		sdb.CreateAccount(alice)
		Expect(sdb.Suicide(alice)).To(BeFalse())
		Expect(sdb.HasSuicided(alice)).To(BeFalse())

		sdb.CreateAccount(bob)
		sdb.SetCode(bob, []byte{1, 2, 3})
		sdb.AddBalance(bob, big.NewInt(10))
		Expect(sdb.Suicide(bob)).To(BeTrue())
		Expect(sdb.GetBalance(bob).Uint64()).To(Equal(uint64(0)))
		Expect(sdb.HasSuicided(bob)).To(BeTrue())
	})

	It("should handle empty", func() {
		sdb.CreateAccount(alice)
		Expect(sdb.Empty(alice)).To(BeTrue())

		sdb.SetCode(alice, []byte{1, 2, 3})
		Expect(sdb.Empty(alice)).To(BeFalse())
	})

	It("should snapshot/revert", func() {
		Expect(func() {
			id := sdb.Snapshot()
			sdb.RevertToSnapshot(id)
		}).ToNot(Panic())
	})

	It("should delete suicides on finalize", func() {
		sdb.CreateAccount(bob)
		sdb.SetCode(bob, []byte{1, 2, 3})
		sdb.AddBalance(bob, big.NewInt(10))
		Expect(sdb.Suicide(bob)).To(BeTrue())
		Expect(sdb.GetBalance(bob).Uint64()).To(Equal(uint64(0)))
		Expect(sdb.HasSuicided(bob)).To(BeTrue())

		sdb.Finalize()
		Expect(sdb.HasSuicided(bob)).To(BeFalse())
	})
})
