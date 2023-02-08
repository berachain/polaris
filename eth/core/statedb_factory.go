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

package core

import (
	"github.com/berachain/stargazer/eth/core/state"
<<<<<<< HEAD
	plugin "github.com/berachain/stargazer/eth/core/state/journal"
=======
	"github.com/berachain/stargazer/eth/core/state/plugin"
>>>>>>> processor
	"github.com/berachain/stargazer/eth/core/vm"
)

type StateDBFactory struct {
	sp StatePlugin
<<<<<<< HEAD
	lj state.LogsJournal
	rj state.RefundJournal
=======
	lp state.LogsPlugin
	rp state.RefundPlugin
>>>>>>> processor
}

func NewStateDBFactory(sp StatePlugin) *StateDBFactory {
	return &StateDBFactory{
		sp: sp,
<<<<<<< HEAD
		lj: plugin.NewLogs(),
		rj: plugin.NewRefund(),
=======
		lp: plugin.NewLogs(),
		rp: plugin.NewRefund(),
>>>>>>> processor
	}
}

func (f *StateDBFactory) Build() (vm.StargazerStateDB, error) {
<<<<<<< HEAD
	return state.NewStateDB(f.sp, f.lj, f.rj)
=======
	return state.NewStateDB(f.sp, f.lp, f.rp)
>>>>>>> processor
}
