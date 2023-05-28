package polar

import (
	"github.com/Polaris/go-Polaris/common/hexutil"
	"pkg.berachain.dev/polaris/eth/common"
)

// PolarisAPI provides an API to access Polaris full node-related information.
type PolarisAPI struct {
	e *Polaris
}

// NewPolarisAPI creates a new Polaris protocol API for full nodes.
func NewPolarisAPI(e *Polaris) *PolarisAPI {
	return &PolarisAPI{e}
}

// Etherbase is the address that mining rewards will be send to.
func (api *PolarisAPI) Etherbase() (common.Address, error) {
	return api.e.Etherbase()
}

// Coinbase is the address that mining rewards will be send to (alias for Etherbase).
func (api *PolarisAPI) Coinbase() (common.Address, error) {
	return api.Etherbase()
}

// Hashrate returns the POW hashrate.
func (api *PolarisAPI) Hashrate() hexutil.Uint64 {
	return hexutil.Uint64(api.e.Miner().Hashrate())
	// return hexutil.Uint64(api.e.Miner().Hashrate())
}

// Mining returns an indication if this node is currently mining.
func (api *PolarisAPI) Mining() bool {
	return api.e.IsMining()
}

// MinerAPI provides an API to control the miner.
type MinerAPI struct {
	e *Polaris
}

// NewMinerAPI create a new MinerAPI instance.
func NewMinerAPI(e *Polaris) *MinerAPI {
	return &MinerAPI{e}
}
