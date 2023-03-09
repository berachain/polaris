package node

import (
	"time"

	"github.com/rs/zerolog/log"
	"pkg.berachain.dev/polaris/playground/chain"
)

// `Runner` is the interface for the node runner.
type Runner interface {
	Start() error
}

// `runner` is the main node runner.
type runner struct {
	blocktime  time.Duration
	rpcService *rpcService
	stop       chan struct{}
	chain      *chain.Playground
}

// `NewRunner` creates a new node runner.
func NewRunner(blocktime time.Duration) Runner {
	// Setup RPC
	rpcService := NewRPCService()

	// Setup Mempool
	mempool := NewMempool()
	return &runner{
		blocktime:  blocktime,
		rpcService: rpcService,
		stop:       make(chan struct{}),
		chain: chain.NewPlayground(
			mempool,
		),
	}
}

// `Start` starts the node
func (r *runner) Start() error {
	for {
		select {
		case <-r.stop:
			log.Info().Msg("chain shutting down")
			close(r.stop)
			return nil
		case <-time.After(r.blocktime):
			log.Error().Msg("producing block")
			if _, err := r.chain.ProduceBlock(); err != nil {
				log.Error().Err(err).Msg("failed to produce block")
				return err
			}
		}
	}
}

// `Stop` stops the node.
func (r *runner) Stop() {
	r.stop <- struct{}{}
}
