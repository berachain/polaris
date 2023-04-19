package config

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/skip-mev/pob/mempool"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/cosmos/x/evm/plugins/txpool"
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
		serializer      txpool.Serializer
		cp              txpool.ConfigurationPlugin
	}

	AuctionBidInfo struct {
		Transactions [][]byte
		Bid          sdk.Coin
	}
)

// NewMempoolConfig returns a new instance of the mempool config.
func NewMempoolConfig(builderContract precompile.StatefulImpl, txDecoder sdk.TxDecoder, cp txpool.ConfigurationPlugin, clientCtx client.Context) *MempoolConfig {
	contractABI, err := abi.JSON(strings.NewReader(bindings.BuilderModuleMetaData.ABI))
	if err != nil {
		panic(err)
	}

	return &MempoolConfig{
		builderContract: builderContract,
		txDecoder:       txDecoder,
		contractABI:     contractABI,
		cp:              cp,
		serializer:      txpool.NewSerializer(cp, clientCtx),
	}
}

// IsAuctionTx defines a function that returns true iff a transaction is an
// auction bid transaction. An transaction is an auction bid transaction if it
// 1. is an EthTransactionRequest
// 2. the transaction's data is a call to the builder contract's auctionBid method
// 3. the transaction's data provides valid arguments to the auctionBid method
func (mempool *MempoolConfig) IsAuctionTx(tx sdk.Tx) (bool, error) {
	// Ensure the transcaction is an EthTransactionRequest
	ethTx, err := getEthTransactionRequest(tx)
	if err != nil {
		return false, err
	}

	if ethTx == nil {
		return false, nil
	}

	return mempool.validateAuctionTx(ethTx)
}

// GetTransactionSigners defines a function that returns the signers of a
// bundle transaction i.e. transaction that was included in the auction transaction's bundle.
func (mempool *MempoolConfig) GetTransactionSigners(tx []byte) (map[string]bool, error) {
	return nil, nil
}

// WrapBundleTransaction defines a function that wraps a bundle transaction into a sdk.Tx.
func (mempool *MempoolConfig) WrapBundleTransaction(tx []byte) (sdk.Tx, error) {
	return nil, nil
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

	ethTx, err := getEthTransactionRequest(tx)
	if err != nil {
		return sdk.Coin{}, err
	}

	if ethTx == nil {
		return sdk.Coin{}, fmt.Errorf("transaction is not an auction transaction")
	}

	bidInfo, err := mempool.getBidInfoFromEthTx(ethTx)
	if err != nil {
		return sdk.Coin{}, err
	}

	return bidInfo.Bid, nil
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

	ethTx, err := getEthTransactionRequest(tx)
	if err != nil {
		return nil, err
	}

	if ethTx == nil {
		return nil, fmt.Errorf("transaction is not an auction transaction")
	}

	bidInfo, err := mempool.getBidInfoFromEthTx(ethTx)
	if err != nil {
		return nil, err
	}

	return bidInfo.Transactions, nil
}

// --------------------------------------------------------------------- //
// ----------------------- Helper Functions ---------------------------- //
// --------------------------------------------------------------------- //

func (mempool *MempoolConfig) validateAuctionTx(ethTx *coretypes.Transaction) (bool, error) {
	if *ethTx.To() != mempool.builderContract.RegistryKey() {
		return false, fmt.Errorf("transaction must be sent to the builder contract")
	}

	// The user should not be sending any value to the builder contract
	if ethTx.Value().Cmp(sdk.ZeroInt().BigInt()) != 0 {
		return false, fmt.Errorf("transaction must not send any value to the builder contract")
	}

	// The user should be sending a valid transaction to the builder contract's bid function
	bidInfo, err := mempool.getBidInfoFromEthTx(ethTx)
	if err != nil {
		return false, fmt.Errorf("transaction must be a valid bid transaction: %w", err)
	}

	// Since we do not have access to valid basic in the mempool, we must ensure that the
	// bid is valid here
	if len(bidInfo.Transactions) == 0 {
		return false, fmt.Errorf("transaction bundle must not be empty")
	}

	for _, tx := range bidInfo.Transactions {
		if len(tx) == 0 {
			return false, fmt.Errorf("transaction bundle must not contain empty transactions")
		}
	}

	if bidInfo.Bid.Denom != mempool.cp.GetEvmDenom() {
		return false, fmt.Errorf("bid must be in %s", mempool.cp.GetEvmDenom())
	}

	return true, nil
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

	auctionBidInfo := &AuctionBidInfo{
		Transactions: bundle,
		Bid:          sdk.NewCoin(bidInfo.Denom, sdk.NewIntFromUint64(bidInfo.Amount)),
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
