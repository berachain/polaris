package core

import (
	"errors"

	"github.com/berachain/polaris/eth/core/state"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

// WriteGenesisBlock inserts the genesis block into the blockchain.
func (bc *blockchain) WriteGenesisBlock(block *ethtypes.Block) error {
	// Get the state with the latest finalize block context.
	sp := bc.spf.NewPluginWithMode(state.Genesis)
	state := state.NewStateDB(sp, bc.pp)

	// TODO: add more validation here.
	if block.NumberU64() != 0 {
		return errors.New("not the genesis block")
	}
	_, err := bc.WriteBlockAndSetHead(block, nil, nil, state, true)
	return err
}
