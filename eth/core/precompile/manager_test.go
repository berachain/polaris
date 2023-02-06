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

package precompile

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/lib/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("controller", func() {
	var c *manager
	var mr *mockRunner
	var ctx context.Context
	var ms *mockSdb

	BeforeEach(func() {
		mr = &mockRunner{}
		ms = &mockSdb{}
		ctx = context.Background()
		c = utils.MustGetAs[*manager](NewManager(mr, ms))
		err := c.Register(&mockStateless{})
		Expect(err).To(BeNil())
	})

	It("should find and run", func() {
		err := c.Reset(ctx)
		Expect(err).To(BeNil())

		pc := c.Get(addr)
		Expect(pc).ToNot(BeNil())

		_, _, err = c.Run(pc, []byte{}, addr, new(big.Int), 10, true)
		Expect(err).To(BeNil())
		Expect(mr.called).To(BeTrue())
	})

	It("should not find an unregistered", func() {
		found := c.Has(common.BytesToAddress([]byte{2}))
		Expect(found).To(BeFalse())
	})

	It("should error on incompatible context", func() {
		err := c.Reset(nil) //nolint:staticcheck // only for tests.
		Expect(err.Error()).To(Equal("the context passed in is nil"))
	})
})

// MOCKS BELOW.

type mockRunner struct {
	called bool
}

func (mr *mockRunner) Run(
	ctx context.Context, ldb vm.LogsDB, pc vm.PrecompileContainer, input []byte,
	caller common.Address, value *big.Int, suppliedGas uint64, readonly bool,
) ([]byte, uint64, error) {
	mr.called = true
	return nil, 0, nil
}

type mockSdb struct {
	vm.StargazerStateDB
}
