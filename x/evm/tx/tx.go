package tx

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"pkg.berachain.dev/stargazer/x/evm/types"
)

// `SignMode_SIGN_MODE_ETHEREUM` defines the sign mode for Ethereum transactions.
const SignMode_SIGN_MODE_ETHEREUM signingtypes.SignMode = 42069

// `CustomSignModeHandlers` returns the custom sign mode handlers for the EVM module.
func CustomSignModeHandlers() []signing.SignModeHandler {
	return []signing.SignModeHandler{
		SignModeEthTxHandler{},
	}
}

var _ signing.SignModeHandler = (*SignModeEthTxHandler)(nil)

// `SignModeEthTx` defines the sign mode for Ethereum transactions.
type SignModeEthTxHandler struct{}

// `Modes` returns the sign modes for the EVM module.
func (s SignModeEthTxHandler) Modes() []signingtypes.SignMode {
	return []signingtypes.SignMode{s.DefaultMode()}
}

// `DefaultMode` returns the default sign mode for the EVM module.
func (s SignModeEthTxHandler) DefaultMode() signingtypes.SignMode {
	return SignMode_SIGN_MODE_ETHEREUM
}

// `GetSignBytes` returns the sign bytes for the given sign mode and transaction.
func (s SignModeEthTxHandler) GetSignBytes(mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx) ([]byte, error) {
	ethTx, ok := tx.GetMsgs()[0].(*types.EthTransactionRequest)
	if !ok {
		return nil, errors.New("expected EthTransactionRequest")
	}
	return ethTx.GetSignBytes()
}
