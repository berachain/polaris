package core

import (
	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/eth/core/state/plugin"
	"github.com/berachain/stargazer/eth/core/vm"
)

type StateDBFactory struct {
	sp StatePlugin
	lp state.LogsPlugin
	rp state.RefundPlugin
}

func NewStateDBFactory(sp StatePlugin) *StateDBFactory {
	return &StateDBFactory{
		sp: sp,
		lp: plugin.NewLogs(),
		rp: plugin.NewRefund(),
	}
}

func (f *StateDBFactory) Build() (vm.StargazerStateDB, error) {
	return state.NewStateDB(f.sp, f.lp, f.rp)
}
