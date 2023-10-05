package types

import (
	"math"

	storetypes "cosmossdk.io/store/types"
)

// noopGasMeter is a gas meter that implements the fuck it we ball
// software design pattern. On a more serious note, since we let
// geth handle gas consumption, but the context forces us to use a GasKV,
// the golang race detector will tell us our miner and txpool gorountines
// are racing each other, to consume gas on the gas meter. We could, add a lock
// but why do that and introduce lock contention, when we could just remove
// the variable the data race is occuring on, since we aren't using the
// data anyways.
type noopGasMeter struct{}

// NewnoopGasMeter returns a new gas meter without a limit.
func NewNoopGasMeter() storetypes.GasMeter {
	return &noopGasMeter{}
}

// GasConsumed returns the gas consumed from the GasMeter.
func (g *noopGasMeter) GasConsumed() storetypes.Gas {
	return 0
}

// GasConsumedToLimit returns the gas consumed from the GasMeter since the gas is not confined to a limit.
func (g *noopGasMeter) GasConsumedToLimit() storetypes.Gas {
	return 0
}

// GasRemaining returns MaxUint64 since limit is not confined in noopGasMeter.
func (g *noopGasMeter) GasRemaining() storetypes.Gas {
	return math.MaxUint64
}

// Limit returns MaxUint64 since limit is not confined in noopGasMeter.
func (g *noopGasMeter) Limit() storetypes.Gas {
	return math.MaxUint64
}

// ConsumeGas adds the given amount of gas to the gas consumed and panics if it overflows the limit.
func (g *noopGasMeter) ConsumeGas(storetypes.Gas, string) {}

// RefundGas is a no-op.
func (g *noopGasMeter) RefundGas(storetypes.Gas, string) {}

// IsPastLimit always returns false since the gas limit is not confined.
func (g *noopGasMeter) IsPastLimit() bool {
	return false
}

// IsOutOfGas returns false since the gas limit is not confined.
func (g *noopGasMeter) IsOutOfGas() bool {
	return false
}

// String returns the noopGasMeter's gas consumed.
func (g *noopGasMeter) String() string {
	return "Fishing? In Lebanon? Me? Never..."
}
