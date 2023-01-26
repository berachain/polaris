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

package precompile_test

import (
	"context"
	"math/big"

	"github.com/berachain/stargazer/core/state"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/testutil"
	"github.com/berachain/stargazer/x/evm/precompile"

	sdk "github.com/cosmos/cosmos-sdk/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("cosmos runner", func() {
	cr := precompile.NewCosmosRunner()

	It("should use correctly consume gas", func() {
		_, remainingGas, err := cr.Run(&mockStateless{}, &mockSDB{&state.StateDB{}}, []byte{}, addr, new(big.Int), 30, false)
		Expect(err).To(BeNil())
		Expect(remainingGas).To(Equal(uint64(10)))
	})

	It("should error on insufficient gas", func() {
		_, _, err := cr.Run(&mockStateless{}, &mockSDB{&state.StateDB{}}, []byte{}, addr, new(big.Int), 5, true)
		Expect(err.Error()).To(Equal("out of gas"))
	})

	It("should plug in custom gas configs", func() {
		*cr = cr.WithKVGasConfig(&sdk.GasConfig{})
		*cr = cr.WithTransientKVGasConfig(&sdk.GasConfig{})
	})
})

// MOCKS BELOW.

type mockSDB struct {
	vm.StargazerStateDB
}

func (msdb *mockSDB) GetContext() context.Context {
	return testutil.NewContextWithMultistores()
}

type mockStateless struct{}

var addr = common.BytesToAddress([]byte{1})

func (ms *mockStateless) Address() common.Address {
	return addr
}

func (ms *mockStateless) Run(
	ctx context.Context, statedb vm.GethStateDB, input []byte,
	caller common.Address, value *big.Int, readonly bool,
) ([]byte, error) {
	sdk.UnwrapSDKContext(ctx).GasMeter().ConsumeGas(10, "")
	return nil, nil
}

func (ms *mockStateless) RequiredGas(input []byte) uint64 {
	return 10
}
