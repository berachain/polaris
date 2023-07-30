package cmd

import (
	"fmt"
	"log"

	cosmosflags "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	genutilflags "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cobra"
	"pkg.berachain.dev/polaris/genbuild/utils"
)

func NewInitCmd() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "init moniker",
		Short: "Initialize the node",
		Args:  cobra.ExactArgs(1),
		Run:   runInit,
	}

	cmd.Flags().String(FlagChainId, "", "chain id")
	if err := cmd.MarkFlagRequired(FlagChainId); err != nil {
		return nil, err
	}

	cmd.Flags().String(FlagHomeDir, "", "home directory")
	if err := cmd.MarkFlagRequired(FlagHomeDir); err != nil {
		return nil, err
	}

	return cmd, nil
}

const (
	denom          = "abera"
	keyName        = "dev"
	initialBalance = 1e26
	bondAmount     = 1e21
)

func runInit(cmd *cobra.Command, args []string) {
	moniker := args[0]
	chainId, err := cmd.Flags().GetString(FlagChainId)
	if err != nil {
		log.Fatal(err)
	}
	homeDir, err := cmd.Flags().GetString(FlagHomeDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := utils.RunV(
		"polard", []string{"init", moniker}, map[string]string{
			cosmosflags.FlagHome:              homeDir,
			cosmosflags.FlagChainID:           chainId,
			genutilflags.FlagDefaultBondDenom: denom,
			genutilflags.FlagOverwrite:        "",
		},
	); err != nil {
		log.Fatal(err)
	}

	if err := utils.RunV(
		"polard", []string{"config", "set", "client", cosmosflags.FlagKeyringBackend, keyring.BackendTest},
		map[string]string{
			cosmosflags.FlagHome: homeDir,
		}); err != nil {
		log.Fatal(err)
	}

	if err := utils.RunV(
		"polard", []string{"config", "set", "client", cosmosflags.FlagChainID, chainId},
		map[string]string{
			cosmosflags.FlagHome: homeDir,
		}); err != nil {
		log.Fatal(err)
	}

	if err := utils.RunV(
		"polard", []string{"keys", "add", keyName},
		map[string]string{
			cosmosflags.FlagHome:           homeDir,
			cosmosflags.FlagKeyringBackend: keyring.BackendTest,
			cosmosflags.FlagKeyType:        fmt.Sprintf("eth_%s", string(hd.Secp256k1Type)),
		}); err != nil {
		log.Fatal(err)
	}

	if err := utils.RunV(
		"polard", []string{"genesis", "add-genesis-account", keyName, fmt.Sprintf("%f%s", initialBalance, denom)},
		map[string]string{
			cosmosflags.FlagHome:           homeDir,
			cosmosflags.FlagKeyringBackend: keyring.BackendTest,
		}); err != nil {
		log.Fatal(err)
	}

	if err := utils.RunV(
		"polard", []string{"genesis", "gentx", keyName, fmt.Sprintf("%f%s", bondAmount, denom)},
		map[string]string{
			cosmosflags.FlagHome:           homeDir,
			cosmosflags.FlagKeyringBackend: keyring.BackendTest,
			cosmosflags.FlagChainID:        chainId,
		}); err != nil {
		log.Fatal(err)
	}
}
