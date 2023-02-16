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

package configuration

import (
	"context"

	"github.com/berachain/stargazer/eth/core"
	"github.com/berachain/stargazer/eth/params"
)

// `plugin` implements the core.ConfigurationPlugin interface.
type plugin struct {
	ctx context.Context
}

// `NewPluginFrom` returns a new plugin instance.
func NewPluginFrom(ctx context.Context) core.ConfigurationPlugin {
	return &plugin{}
}

// `Prepare` implements the core.ConfigurationPlugin interface.
func (p *plugin) Prepare(ctx context.Context) {
	p.ctx = ctx
}

// `ChainConfig` implements the core.ConfigurationPlugin interface.
func (p *plugin) ChainConfig() *params.ChainConfig {
	return params.DefaultChainConfig
}

// `ExtraEips` implements the core.ConfigurationPlugin interface.
func (p *plugin) ExtraEips() []int {
	return []int{}
}
