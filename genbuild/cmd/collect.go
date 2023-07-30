package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	cosmosflags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/cosmos/cosmos-sdk/types/tx"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
	polariscryptocodec "pkg.berachain.dev/polaris/cosmos/crypto/codec"
	polaristypes "pkg.berachain.dev/polaris/cosmos/types"
	"pkg.berachain.dev/polaris/genbuild/utils"
)

func NewCollectCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "collect",
		Short: "Collect genesis txs and update genesis.json file",
		Args:  cobra.NoArgs,
		Run:   runCollect,
	}

	cmd.Flags().String(FlagHomeDir, "", "home directory")
	if err := cmd.MarkFlagRequired(FlagHomeDir); err != nil {
		return nil, err
	}

	return cmd, nil
}

func runCollect(cmd *cobra.Command, args []string) {
	homeDir, err := cmd.Flags().GetString(FlagHomeDir)
	if err != nil {
		log.Fatal(err)
	}

	keyAddresses, err := getKeyAddresses(homeDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, keyAddress := range keyAddresses {
		if err := utils.RunV(
			"polard", []string{"genesis", "add-genesis-account", keyAddress, fmt.Sprintf("%f%s", initialBalance, denom)},
			map[string]string{
				cosmosflags.FlagHome: homeDir,
				// TODO: remove append, check existing keys on current node and ignore accordingly
				"append": "",
			},
		); err != nil {
			log.Fatal(err)
		}
	}

	if err := utils.RunV(
		"polard", []string{"genesis", "collect-gentxs"}, map[string]string{
			cosmosflags.FlagHome: homeDir,
		},
	); err != nil {
		log.Fatal(err)
	}

	if err := utils.RunV(
		"polard", []string{"genesis", "validate-genesis"}, map[string]string{
			cosmosflags.FlagHome: homeDir,
		},
	); err != nil {
		log.Fatal(err)
	}
}

func getKeyAddresses(homeDir string) ([]string, error) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	// cosmos.staking.v1beta1.MsgCreateValidator
	stakingtypes.RegisterInterfaces(interfaceRegistry)
	// cosmos.crypto.ed25519.PubKey
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	// polaris.crypto.ethsecp256k1.v1.PubKey
	polariscryptocodec.RegisterInterfaces(interfaceRegistry)

	cdc := codec.NewProtoCodec(interfaceRegistry)

	genTxsDir := filepath.Join(homeDir, "config", "gentx")
	genTxFiles, err := os.ReadDir(genTxsDir)
	if err != nil {
		return nil, err
	}

	var addresses []string
	for _, genTxFile := range genTxFiles {
		if genTxFile.IsDir() {
			continue
		}
		if !strings.HasSuffix(genTxFile.Name(), ".json") {
			continue
		}
		jsonRawTx, err := os.ReadFile(filepath.Join(genTxsDir, genTxFile.Name()))
		if err != nil {
			return nil, err
		}
		var genTx tx.Tx
		if err := cdc.UnmarshalJSON(jsonRawTx, &genTx); err != nil {
			return nil, err
		}
		publicKey := genTx.AuthInfo.SignerInfos[0].PublicKey.GetCachedValue().(cryptotypes.PubKey)
		address, err := bech32.ConvertAndEncode(polaristypes.Bech32PrefixAccAddr, publicKey.Address())
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, err
}
