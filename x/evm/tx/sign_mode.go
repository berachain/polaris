package tx

import (
	"fmt"

	errorsmod "cosmossdk.io/errors"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	types "github.com/cosmos/cosmos-sdk/types/tx"
	signingtypes "github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	evmtypes "pkg.berachain.dev/stargazer/x/evm/types"
)

var _ signing.SignModeHandler = signModeDirectAuxHandler{}

// signModeDirectAuxHandler defines the SIGN_MODE_DIRECT_AUX SignModeHandler
type signModeDirectAuxHandler struct{}

// DefaultMode implements SignModeHandler.DefaultMode
func (signModeDirectAuxHandler) DefaultMode() signingtypes.SignMode {
	return signingtypes.SignMode_SIGN_MODE_DIRECT_AUX
}

// Modes implements SignModeHandler.Modes
func (signModeDirectAuxHandler) Modes() []signingtypes.SignMode {
	return []signingtypes.SignMode{signingtypes.SignMode_SIGN_MODE_DIRECT_AUX}
}

// GetSignBytes implements SignModeHandler.GetSignBytes
func (signModeDirectAuxHandler) GetSignBytes(
	mode signingtypes.SignMode, data signing.SignerData, tx sdk.Tx,
) ([]byte, error) {
	if mode != signingtypes.SignMode_SIGN_MODE_DIRECT_AUX {
		return nil, fmt.Errorf("expected %s, got %s", signingtypes.SignMode_SIGN_MODE_DIRECT_AUX, mode)
	}

	pkAny, err := codectypes.NewAnyWithValue(data.PubKey)
	if err != nil {
		return nil, err
	}

	addr := data.Address
	if addr == "" {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "got empty address in %s handler", signingtypes.SignMode_SIGN_MODE_DIRECT_AUX)
	}

	ethTx := tx.GetMsgs()[0].(*evmtypes.EthTransactionRequest).AsTransaction()

	signDocDirectAux := types.SignDocDirectAux{
		BodyBytes:     ethTx.Data(),
		ChainId:       data.ChainID,
		AccountNumber: data.AccountNumber,
		Sequence:      ethTx.Nonce(),
		Tip:           &types.Tip{},
		PublicKey:     pkAny,
	}

	return signDocDirectAux.Marshal()
}
