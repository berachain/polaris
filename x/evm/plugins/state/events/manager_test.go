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

package events_test

import (
	"github.com/berachain/stargazer/testutil"
	"github.com/berachain/stargazer/x/evm/plugins/state"
	"github.com/berachain/stargazer/x/evm/plugins/state/events"
	"github.com/berachain/stargazer/x/evm/plugins/state/events/mock"
	sdk "github.com/cosmos/cosmos-sdk/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	var cem state.ControllableEventManager
	var ctx sdk.Context
	var ldb *mock.LogsDBMock

	BeforeEach(func() {
		ldb = mock.NewEmptyLogsDB()

		ctx = testutil.NewContext()
		ctx.EventManager().EmitEvent(sdk.NewEvent("1"))

		cem = events.NewManagerFrom(ctx.EventManager(), mock.NewPrecompileLogFactory())
		ctx = ctx.WithEventManager(cem)
		Expect(len(ctx.EventManager().Events())).To(Equal(1))
		Expect(cem.Events()).To(HaveLen(1))
	})

	It("should have the right registry key", func() {
		Expect(cem.RegistryKey()).To(Equal("events"))
	})

	It("should correctly snapshot/revert", func() {
		ctx.EventManager().EmitEvent(sdk.NewEvent("2"))
		Expect(len(ctx.EventManager().Events())).To(Equal(2))

		snap := cem.Snapshot()
		ctx.EventManager().EmitEvent(sdk.NewEvent("3"))
		Expect(len(ctx.EventManager().Events())).To(Equal(3))

		cem.RevertToSnapshot(snap)
		Expect(len(ctx.EventManager().Events())).To(Equal(2))
	})

	It("should not build eth logs when not in precompile", func() {
		ctx.EventManager().EmitEvent(sdk.NewEvent("2"))
		Expect(len(ctx.EventManager().Events())).To(Equal(2))
		Expect(len(ldb.AddLogCalls())).To(Equal(0))

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent("3"),
			sdk.NewEvent("4"),
		})
		Expect(len(ctx.EventManager().Events())).To(Equal(4))
		Expect(len(ldb.AddLogCalls())).To(Equal(0))
	})

	It("should panic when building eth logs fails", func() {
		cem.BeginPrecompileExecution(ldb)

		Expect(func() {
			ctx.EventManager().EmitEvent(sdk.NewEvent("non-eth-event"))
		}).To(Panic())
	})

	It("should build eth logs from cosmos events during precompile", func() {
		cem.BeginPrecompileExecution(ldb)

		ctx.EventManager().EmitEvent(sdk.NewEvent("2"))
		Expect(len(ctx.EventManager().Events())).To(Equal(2))
		Expect(len(ldb.AddLogCalls())).To(Equal(1))

		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent("3"),
			sdk.NewEvent("4"),
		})
		Expect(len(ctx.EventManager().Events())).To(Equal(4))
		Expect(len(ldb.AddLogCalls())).To(Equal(3))

		cem.EndPrecompileExecution()

		Expect(func() { cem.Finalize() }).ToNot(Panic())
	})
})
