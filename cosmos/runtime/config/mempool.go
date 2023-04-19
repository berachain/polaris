package config

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/skip-mev/pob/mempool"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/x/evm/types"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/core/precompile"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

var _ mempool.Config = (*MempoolConfig)(nil)

type (
	MempoolConfig struct {
		builderContract precompile.StatefulImpl
		txDecoder       sdk.TxDecoder
		contractABI     abi.ABI
		host            Serializer
		evmDenom        string
	}

	AuctionBidInfo struct {
		Transactions [][]byte
		Bid          sdk.Coin
		Timeout      uint64
	}

	Serializer interface {
		Serialize(tx *coretypes.Transaction) ([]byte, error)
		SerializeToSdkTx(tx *coretypes.Transaction) (sdk.Tx, error)
	}
)

// NewMempoolConfig returns a new instance of the mempool config.
func NewMempoolConfig(builderContract precompile.StatefulImpl, txDecoder sdk.TxDecoder, host Serializer, denom string) *MempoolConfig {
	contractABI, err := abi.JSON(strings.NewReader(bindings.BuilderModuleMetaData.ABI))
	if err != nil {
		panic(err)
	}

	return &MempoolConfig{
		builderContract: builderContract,
		txDecoder:       txDecoder,
		contractABI:     contractABI,
		host:            host,
		evmDenom:        denom,
	}
}

// IsAuctionTx defines a function that returns true iff a transaction is an
// auction bid transaction.
func (mempool *MempoolConfig) IsAuctionTx(tx sdk.Tx) (bool, error) {
	// Ensure the transcaction is an EthTransactionRequest
	ethTx, err := getEthTransactionRequest(tx)
	if err != nil {
		return false, err
	}

	if ethTx == nil {
		return false, nil
	}

	// Transaction must be sent to the builder contract address to be considered a bid
	if *ethTx.To() != mempool.builderContract.RegistryKey() {
		return false, nil
	}

	return mempool.validateAuctionTx(ethTx)
}

// GetTransactionSigners defines a function that returns the signers of a
// bundle transaction i.e. transaction that was included in the auction transaction's bundle.
func (mempool *MempoolConfig) GetTransactionSigners(tx []byte) (map[string]struct{}, error) {
	ethTx := &coretypes.Transaction{}
	if err := ethTx.UnmarshalBinary(tx); err != nil {
		return nil, err
	}

	from, err := getFrom(ethTx)
	if err != nil {
		return nil, err
	}

	signer := cosmlib.AddressToAccAddress(from).String()
	signers := map[string]struct{}{
		signer: {},
	}

	return signers, nil
}

// WrapBundleTransaction defines a function that wraps a bundle transaction into a sdk.Tx.
func (mempool *MempoolConfig) WrapBundleTransaction(tx []byte) (sdk.Tx, error) {
	ethTx := &coretypes.Transaction{}
	if err := ethTx.UnmarshalBinary(tx); err != nil {
		return nil, err
	}

	sdkTx, err := mempool.host.SerializeToSdkTx(ethTx)
	if err != nil {
		return nil, err
	}

	return sdkTx, nil
}

// GetBidder defines a function that returns the bidder of an auction transaction transaction.
func (mempool *MempoolConfig) GetBidder(tx sdk.Tx) (sdk.AccAddress, error) {
	isAuctionTx, err := mempool.IsAuctionTx(tx)
	if err != nil {
		return nil, err
	}

	if !isAuctionTx {
		return nil, fmt.Errorf("transaction is not an auction transaction")
	}

	ethTx, err := getEthTransactionRequest(tx)
	if err != nil {
		return nil, err
	}

	if ethTx == nil {
		return nil, fmt.Errorf("transaction is not an auction transaction")
	}

	from, err := getFrom(ethTx)
	if err != nil {
		return nil, err
	}

	bidder := cosmlib.AddressToAccAddress(from)

	return bidder, nil
}

// GetBid defines a function that returns the bid of an auction transaction.
func (mempool *MempoolConfig) GetBid(tx sdk.Tx) (sdk.Coin, error) {
	isAuctionTx, err := mempool.IsAuctionTx(tx)
	if err != nil {
		return sdk.Coin{}, err
	}

	if !isAuctionTx {
		return sdk.Coin{}, fmt.Errorf("transaction is not an auction transaction")
	}

	auctionBidInfo, err := mempool.getBidInfoFromSdkTx(tx)
	if err != nil {
		return sdk.Coin{}, err
	}

	return auctionBidInfo.Bid, nil
}

// GetBundledTransactions defines a function that returns the bundled transactions
// that the user wants to execute at the top of the block given an auction transaction.
func (mempool *MempoolConfig) GetBundledTransactions(tx sdk.Tx) ([][]byte, error) {
	isAuctionTx, err := mempool.IsAuctionTx(tx)
	if err != nil {
		return nil, err
	}

	if !isAuctionTx {
		return nil, fmt.Errorf("transaction is not an auction transaction")
	}

	auctionBidInfo, err := mempool.getBidInfoFromSdkTx(tx)
	if err != nil {
		return nil, err
	}

	return auctionBidInfo.Transactions, nil
}

// HasValidTimeout defines a function that returns true iff the auction transaction
// has a valid timeout.
func (mempool *MempoolConfig) HasValidTimeout(ctx sdk.Context, tx sdk.Tx) error {
	if isAuctionTx, err := mempool.IsAuctionTx(tx); err != nil || !isAuctionTx {
		return err
	}

	auctionBidInfo, err := mempool.getBidInfoFromSdkTx(tx)
	if err != nil {
		return err
	}

	if auctionBidInfo.Timeout < uint64(ctx.BlockHeight()) {
		return fmt.Errorf("auction transaction has an invalid timeout")
	}

	return nil
}

// --------------------------------------------------------------------- //
// ----------------------- Helper Functions ---------------------------- //
// --------------------------------------------------------------------- //

// validateAuctionTx returns true if the ethereum transaction is an auction bid transaction. Since
// we do not have access to valid basic in the mempool, we must valid it here.
func (mempool *MempoolConfig) validateAuctionTx(ethTx *coretypes.Transaction) (bool, error) {
	// The user should not be sending any value to the builder contract
	if ethTx.Value().Cmp(sdk.ZeroInt().BigInt()) != 0 {
		return false, fmt.Errorf("a bid transaction must not send any %s to the builder contract", mempool.evmDenom)
	}

	// The user should be sending a valid transaction to the builder contract's bid function
	bidInfo, err := mempool.getBidInfoFromEthTx(ethTx)
	if err != nil {
		return false, fmt.Errorf("transaction must be a valid bid transaction: %w", err)
	}

	// Since we do not have access to valid basic in the mempool, we must ensure that the
	// bid is valid here
	// if len(bidInfo.Transactions) == 0 {
	// 	return false, fmt.Errorf("bundle of transactions must not be empty")
	// }

	for _, tx := range bidInfo.Transactions {
		if len(tx) == 0 {
			return false, fmt.Errorf("transaction bundle must not contain empty transactions")
		}
	}

	if bidInfo.Bid.Denom != mempool.evmDenom {
		return false, fmt.Errorf("bid must be in %s but got %s", mempool.evmDenom, bidInfo.Bid.Denom)
	}

	return true, nil
}

// getBidInfoFromSdkTx returns the bid amount and the bundle of transactions from a
// Cosmos SDK transaction.
func (mempool *MempoolConfig) getBidInfoFromSdkTx(tx sdk.Tx) (*AuctionBidInfo, error) {
	ethTx, err := getEthTransactionRequest(tx)
	if err != nil {
		return nil, err
	}

	if ethTx == nil {
		return nil, fmt.Errorf("transaction is not an auction transaction")
	}

	return mempool.getBidInfoFromEthTx(ethTx)
}

// getBidInfoFromEthTx returns the bid amount and the bundle of transactions from an
// Ethereum transaction.
func (mempool *MempoolConfig) getBidInfoFromEthTx(ethTx *coretypes.Transaction) (*AuctionBidInfo, error) {
	data := ethTx.Data()
	if len(data) <= 4 {
		return nil, fmt.Errorf("transaction data is too short")
	}

	// Get the method name and the inputs from the transaction data
	methodSigData := data[:4]
	method, err := mempool.contractABI.MethodById(methodSigData)
	if err != nil {
		return nil, err
	}

	// Get the inputs from the transaction data
	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		return nil, err
	}

	bidInfo, ok := inputsMap["bid"].(struct {
		Amount uint64 "json:\"amount\""
		Denom  string "json:\"denom\""
	})
	if !ok {
		return nil, fmt.Errorf("invalid bid type: %T", inputsMap["bid"])
	}

	bundle, ok := inputsMap["transactions"].([][]byte)
	if !ok {
		return nil, fmt.Errorf("invalid bundle type: %T", inputsMap["bundle"])
	}

	timeout, ok := inputsMap["timeout"].(uint64)
	if !ok {
		return nil, fmt.Errorf("invalid timeout type: %T", inputsMap["timeout"])
	}

	auctionBidInfo := &AuctionBidInfo{
		Transactions: bundle,
		Bid:          sdk.NewCoin(bidInfo.Denom, sdk.NewIntFromUint64(bidInfo.Amount)),
		Timeout:      timeout,
	}

	return auctionBidInfo, nil
}

// getFrom returns the sender of an Ethereum transaction.
func getFrom(tx *coretypes.Transaction) (common.Address, error) {
	from, err := gethtypes.Sender(gethtypes.LatestSignerForChainID(tx.ChainId()), tx)
	return from, err
}

// getEthTransactionRequest returns the EthTransactionRequest message from a
// sdk transaction. If the transaction is not an EthTransactionRequest, it returns
// nil.
func getEthTransactionRequest(tx sdk.Tx) (*coretypes.Transaction, error) {
	msgEthTx := make([]*coretypes.Transaction, 0)
	for _, msg := range tx.GetMsgs() {
		if ethTxMsg, ok := msg.(*types.EthTransactionRequest); ok {
			msgEthTx = append(msgEthTx, ethTxMsg.AsTransaction())
		}
	}

	switch {
	case len(msgEthTx) == 0:
		return nil, nil
	case len(msgEthTx) == 1 && len(tx.GetMsgs()) == 1:
		return msgEthTx[0], nil
	default:
		return nil, fmt.Errorf("invalid transaction: %T", tx)
	}
}
