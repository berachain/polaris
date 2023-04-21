package mempool

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/skip-mev/pob/mempool"
	bindings "pkg.berachain.dev/polaris/contracts/bindings/cosmos/precompile"
	cosmlib "pkg.berachain.dev/polaris/cosmos/lib"
	"pkg.berachain.dev/polaris/eth/common"
	coretypes "pkg.berachain.dev/polaris/eth/core/types"
)

// In order for the mempool to differentiate between normal and bid transactions, the application
// must implement the mempool.Config interface.
var _ mempool.Config = (*Config)(nil)

type (
	// Config defines the necessary functionality and fields required to implement
	// the mempool.Config interface. It is used to validate auction transactions, wrap bundle transactions
	// and get the signers of bundle transactions.
	Config struct {
		// builderContract is the address of the builder precompile contract
		builderContract common.Address
		// txDecoder is the transaction decoder used to decode Cosmos SDK transactions
		txDecoder sdk.TxDecoder
		// contractABI is the ABI of the builder contract used to decode auction bids
		contractABI abi.ABI
		// serializer is the serializer used to serialize ethereum transactions to Cosmos SDK transactions
		serializer Serializer
		// evmDenom is the denom of the evm coin which must be used for auction bids
		evmDenom string
	}

	// Serializer defines the necessary functionality to serialize a ethereum transactions to Cosmos SDK transactions.
	Serializer interface {
		SerializeToSdkTx(tx *coretypes.Transaction) (sdk.Tx, error)
	}
)

// NewMempoolConfig returns a new instance of the mempool config.
func NewMempoolConfig(builderContract common.Address, txDecoder sdk.TxDecoder, serializer Serializer, denom string) *Config {
	contractABI, err := abi.JSON(strings.NewReader(bindings.BuilderModuleMetaData.ABI))
	if err != nil {
		panic(err)
	}

	return &Config{
		builderContract: builderContract,
		txDecoder:       txDecoder,
		contractABI:     contractABI,
		serializer:      serializer,
		evmDenom:        denom,
	}
}

// IsAuctionTx defines a function that returns true iff a transaction is an
// auction bid transaction. We define an auction bid transaction to be a transaction
// that
// 1. is a Cosmos SDK transaction that contains a single Ethereum transaction request
// 2. the Ethereum transaction request is sent to the builder contract address
// 3. the Ethereum transaction request is a valid auction bid transaction
func (txConfig *Config) IsAuctionTx(tx sdk.Tx) (bool, error) {
	// Ensure the transcaction is an EthTransactionRequest
	ethTx, err := getEthTransactionRequest(tx)
	if err != nil || ethTx == nil {
		return false, nil
	}

	// Transaction must be sent to the builder contract address to be considered a bid
	if *ethTx.To() != txConfig.builderContract {
		return false, nil
	}

	return txConfig.validateAuctionTx(ethTx)
}

// GetTransactionSigners defines a function that returns the signers of a transaction that
// is included in a searchers bundle. In this case, each transaction in the bundle is a
// core ethereum transaction type as bytes.
func (txConfig *Config) GetTransactionSigners(tx []byte) (map[string]struct{}, error) {
	ethTx := &coretypes.Transaction{}
	if err := ethTx.UnmarshalBinary(tx); err != nil {
		return nil, err
	}

	from, err := getFromEthTx(ethTx)
	if err != nil {
		return nil, err
	}

	signer := cosmlib.AddressToAccAddress(from).String()
	signers := map[string]struct{}{
		signer: {},
	}

	return signers, nil
}

// WrapBundleTransaction defines a function that wraps a bundle transaction (eth core transaction type) into a sdk.Tx.
func (txConfig *Config) WrapBundleTransaction(tx []byte) (sdk.Tx, error) {
	ethTx := &coretypes.Transaction{}
	if err := ethTx.UnmarshalBinary(tx); err != nil {
		return nil, err
	}

	sdkTx, err := txConfig.serializer.SerializeToSdkTx(ethTx)
	if err != nil {
		return nil, err
	}

	return sdkTx, nil
}

// GetBidder defines a function that returns the bidder of an auction bid transaction.
func (txConfig *Config) GetBidder(tx sdk.Tx) (sdk.AccAddress, error) {
	isAuctionTx, err := txConfig.IsAuctionTx(tx)
	if err != nil {
		return nil, err
	}

	if !isAuctionTx {
		return nil, fmt.Errorf("transaction is not an auction transaction")
	}

	auctionBidInfo, err := txConfig.getBidInfoFromSdkTx(tx)
	if err != nil {
		return nil, err
	}

	return auctionBidInfo.Bidder, nil
}

// GetBid defines a function that returns the bid of an auction transaction.
func (txConfig *Config) GetBid(tx sdk.Tx) (sdk.Coin, error) {
	isAuctionTx, err := txConfig.IsAuctionTx(tx)
	if err != nil {
		return sdk.Coin{}, err
	}

	if !isAuctionTx {
		return sdk.Coin{}, fmt.Errorf("transaction is not an auction transaction")
	}

	auctionBidInfo, err := txConfig.getBidInfoFromSdkTx(tx)
	if err != nil {
		return sdk.Coin{}, err
	}

	return auctionBidInfo.Bid, nil
}

// GetBundledTransactions defines a function that returns the bundled transactions
// that the user wants to execute at the top of the block given an auction transaction.
func (txConfig *Config) GetBundledTransactions(tx sdk.Tx) ([][]byte, error) {
	isAuctionTx, err := txConfig.IsAuctionTx(tx)
	if err != nil {
		return nil, err
	}

	if !isAuctionTx {
		return nil, fmt.Errorf("transaction is not an auction transaction")
	}

	auctionBidInfo, err := txConfig.getBidInfoFromSdkTx(tx)
	if err != nil {
		return nil, err
	}

	return auctionBidInfo.Transactions, nil
}

// GetTimeoutHeight defines a function that returns the timeout height of an auction transaction.
func (txConfig *Config) GetTimeout(tx sdk.Tx) (uint64, error) {
	isAuctionTx, err := txConfig.IsAuctionTx(tx)
	if err != nil {
		return 0, err
	}

	if !isAuctionTx {
		return 0, fmt.Errorf("transaction is not an auction transaction")
	}

	auctionBidInfo, err := txConfig.getBidInfoFromSdkTx(tx)
	if err != nil {
		return 0, err
	}

	return auctionBidInfo.Timeout, nil
}

func (txConfig *Config) GetTimeoutHeight(tx sdk.Tx) (uint64, error) {
	return txConfig.GetTimeout(tx)
}

// GetAuctionBidInfo defines a function that returns the auction bid info of an auction transaction.
func (txConfig *Config) GetAuctionBidInfo(tx sdk.Tx) (mempool.AuctionBidInfo, error) {
	bid, err := txConfig.GetBid(tx)
	if err != nil {
		return mempool.AuctionBidInfo{}, err
	}

	bidder, err := txConfig.GetBidder(tx)
	if err != nil {
		return mempool.AuctionBidInfo{}, err
	}

	bundle, err := txConfig.GetBundledTransactions(tx)
	if err != nil {
		return mempool.AuctionBidInfo{}, err
	}

	timeout, err := txConfig.GetTimeout(tx)
	if err != nil {
		return mempool.AuctionBidInfo{}, err
	}

	return mempool.AuctionBidInfo{
		Bid:          bid,
		Bidder:       bidder,
		Transactions: bundle,
		Timeout:      timeout,
	}, nil
}

// GetBundleSigners defines a function that returns the signers of each transaction in a bundle.
func (txConfig *Config) GetBundleSigners(txs [][]byte) ([]map[string]struct{}, error) {
	signers := make([]map[string]struct{}, len(txs))

	for index, tx := range txs {
		txSigners, err := txConfig.GetTransactionSigners(tx)
		if err != nil {
			return nil, err
		}

		signers[index] = txSigners
	}

	return signers, nil
}
