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
	"github.com/berachain/stargazer/core/precompile/registry"
	"github.com/berachain/stargazer/core/vm"
	"github.com/berachain/stargazer/lib/common"
)

// Compile-time assertion to ensure `controller` adheres to `PrecompileController`.
var _ vm.PrecompileController = (*controller)(nil)

// `controller` is a struct that embeds a `vm.PrecompileRunner` and uses a precompile registry.
type controller struct {
	// `PrecompileRunner` will run the precompile in a custom precompile environment.
	vm.PrecompileRunner

	// `registry` allows the `controller` to search for a precompile container at an address.
	registry *registry.Registry
}

// `NewController` creates and returns a `controller` with the given precompile registry and
// precompile runner.
//
//nolint:revive // this is only used as a `vm.PrecompileController`.
func NewController(registry *registry.Registry, runner vm.PrecompileRunner) *controller {
	return &controller{
		PrecompileRunner: runner,
		registry:         registry,
	}
}

// `Exists` searches the registry at the given `addr` for a precompile container.
//
// `Exists` implements `vm.PrecompileContainer`.
func (c *controller) Exists(addr common.Address) (vm.PrecompileContainer, bool) {
	return c.registry.Lookup(addr)
}
