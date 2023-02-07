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
	libtypes "github.com/berachain/stargazer/lib/types"
	"github.com/berachain/stargazer/testutil"
	"github.com/berachain/stargazer/x/evm/plugin/state/events"
	sdk "github.com/cosmos/cosmos-sdk/types"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Manager", func() {
	var cem libtypes.Controllable[string]
	var ctx sdk.Context

	It("should correctly control the event manager", func() {
		ctx = testutil.NewContext()

		// before the sdb tx
		ctx.EventManager().EmitEvent(sdk.NewEvent("1"))

		// sdb setup
		cem = events.NewManagerFrom(ctx.EventManager())

		// check the controllable event manager is hooked up to context
		Expect(len(ctx.EventManager().Events())).To(Equal(1))

		// chcek registry key
		Expect(cem.RegistryKey()).To(Equal("events"))

		// add to event manager
		ctx.EventManager().EmitEvent(sdk.NewEvent("2"))
		Expect(len(ctx.EventManager().Events())).To(Equal(2))

		snap := cem.Snapshot()
		ctx.EventManager().EmitEvent(sdk.NewEvent("3"))
		Expect(len(ctx.EventManager().Events())).To(Equal(3))

		cem.RevertToSnapshot(snap)
		Expect(len(ctx.EventManager().Events())).To(Equal(2))
	})
})
