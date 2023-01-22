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

package log_test

import (
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Log Registry", func() {
	var lr *log.Registry
	var addr = common.BytesToAddress([]byte("my precompile address"))
	var abiEvent abi.Event

	BeforeEach(func() {
		lr = log.NewRegistry()
		abiEvent = abi.Event{Name: "CancelUnbondingDelegation"}
	})

	Describe("Test Register Event", func() {
		It("should handle registration properly", func() {
			err := lr.RegisterEvent(addr, abiEvent, nil)
			Expect(err).To(BeNil())

			err = lr.RegisterEvent(addr, abiEvent, nil)
			Expect(err.Error()).To(Equal("this Ethereum event is already registered: CancelUnbondingDelegation"))
		})
	})

	Describe("Test Get Precompile Log", func() {
		It("should correctly return existing/non-existing logs", func() {
			// event not registeredß
			log := lr.GetPrecompileLog("cancel_unbonding_delegation")
			Expect(log).To(BeNil())

			// valid event registered
			err := lr.RegisterEvent(addr, abiEvent, nil)
			Expect(err).To(BeNil())
			log = lr.GetPrecompileLog("cancel_unbonding_delegation")
			Expect(log).ToNot(BeNil())

			// event that doesn't exist
			log = lr.GetPrecompileLog("cancel-unbonding-delegation")
			Expect(log).To(BeNil())
		})
	})
})
