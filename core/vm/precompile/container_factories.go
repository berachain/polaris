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
	"cosmossdk.io/errors"
	"github.com/berachain/stargazer/core/vm/precompile/container"
	"github.com/berachain/stargazer/core/vm/precompile/container/types"
	"github.com/berachain/stargazer/lib/common"
	"github.com/berachain/stargazer/types/abi"
)

const (
	statefulContainerFactoryName = `StatefulContainerFactory`
	dynamicContainerFactoryName  = `DynamicContainerFactory`
)

var (
	_ AbstractContainerFactory = (*StatelessContainerFactory)(nil)
	_ AbstractContainerFactory = (*StatelessContainerFactory)(nil)
	_ AbstractContainerFactory = (*DynamicContainerFactory)(nil)
)

// ===========================================================================
// Stateless Container Factory
// ===========================================================================

type StatelessContainerFactory struct{}

func NewStatelessContainerFactory() *StatelessContainerFactory {
	return &StatelessContainerFactory{}
}

func (scf *StatelessContainerFactory) Build(
	bci BaseContractImpl,
) (types.PrecompileContainer, error) {
	pc, ok := bci.(StatelessContractImpl)
	if !ok {
		return nil, ErrNotStatelessPrecompile
	}
	return pc, nil
}

// ===========================================================================
// Stateful Container Factory
// ===========================================================================

type StatefulContainerFactory struct{}

func NewStatefulContainerFactory() *StatefulContainerFactory {
	return &StatefulContainerFactory{}
}

func (scf *StatefulContainerFactory) Build(
	bci BaseContractImpl,
) (types.PrecompileContainer, error) {
	sci, ok := bci.(StatefulContractImpl)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, statefulContainerFactoryName)
	}

	var err error
	var idsToMethods map[common.Hash]*types.PrecompileMethod

	// add precompile methods to Stateful Container, if any
	precompileMethods := sci.PrecompileMethods()
	if precompileMethods != nil {
		// validate precompile methods
		for _, pm := range precompileMethods {
			if err = pm.ValidateBasic(); err != nil {
				return nil, err
			}
		}

		idsToMethods, err = scf.buildIdsToMethods(precompileMethods, sci.ABIMethods())
		if err != nil {
			return nil, err
		}
	}

	return container.NewStatefulContainer(idsToMethods), nil
}

func (scf *StatefulContainerFactory) buildIdsToMethods(
	precompileMethods types.PrecompileMethods,
	abiMethods map[string]abi.Method,
) (map[common.Hash]*types.PrecompileMethod, error) {
	// match every ABI method to corresponding precompile method
	idsToMethods := make(map[common.Hash]*types.PrecompileMethod)
	for name := range abiMethods {
		abiMethod := abiMethods[name]

		// find the corresponding precompile method for abiMethod based on signature
		var precompileMethod *types.PrecompileMethod
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
		idsToMethods[common.BytesToHash(abiMethod.ID)] = precompileMethod
	}
	return idsToMethods, nil
}

// ===========================================================================
// Dynamic Container Factory
// ===========================================================================

type DynamicContainerFactory struct{}

func NewDynamicContainerFactory() *DynamicContainerFactory {
	return &DynamicContainerFactory{}
}

func (dcf *DynamicContainerFactory) Build(
	bci BaseContractImpl,
) (types.PrecompileContainer, error) {
	dci, ok := bci.(DynamicContractImpl)
	if !ok {
		return nil, errors.Wrap(ErrWrongContainerFactory, dynamicContainerFactoryName)
	}

	return NewStatefulContainerFactory().Build(dci)
}
