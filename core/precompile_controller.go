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
	"github.com/berachain/stargazer/core/precompile"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
)

// Compile-time assertion to ensure `precompileController` adheres to `PrecompileController`.
var _ vm.PrecompileController = (*precompileController)(nil)

// `precompileController` is a struct that embeds a `vm.PrecompileRunner` and uses a precompile precompile.
type precompileController struct {
	// `PrecompileRunner` will run the precompile in a custom precompile environment.
	vm.PrecompileRunner

	// `registry` allows the `precompileController` to search for a precompile container at an address.
	registry *precompile.Registry
}

// `NewPrecompileController` creates and returns a `precompileController` with the given precompile
// registry and precompile runner.
//
//nolint:revive // this is only used as a `vm.PrecompileController`.
func NewPrecompileController(
	registry *precompile.Registry, runner vm.PrecompileRunner,
) *precompileController {
	return &precompileController{
		PrecompileRunner: runner,
		registry:         registry,
	}
}

// `Exists` searches the registry at the given `addr` for a precompile container.
//
// `Exists` implements `vm.PrecompileContainer`.
func (c *precompileController) Exists(addr common.Address) (vm.PrecompileContainer, bool) {
	return c.registry.Lookup(addr)
}
