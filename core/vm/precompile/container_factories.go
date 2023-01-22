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

package precompile

import (
	"github.com/berachain/stargazer/core/vm/precompile/container"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/core/vm/precompile/log"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/lib/utils"
	"github.com/berachain/stargazer/types/abi"
)

const (
	// container impl names stored as constants, to be used in error messages.
	statelessContainerName = `StatelessContainerImpl`
	statefulContainerName  = `StatefulContainerImpl`
	dynamicContainerName   = `DynamicContainerImpl`
)

// Compile-time assertions to ensure these container factories adhere to
// `AbstractContainerFactory`.
var (
	_ AbstractContainerFactory = (*StatelessContainerFactory)(nil)
	_ AbstractContainerFactory = (*StatefulContainerFactory)(nil)
	_ AbstractContainerFactory = (*DynamicContainerFactory)(nil)
)

// ===========================================================================
// Stateless Container Factory
// ===========================================================================

// `StatelessContainerFactory` is used to build stateless precompile containers.
type StatelessContainerFactory struct{}

// `NewStatelessContainerFactory` creates and returns a new `StatelessContainerFactory`.
func NewStatelessContainerFactory() *StatelessContainerFactory {
	return &StatelessContainerFactory{}
}

// `Build` returns a stateless precompile container for the given base contract implememntation.
// This function will return an error if the given contract is not a stateless implementation.
//
// `Build` implements `AbstractContainerFactory`.
func (scf *StatelessContainerFactory) Build(
	bci BaseContractImpl,
) (types.PrecompileContainer, error) {
	pc, ok := utils.GetAs[StatelessContractImpl](bci)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, statelessContainerName)
	}
	return pc, nil
}

// ===========================================================================
// Stateful Container Factory
// ===========================================================================

// `StatefulContainerFactory` is used to build stateful precompile containers.
type StatefulContainerFactory struct {
	// `lr` is used to register stateful precompiles' event logs, if any.
	lr *LogRegistry
}

// `NewStatefulContainerFactory` creates and returns a new `StatefulContainerFactory`.
func NewStatefulContainerFactory(lr *LogRegistry) *StatefulContainerFactory {
	return &StatefulContainerFactory{
		lr: lr,
	}
}

// `Build` returns a stateful precompile container for the given base contract implementation.
// This function will return an error if the given contract is not a stateful implementation.
//
// `Build` implements `AbstractContainerFactory`.
func (scf *StatefulContainerFactory) Build(
	bci BaseContractImpl,
) (types.PrecompileContainer, error) {
	sci, ok := utils.GetAs[StatefulContractImpl](bci)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, statefulContainerName)
	}

	var err error

	// add precompile methods to stateful container, if any exist
	var idsToMethods map[string]*types.Method
	if precompileMethods := sci.PrecompileMethods(); precompileMethods != nil {
		idsToMethods, err = scf.buildIdsToMethods(precompileMethods, sci.ABIMethods())
		if err != nil {
			return nil, err
		}
	}

	// add precompile events to stateful container, if any exist
	if precompileEvents := sci.ABIEvents(); precompileEvents != nil {
		customValueDecoders := sci.CustomValueDecoders()
		for _, abiEvent := range precompileEvents {
			// add value decoders if the event is custom
			var customEventValueDecoders log.ValueDecoders
			if customValueDecoders != nil {
				customEventValueDecoders = customValueDecoders[EventType(abiEvent.Name)]
			}
			// register the event to the precompiles' log registry
			err = scf.lr.RegisterEvent(sci.Address(), abiEvent, customEventValueDecoders)
			if err != nil {
				return nil, err
			}
		}
	}

	return container.NewStatefulContainer(idsToMethods), nil
}

// `buildIdsToMethods` builds the stateful precompile container for the given `precompileMethods`
// and `abiMethods`. This function will return an error if every method in `abiMethods` does not
// have a valid, corresponding `types.Method`.
func (scf *StatefulContainerFactory) buildIdsToMethods(
	precompileMethods types.Methods,
	abiMethods map[string]abi.Method,
) (map[string]*types.Method, error) {
	// validate precompile methods
	for _, pm := range precompileMethods {
		if err := pm.ValidateBasic(); err != nil {
			return nil, err
		}
	}

	// match every ABI method to corresponding precompile method
	idsToMethods := make(map[string]*types.Method)
	for name := range abiMethods {
		abiMethod := abiMethods[name]

		// find the corresponding precompile method for abiMethod based on signature
		var precompileMethod *types.Method
		i := 0
		for ; i < len(precompileMethods); i++ {
			if precompileMethods[i].AbiSig == abiMethod.Sig {
				precompileMethod = precompileMethods[i]
				break
			}
		}
		if i == len(precompileMethods) {
			return nil, errors.Wrap(ErrNoPrecompileMethodForABIMethod, abiMethod.Sig)
		}

		// attach the ABI method to the precompile method for stateful container to handle
		precompileMethod.AbiMethod = &abiMethod
		idsToMethods[utils.UnsafeBytesToStr(abiMethod.ID)] = precompileMethod
	}
	return idsToMethods, nil
}

// ===========================================================================
// Dynamic Container Factory
// ===========================================================================

// `DynamicContainerFactory` is used to build dynamic precompile containers.
type DynamicContainerFactory struct {
	*StatefulContainerFactory
}

// `NewDynamicContainerFactory` creates and returns a new `DynamicContainerFactory` for the given
// log registry `lr`.
func NewDynamicContainerFactory(lr *LogRegistry) *DynamicContainerFactory {
	return &DynamicContainerFactory{
		StatefulContainerFactory: NewStatefulContainerFactory(lr),
	}
}

// `Build` returns a dynamic precompile container for the given base contract implememntation.
// This function will return an error if the given contract is not a dyanmic implementation.
//
// `Build` implements `AbstractContainerFactory`.
func (dcf *DynamicContainerFactory) Build(
	bci BaseContractImpl,
) (types.PrecompileContainer, error) {
	dci, ok := utils.GetAs[DynamicContractImpl](bci)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, dynamicContainerName)
	}

	return dcf.StatefulContainerFactory.Build(dci)
}
