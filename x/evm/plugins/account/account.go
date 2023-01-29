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

package account

import (
	"context"

	"github.com/berachain/stargazer/eth/core/state"
	"github.com/berachain/stargazer/lib/common"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ state.AccountPlugin = (*Plugin)(nil)

type Plugin struct {
	// keepers used for balance and account information.
	ak AccountKeeper

	// kvstore cachekv.Store
}

func NewPlugin(ak AccountKeeper) *Plugin {
	return &Plugin{
		ak: ak,
	}
}

func (p *Plugin) CreateAccount(ctx context.Context, addr common.Address) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	acc := p.ak.NewAccountWithAddress(sCtx, addr[:])

	// save the new account in the account keeper
	p.ak.SetAccount(sCtx, acc)
}

func (p *Plugin) HasAccount(ctx context.Context, addr common.Address) bool {
	return p.ak.HasAccount(sdk.UnwrapSDKContext(ctx), addr[:])
}

// `DeleteAccount` deletes the account associated with the given address.
func (p *Plugin) DeleteAccount(ctx context.Context, addr common.Address) {
	acc := p.ak.GetAccount(sdk.UnwrapSDKContext(ctx), addr[:])
	if acc == nil {
		return
	}
	p.ak.RemoveAccount(sdk.UnwrapSDKContext(ctx), acc)
}

// `GetNonce` returns the nonce of the account associated with the given address.
func (p *Plugin) GetNonce(ctx context.Context, addr common.Address) uint64 {
	acc := p.ak.GetAccount(sdk.UnwrapSDKContext(ctx), addr[:])
	if acc == nil {
		return 0
	}
	return acc.GetSequence()
}

// `SetNonce` sets the nonce of the account associated with the given address.
func (p *Plugin) SetNonce(ctx context.Context, addr common.Address, nonce uint64) {
	sCtx := sdk.UnwrapSDKContext(ctx)
	acc := p.ak.GetAccount(sCtx, addr[:])
	if acc == nil {
		acc = p.ak.NewAccountWithAddress(sCtx, addr[:])
	}

	if err := acc.SetSequence(nonce); err != nil {
		panic(err)
	}

	p.ak.SetAccount(sCtx, acc)
}

func (p *Plugin) RevertToSnapshot(id int) {
	// noop
}

func (p *Plugin) Snapshot() int {
	// noop
	return 0
}
