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

package mock

import (
	"math/big"
	"time"

	"github.com/berachain/stargazer/eth/common"
	ethereumcorevm "github.com/ethereum/go-ethereum/core/vm"
)

//go:generate moq -out ./logger.mock.go -pkg mock ../ EVMLogger

func NewEVMLoggerMock() *EVMLoggerMock {
	mockedEVMLogger := &EVMLoggerMock{
		CaptureEndFunc: func(output []byte, gasUsed uint64, t time.Duration, err error) {
			// no-op
		},
		CaptureEnterFunc: func(typ ethereumcorevm.OpCode,
			from common.Address, to common.Address, input []byte, gas uint64, value *big.Int) {
			// no-op
		},
		CaptureExitFunc: func(output []byte, gasUsed uint64, err error) {
			// no-op
		},
		CaptureFaultFunc: func(pc uint64,
			op ethereumcorevm.OpCode, gas uint64, cost uint64,
			scope *ethereumcorevm.ScopeContext, depth int, err error) {
			// no-op
		},
		CaptureStartFunc: func(env *ethereumcorevm.EVM,
			from common.Address, to common.Address, create bool, input []byte, gas uint64,
			value *big.Int) {
			// no-op
		},
		CaptureStateFunc: func(pc uint64,
			op ethereumcorevm.OpCode, gas uint64, cost uint64,
			scope *ethereumcorevm.ScopeContext, rData []byte, depth int, err error) {
			// no-op
		},
		CaptureTxEndFunc: func(restGas uint64) {
			// no-op
		},
		CaptureTxStartFunc: func(gasLimit uint64) {
			// no-op
		},
	}
	return mockedEVMLogger
}
