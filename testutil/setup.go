// Copyright (C) 2022, Berachain Foundation. All rights reserved.
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

package testutil

import (
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/berachain/stargazer/common"
)

var (
	AccKey     = sdk.NewKVStoreKey("acc")
	BankKey    = sdk.NewKVStoreKey("bank")
	EvmKey     = sdk.NewKVStoreKey("evm")
	StakingKey = sdk.NewKVStoreKey("staking")
	Alice      = common.BytesToAddress([]byte("alice"))
	Bob        = common.BytesToAddress([]byte("bob"))
)

func SetupMinimalKeepers() (sdk.Context,
	authkeeper.AccountKeeper, bankkeeper.BaseKeeper, stakingkeeper.Keeper) {
	dbm := db.NewMemDB()
	ms := store.NewCommitMultiStore(dbm)

	ms.MountStoreWithDB(AccKey, storetypes.StoreTypeIAVL, dbm)
	ms.MountStoreWithDB(BankKey, storetypes.StoreTypeIAVL, dbm)
	ms.MountStoreWithDB(EvmKey, storetypes.StoreTypeIAVL, dbm)
	ms.MountStoreWithDB(StakingKey, storetypes.StoreTypeIAVL, dbm)

	err := ms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	ctx := sdk.NewContext(ms, tmproto.Header{}, false, log.TestingLogger())

	encodingConfig := testutil.MakeTestEncodingConfig(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
	)

	ak := authkeeper.NewAccountKeeper(
		encodingConfig.Codec,
		AccKey,
		authtypes.ProtoBaseAccount,
		map[string][]string{
			stakingtypes.NotBondedPoolName: {authtypes.Minter, authtypes.Burner},
			stakingtypes.BondedPoolName:    {authtypes.Minter, authtypes.Burner},
			"evm":                          {authtypes.Minter, authtypes.Burner},
			"staking":                      {authtypes.Minter, authtypes.Burner},
		},
		"bera",
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	ak.SetModuleAccount(ctx,
		authtypes.NewEmptyModuleAccount("evm", authtypes.Minter, authtypes.Burner))

	bk := bankkeeper.NewBaseKeeper(
		encodingConfig.Codec,
		BankKey,
		ak,
		nil,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	sk := stakingkeeper.NewKeeper(
		encodingConfig.Codec,
		StakingKey,
		ak,
		bk,
		authtypes.NewModuleAddress(govtypes.ModuleName).String(),
	)

	return ctx, ak, bk, *sk
}
