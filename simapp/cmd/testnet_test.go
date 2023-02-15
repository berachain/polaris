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

//nolint:mnd // from sdk.
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/spf13/viper"

	"cosmossdk.io/simapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutiltest "github.com/cosmos/cosmos-sdk/x/genutil/client/testutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/staking"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Testnet command", func() {
	var home string
	var dir *os.File
	BeforeEach(func() {
		home = os.TempDir()
		dir = os.NewFile(0, os.DevNull)
		v, err := dir.ReadDir(-1)
		if err == nil {
			Expect(dir.Seek(0, 0)).To(BeNil())
		}
		if len(v) == 0 {
			os.RemoveAll(home)
		}
	})

	It("should generate a testnet with given options", func() {
		encodingConfig := moduletestutil.MakeTestEncodingConfig(staking.AppModuleBasic{},
			auth.AppModuleBasic{})
		logger := log.NewNopLogger()
		cfg, err := genutiltest.CreateDefaultCometConfig(home)
		Expect(err).To(BeNil())
		err = genutiltest.ExecInitCmd(simapp.ModuleBasics, home, encodingConfig.Codec)
		Expect(err).To(BeNil())

		serverCtx := server.NewContext(viper.New(), cfg, logger)
		clientCtx := client.Context{}.
			WithCodec(encodingConfig.Codec).
			WithHomeDir(home).
			WithTxConfig(encodingConfig.TxConfig)

		ctx := context.Background()
		ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
		ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
		cmd := testnetInitFilesCmd(simapp.ModuleBasics, banktypes.GenesisBalancesIterator{})
		cmd.SetArgs([]string{fmt.Sprintf("--%s=test", flags.FlagKeyringBackend),
			fmt.Sprintf("--output-dir=%s", home)})
		err = cmd.ExecuteContext(ctx)
		Expect(err).To(BeNil())

		genFile := cfg.GenesisFile()
		appState, _, err := genutiltypes.GenesisStateFromGenFile(genFile)
		Expect(err).To(BeNil())

		bankGenState := banktypes.GetGenesisStateFromAppState(encodingConfig.Codec, appState)
		Expect(bankGenState.Supply.IsZero()).To(BeFalse())
	})
})
