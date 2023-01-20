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

	"github.com/berachain/stargazer/core/vm/precompile"
	"github.com/berachain/stargazer/lib/common"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Container Factories", func() {
	Context("Stateless Container Factory", func() {
		var scf *precompile.StatelessContainerFactory

		BeforeEach(func() {
			scf = precompile.NewStatelessContainerFactory()
		})

		It("should build stateless precompile containers", func() {
			pc, err := scf.Build(&mockStateless{})
			Expect(err).To(BeNil())
			Expect(pc).ToNot(BeNil())
		})
	})
})

// MOCKS BELOW.

type mockStateless struct{}

func (ms *mockStateless) Address() common.Address {
	return common.Address{}
}

func (ms *mockStateless) RequiredGas(input []byte) uint64 {
	return 0
}

func (ms *mockStateless) Run(
	ctx context.Context, input []byte, caller common.Address,
	value *big.Int, readonly bool,
) ([]byte, error) {
	return nil, nil
}
