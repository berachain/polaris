package api

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/ethapi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/sha3"
	"pkg.berachain.dev/polaris/eth/common"
	"pkg.berachain.dev/polaris/eth/common/hexutil"
	"pkg.berachain.dev/polaris/eth/core/types"
	"pkg.berachain.dev/polaris/eth/core/vm"
	"pkg.berachain.dev/polaris/eth/params"
)

// FlashBeraBackend is the backend implementation for the FlashBera API.
type FlashBeraBackend interface {
	SendBundle(ctx context.Context, txs types.Transactions, blockNumber rpc.BlockNumber, uuid uuid.UUID, signingAddress common.Address, minTimestamp uint64, maxTimestamp uint64, revertingTxHashes []common.Hash) error
	ChainConfig() *params.ChainConfig
	HeaderByNumber(ctx context.Context, number rpc.BlockNumber) (*types.Header, error)
	StateAndHeaderByNumberOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash) (state.StateDBI, *types.Header, error)
	GetEVM(context.Context, *core.Message, vm.GethStateDB,
		*types.Header, *vm.Config) (*vm.GethEVM, func() error, error)
	RPCGasCap() uint64
}

// --------------------------------------- FlashBots ------------------------------------------- //
type PrivateTxBundleAPI struct {
	b FlashBeraBackend
}

// NewPrivateTxBundleAPI creates a new Tx Bundle API instance.
func NewPrivateTxBundleAPI(b FlashBeraBackend) *PrivateTxBundleAPI {
	return &PrivateTxBundleAPI{b}
}

// SendBundleArgs represents the arguments for a SendBundle call.
type SendBundleArgs struct {
	Txs               []hexutil.Bytes `json:"txs"`
	BlockNumber       rpc.BlockNumber `json:"blockNumber"`
	ReplacementUuid   *uuid.UUID      `json:"replacementUuid"`
	SigningAddress    *common.Address `json:"signingAddress"`
	MinTimestamp      *uint64         `json:"minTimestamp"`
	MaxTimestamp      *uint64         `json:"maxTimestamp"`
	RevertingTxHashes []common.Hash   `json:"revertingTxHashes"`
}

// SendBundle will add the signed transaction to the transaction pool.
// The sender is responsible for signing the transaction and using the correct nonce and ensuring validity
func (s *PrivateTxBundleAPI) SendBundle(ctx context.Context, args SendBundleArgs) error {
	var txs types.Transactions
	if len(args.Txs) == 0 {
		return errors.New("bundle missing txs")
	}
	if args.BlockNumber == 0 {
		return errors.New("bundle missing blockNumber")
	}

	for _, encodedTx := range args.Txs {
		tx := new(types.Transaction)
		if err := tx.UnmarshalBinary(encodedTx); err != nil {
			return err
		}
		txs = append(txs, tx)
	}

	var replacementUuid uuid.UUID
	if args.ReplacementUuid != nil {
		replacementUuid = *args.ReplacementUuid
	}

	var signingAddress common.Address
	if args.SigningAddress != nil {
		signingAddress = *args.SigningAddress
	}

	var minTimestamp, maxTimestamp uint64
	if args.MinTimestamp != nil {
		minTimestamp = *args.MinTimestamp
	}
	if args.MaxTimestamp != nil {
		maxTimestamp = *args.MaxTimestamp
	}

	go s.b.SendBundle(ctx, txs, args.BlockNumber, replacementUuid, signingAddress, minTimestamp, maxTimestamp, args.RevertingTxHashes)

	return nil
}

// --------------------------------------- Bundle ------------------------------------------- //

// BundleAPI offers an API for accepting bundled transactions. The BundleAPI has been heavily
// inspired by the original mev-geth implementation.
// (https://github.com/flashbots/mev-geth/blob/master/internal/ethapi/api.go#L2038)
type BundleAPI struct {
	b     FlashBeraBackend
	chain core.ChainContext
}

// NewBundleAPI creates a new Tx Bundle API instance.
func NewBundleAPI(b FlashBeraBackend, chain core.ChainContext) *BundleAPI {
	return &BundleAPI{b, chain}
}

// CallBundleArgs represents the arguinterface ments for a call.
type CallBundleArgs struct {
	Txs                    []hexutil.Bytes       `json:"txs"`
	BlockNumber            rpc.BlockNumber       `json:"blockNumber"`
	StateBlockNumberOrHash rpc.BlockNumberOrHash `json:"stateBlockNumber"`
	Coinbase               *string               `json:"coinbase"`
	Timestamp              *uint64               `json:"timestamp"`
	Timeout                *int64                `json:"timeout"`
	GasLimit               *uint64               `json:"gasLimit"`
	Difficulty             *big.Int              `json:"difficulty"`
	BaseFee                *big.Int              `json:"baseFee"`
}

// CallBundle will simulate a bundle of transactions at the top of a given block
// number with the state of another (or the same) block. This can be used to
// simulate future blocks with the current state, or it can be used to simulate
// a past block.
// The sender is responsible for signing the transactions and using the correct
// nonce and ensuring validity
func (s *BundleAPI) CallBundle(ctx context.Context, args CallBundleArgs) (map[string]interface{}, error) {
	if len(args.Txs) == 0 {
		return nil, errors.New("bundle missing txs")
	}
	if args.BlockNumber == 0 {
		return nil, errors.New("bundle missing blockNumber")
	}

	var txs types.Transactions
	for _, encodedTx := range args.Txs {
		tx := new(types.Transaction)
		if err := tx.UnmarshalBinary(encodedTx); err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	defer func(start time.Time) { log.Debug("Executing EVM call finished", "runtime", time.Since(start)) }(time.Now())

	timeoutMilliSeconds := int64(5000)
	if args.Timeout != nil {
		timeoutMilliSeconds = *args.Timeout
	}
	timeout := time.Millisecond * time.Duration(timeoutMilliSeconds)
	state, parent, err := s.b.StateAndHeaderByNumberOrHash(ctx, args.StateBlockNumberOrHash)
	if state == nil || err != nil {
		return nil, err
	}
	blockNumber := big.NewInt(int64(args.BlockNumber))

	timestamp := parent.Time + 1
	if args.Timestamp != nil {
		timestamp = *args.Timestamp
	}
	coinbase := parent.Coinbase
	if args.Coinbase != nil {
		coinbase = common.HexToAddress(*args.Coinbase)
	}
	difficulty := parent.Difficulty
	if args.Difficulty != nil {
		difficulty = args.Difficulty
	}
	gasLimit := parent.GasLimit
	if args.GasLimit != nil {
		gasLimit = *args.GasLimit
	}
	var baseFee *big.Int
	if args.BaseFee != nil {
		baseFee = args.BaseFee
	} else if s.b.ChainConfig().IsLondon(big.NewInt(args.BlockNumber.Int64())) {
		baseFee = misc.CalcBaseFee(s.b.ChainConfig(), parent)
	}
	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     blockNumber,
		GasLimit:   gasLimit,
		Time:       timestamp,
		Difficulty: difficulty,
		Coinbase:   coinbase,
		BaseFee:    baseFee,
	}

	// Setup context so it may be cancelled the call has completed
	// or, in case of unmetered gas, setup a context with a timeout.
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}
	// Make sure the context is cancelled when the call has completed
	// this makes sure resources are cleaned up.
	defer cancel()

	// Setup the gas pool (also for unmetered requests)
	// and apply the message.
	gp := new(core.GasPool).AddGas(math.MaxUint64)

	results := []map[string]interface{}{}
	coinbaseBalanceBefore := state.GetBalance(coinbase)

	bundleHash := sha3.NewLegacyKeccak256()
	signer := types.MakeSigner(s.b.ChainConfig(), blockNumber)
	var totalGasUsed uint64
	gasFees := new(big.Int)

	msg, err := core.TransactionToMessage(txs[0], signer, header.BaseFee)
	if err != nil {
		return nil, err
	}

	vmenv, vmError, err := s.b.GetEVM(ctx, msg, state, header, &vm.Config{})
	if err != nil {
		return nil, err
	}

	for i, tx := range txs {
		coinbaseBalanceBeforeTx := state.GetBalance(coinbase)
		state.SetTxContext(tx.Hash(), i)

		receipt, result, err := core.ApplyTransactionWithEVMWithResult(vmenv, s.b.ChainConfig(), &coinbase, gp, state, header, tx, &header.GasUsed)
		if err != nil {
			return nil, fmt.Errorf("err: %w; txhash %s", err, tx.Hash())
		}
		if err := vmError(); err != nil {
			return nil, fmt.Errorf("execution error: %v", err)
		}

		txHash := tx.Hash().String()
		from, err := types.Sender(signer, tx)
		if err != nil {
			return nil, fmt.Errorf("err: %w; txhash %s", err, tx.Hash())
		}
		to := "0x"
		if tx.To() != nil {
			to = tx.To().String()
		}
		jsonResult := map[string]interface{}{
			"txHash":      txHash,
			"gasUsed":     receipt.GasUsed,
			"fromAddress": from.String(),
			"toAddress":   to,
		}
		totalGasUsed += receipt.GasUsed
		gasPrice, err := tx.EffectiveGasTip(header.BaseFee)
		if err != nil {
			return nil, fmt.Errorf("err: %w; txhash %s", err, tx.Hash())
		}
		gasFeesTx := new(big.Int).Mul(big.NewInt(int64(receipt.GasUsed)), gasPrice)
		gasFees.Add(gasFees, gasFeesTx)
		bundleHash.Write(tx.Hash().Bytes())
		if result.Err != nil {
			jsonResult["error"] = result.Err.Error()
			revert := result.Revert()
			if len(revert) > 0 {
				jsonResult["revert"] = string(revert)
			}
		} else {
			dst := make([]byte, hex.EncodedLen(len(result.Return())))
			hex.Encode(dst, result.Return())
			jsonResult["value"] = "0x" + string(dst)
		}
		coinbaseDiffTx := new(big.Int).Sub(state.GetBalance(coinbase), coinbaseBalanceBeforeTx)
		jsonResult["coinbaseDiff"] = coinbaseDiffTx.String()
		jsonResult["gasFees"] = gasFeesTx.String()
		jsonResult["ethSentToCoinbase"] = new(big.Int).Sub(coinbaseDiffTx, gasFeesTx).String()
		jsonResult["gasPrice"] = new(big.Int).Div(coinbaseDiffTx, big.NewInt(int64(receipt.GasUsed))).String()
		jsonResult["gasUsed"] = receipt.GasUsed
		results = append(results, jsonResult)
	}

	ret := map[string]interface{}{}
	ret["results"] = results
	coinbaseDiff := new(big.Int).Sub(state.GetBalance(coinbase), coinbaseBalanceBefore)
	ret["coinbaseDiff"] = coinbaseDiff.String()
	ret["gasFees"] = gasFees.String()
	ret["ethSentToCoinbase"] = new(big.Int).Sub(coinbaseDiff, gasFees).String()
	ret["bundleGasPrice"] = new(big.Int).Div(coinbaseDiff, big.NewInt(int64(totalGasUsed))).String()
	ret["totalGasUsed"] = totalGasUsed
	ret["stateBlockNumber"] = parent.Number.Int64()

	ret["bundleHash"] = "0x" + common.Bytes2Hex(bundleHash.Sum(nil))
	return ret, nil
}

// EstimateGasBundleArgs represents the arguments for a call
type EstimateGasBundleArgs struct {
	Txs                    []ethapi.TransactionArgs `json:"txs"`
	BlockNumber            rpc.BlockNumber          `json:"blockNumber"`
	StateBlockNumberOrHash rpc.BlockNumberOrHash    `json:"stateBlockNumber"`
	Coinbase               *string                  `json:"coinbase"`
	Timestamp              *uint64                  `json:"timestamp"`
	Timeout                *int64                   `json:"timeout"`
}

func (s *BundleAPI) EstimateGasBundle(ctx context.Context, args EstimateGasBundleArgs) (map[string]interface{}, error) {
	if len(args.Txs) == 0 {
		return nil, errors.New("bundle missing txs")
	}
	if args.BlockNumber == 0 {
		return nil, errors.New("bundle missing blockNumber")
	}

	timeoutMS := int64(5000)
	if args.Timeout != nil {
		timeoutMS = *args.Timeout
	}
	timeout := time.Millisecond * time.Duration(timeoutMS)

	state, parent, err := s.b.StateAndHeaderByNumberOrHash(ctx, args.StateBlockNumberOrHash)
	if state == nil || err != nil {
		return nil, err
	}
	blockNumber := big.NewInt(int64(args.BlockNumber))
	timestamp := parent.Time + 1
	if args.Timestamp != nil {
		timestamp = *args.Timestamp
	}
	coinbase := parent.Coinbase
	if args.Coinbase != nil {
		coinbase = common.HexToAddress(*args.Coinbase)
	}

	header := &types.Header{
		ParentHash: parent.Hash(),
		Number:     blockNumber,
		GasLimit:   parent.GasLimit,
		Time:       timestamp,
		Difficulty: parent.Difficulty,
		Coinbase:   coinbase,
		BaseFee:    parent.BaseFee,
	}

	// Setup context so it may be cancelled when the call
	// has completed or, in case of unmetered gas, setup
	// a context with a timeout
	var cancel context.CancelFunc
	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	// Make sure the context is cancelled when the call has completed
	// This makes sure resources are cleaned up
	defer cancel()

	// RPC Call gas cap
	globalGasCap := s.b.RPCGasCap()

	// Results
	results := []map[string]interface{}{}

	// Gas pool
	gp := new(core.GasPool).AddGas(math.MaxUint64)

	// Feed each of the transactions into the VM ctx
	// And try and estimate the gas used
	for i, txArgs := range args.Txs {
		// Since its a txCall we'll just prepare the
		// state with a random hash
		var randomHash common.Hash
		rand.Read(randomHash[:])

		// New random hash since its a call
		state.SetTxContext(randomHash, i)

		// Convert tx args to msg to apply state transition
		msg, err := txArgs.ToMessage(globalGasCap, header.BaseFee)
		if err != nil {
			return nil, err
		}

		// Get EVM Environment
		vmenv, vmError, err := s.b.GetEVM(ctx, msg, state, header, &vm.Config{NoBaseFee: true})
		if err != nil {
			return nil, err
		}

		// Apply state transition
		result, err := core.ApplyMessage(vmenv, msg, gp)
		if err != nil {
			return nil, err
		}

		// Check for the vm error
		if err := vmError(); err != nil {
			return nil, fmt.Errorf("execution error: %v", err)
		}

		// Modifications are committed to the state
		// Only delete empty objects if EIP158/161 (a.k.a Spurious Dragon) is in effect
		state.Finalise(vmenv.ChainConfig().IsEIP158(blockNumber))

		// Append result
		jsonResult := map[string]interface{}{
			"gasUsed": result.UsedGas,
		}
		results = append(results, jsonResult)
	}

	// Return results
	ret := map[string]interface{}{}
	ret["results"] = results

	return ret, nil
}
