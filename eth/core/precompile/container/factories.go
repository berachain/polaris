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

package container

import (
	"github.com/berachain/stargazer/eth/core/vm"
	"github.com/berachain/stargazer/eth/types/abi"
	"github.com/berachain/stargazer/lib/errors"
	"github.com/berachain/stargazer/lib/utils"
)

const (
	// container impl names stored as constants, to be used in error messages.
	statelessContainerName = `StatelessContainerImpl`
	statefulContainerName  = `StatefulContainerImpl`
	dynamicContainerName   = `DynamicContainerImpl`
)

// Compile-time assertions to ensure these container factories adhere to `AbstractFactory`.
var (
	_ AbstractFactory = (*StatelessFactory)(nil)
	_ AbstractFactory = (*StatefulFactory)(nil)
	_ AbstractFactory = (*DynamicFactory)(nil)
)

// ===========================================================================
// Stateless Container Factory
// ===========================================================================

// `StatelessFactory` is used to build stateless precompile containers.
type StatelessFactory struct{}

// `NewStatelessFactory` creates and returns a new `StatelessFactory`.
func NewStatelessFactory() *StatelessFactory {
	return &StatelessFactory{}
}

// `Build` returns a stateless precompile container for the given base contract implememntation.
// This function will return an error if the given contract is not a stateless implementation.
//
// `Build` implements `AbstractFactory`.
func (sf *StatelessFactory) Build(
	rj vm.RegistrablePrecompile,
) (vm.PrecompileContainer, error) {
	pc, ok := utils.GetAs[StatelessPrecompileImpl](rj)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, statelessContainerName)
	}
	return pc, nil
}

// ===========================================================================
// Stateful Container Factory
// ===========================================================================

// `StatefulFactory` is used to build stateful precompile containers.
type StatefulFactory struct {
}

// `NewStatefulFactory` creates and returns a new `StatefulFactory`.
func NewStatefulFactory() *StatefulFactory {
	return &StatefulFactory{}
}

// `Build` returns a stateful precompile container for the given base contract implementation.
// This function will return an error if the given contract is not a stateful implementation.
//
// `Build` implements `AbstractFactory`.
func (sf *StatefulFactory) Build(
	rj vm.RegistrablePrecompile,
) (vm.PrecompileContainer, error) {
	sci, ok := utils.GetAs[StatefulPrecompileImpl](rj)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, statefulContainerName)
	}

	var err error

	// add precompile methods to stateful container, if any exist
	var idsToMethods map[string]*Method
	if precompileMethods := sci.PrecompileMethods(); precompileMethods != nil {
		idsToMethods, err = sf.buildIdsToMethods(precompileMethods, sci.ABIMethods())
		if err != nil {
			return nil, err
		}
	}

	return NewStateful(rj, idsToMethods), nil
}

// `buildIdsToMethods` builds the stateful precompile container for the given `precompileMethods`
// and `abiMethods`. This function will return an error if every method in `abiMethods` does not
// have a valid, corresponding `Method`.
func (sf *StatefulFactory) buildIdsToMethods(
	precompileMethods Methods,
	abiMethods map[string]abi.Method,
) (map[string]*Method, error) {
	// validate precompile methods
	for _, pm := range precompileMethods {
		if err := pm.ValidateBasic(); err != nil {
			return nil, err
		}
	}

	// match every ABI method to corresponding precompile method
	idsToMethods := make(map[string]*Method)
	for name := range abiMethods {
		abiMethod := abiMethods[name]

		// find the corresponding precompile method for abiMethod based on signature
		var precompileMethod *Method
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

// `DynamicFactory` is used to build dynamic precompile containers.
type DynamicFactory struct {
	*StatefulFactory
}

// `NewDynamicFactory` creates and returns a new `DynamicFactory` for the given
// log registry `lr`.
func NewDynamicFactory() *DynamicFactory {
	return &DynamicFactory{
		StatefulFactory: NewStatefulFactory(),
	}
}

// `Build` returns a dynamic precompile container for the given base contract implememntation.
// This function will return an error if the given contract is not a dyanmic implementation.
//
// `Build` implements `AbstractFactory`.
func (dcf *DynamicFactory) Build(
	rj vm.RegistrablePrecompile,
) (vm.PrecompileContainer, error) {
	dci, ok := utils.GetAs[DynamicPrecompileImpl](rj)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, dynamicContainerName)
	}

	return dcf.StatefulFactory.Build(dci)
}
