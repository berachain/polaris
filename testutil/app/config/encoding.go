package config

import (
	"cosmossdk.io/simapp/params"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	evmante "pkg.berachain.dev/stargazer/x/evm/ante"
)

func MakeEncodingConfig() params.EncodingConfig {
	cdc := codec.NewLegacyAmino()
	interfaceRegistry := types.NewInterfaceRegistry()
	codec := codec.NewProtoCodec(interfaceRegistry)

	txConfig := tx.NewTxConfig(
		codec,
		append(tx.DefaultSignModes, []signingtypes.SignMode{evmante.SignMode_SIGN_MODE_ETHEREUM}...),
		[]signing.SignModeHandler{evmante.SignModeEthTxHandler{}}...,
	)

	return params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             codec,
		TxConfig:          txConfig,
		Amino:             cdc,
	}
}
