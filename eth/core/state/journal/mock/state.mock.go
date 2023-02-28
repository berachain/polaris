// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"pkg.berachain.dev/stargazer/eth/core/state"
	"sync"
)

// Ensure, that PluginMock does implement state.Plugin.
// If this is not the case, regenerate this file with moq.
var _ state.Plugin = &PluginMock{}

// PluginMock is a mock implementation of state.Plugin.
//
//	func TestSomethingThatUsesPlugin(t *testing.T) {
//
//		// make and configure a mocked state.Plugin
//		mockedPlugin := &PluginMock{
//			AddBalanceFunc: func(address common.Address, intMoqParam *big.Int)  {
//				panic("mock out the AddBalance method")
//			},
//			CreateAccountFunc: func(address common.Address)  {
//				panic("mock out the CreateAccount method")
//			},
//			DeleteSuicidesFunc: func(addresss []common.Address)  {
//				panic("mock out the DeleteSuicides method")
//			},
//			ExistFunc: func(address common.Address) bool {
//				panic("mock out the Exist method")
//			},
//			FinalizeFunc: func()  {
//				panic("mock out the Finalize method")
//			},
//			ForEachStorageFunc: func(address common.Address, fn func(common.Hash, common.Hash) bool) error {
//				panic("mock out the ForEachStorage method")
//			},
//			GetBalanceFunc: func(address common.Address) *big.Int {
//				panic("mock out the GetBalance method")
//			},
//			GetCodeFunc: func(address common.Address) []byte {
//				panic("mock out the GetCode method")
//			},
//			GetCodeHashFunc: func(address common.Address) common.Hash {
//				panic("mock out the GetCodeHash method")
//			},
//			GetCodeSizeFunc: func(address common.Address) int {
//				panic("mock out the GetCodeSize method")
//			},
//			GetCommittedStateFunc: func(address common.Address, hash common.Hash) common.Hash {
//				panic("mock out the GetCommittedState method")
//			},
//			GetNonceFunc: func(address common.Address) uint64 {
//				panic("mock out the GetNonce method")
//			},
//			GetStateFunc: func(address common.Address, hash common.Hash) common.Hash {
//				panic("mock out the GetState method")
//			},
//			RegistryKeyFunc: func() string {
//				panic("mock out the RegistryKey method")
//			},
//			ResetFunc: func(contextMoqParam context.Context)  {
//				panic("mock out the Reset method")
//			},
//			RevertToSnapshotFunc: func(n int)  {
//				panic("mock out the RevertToSnapshot method")
//			},
//			SetBalanceFunc: func(address common.Address, intMoqParam *big.Int)  {
//				panic("mock out the SetBalance method")
//			},
//			SetCodeFunc: func(address common.Address, bytes []byte)  {
//				panic("mock out the SetCode method")
//			},
//			SetNonceFunc: func(address common.Address, v uint64)  {
//				panic("mock out the SetNonce method")
//			},
//			SetStateFunc: func(address common.Address, hash1 common.Hash, hash2 common.Hash)  {
//				panic("mock out the SetState method")
//			},
//			SetStorageFunc: func(addr common.Address, storage map[common.Hash]common.Hash)  {
//				panic("mock out the SetStorage method")
//			},
//			SnapshotFunc: func() int {
//				panic("mock out the Snapshot method")
//			},
//			SubBalanceFunc: func(address common.Address, intMoqParam *big.Int)  {
//				panic("mock out the SubBalance method")
//			},
//			TransferBalanceFunc: func(address1 common.Address, address2 common.Address, intMoqParam *big.Int)  {
//				panic("mock out the TransferBalance method")
//			},
//		}
//
//		// use mockedPlugin in code that requires state.Plugin
//		// and then make assertions.
//
//	}
type PluginMock struct {
	// AddBalanceFunc mocks the AddBalance method.
	AddBalanceFunc func(address common.Address, intMoqParam *big.Int)

	// CreateAccountFunc mocks the CreateAccount method.
	CreateAccountFunc func(address common.Address)

	// DeleteSuicidesFunc mocks the DeleteSuicides method.
	DeleteSuicidesFunc func(addresss []common.Address)

	// ExistFunc mocks the Exist method.
	ExistFunc func(address common.Address) bool

	// FinalizeFunc mocks the Finalize method.
	FinalizeFunc func()

	// ForEachStorageFunc mocks the ForEachStorage method.
	ForEachStorageFunc func(address common.Address, fn func(common.Hash, common.Hash) bool) error

	// GetBalanceFunc mocks the GetBalance method.
	GetBalanceFunc func(address common.Address) *big.Int

	// GetCodeFunc mocks the GetCode method.
	GetCodeFunc func(address common.Address) []byte

	// GetCodeHashFunc mocks the GetCodeHash method.
	GetCodeHashFunc func(address common.Address) common.Hash

	// GetCodeSizeFunc mocks the GetCodeSize method.
	GetCodeSizeFunc func(address common.Address) int

	// GetCommittedStateFunc mocks the GetCommittedState method.
	GetCommittedStateFunc func(address common.Address, hash common.Hash) common.Hash

	// GetNonceFunc mocks the GetNonce method.
	GetNonceFunc func(address common.Address) uint64

	// GetStateFunc mocks the GetState method.
	GetStateFunc func(address common.Address, hash common.Hash) common.Hash

	// RegistryKeyFunc mocks the RegistryKey method.
	RegistryKeyFunc func() string

	// ResetFunc mocks the Reset method.
	ResetFunc func(contextMoqParam context.Context)

	// RevertToSnapshotFunc mocks the RevertToSnapshot method.
	RevertToSnapshotFunc func(n int)

	// SetBalanceFunc mocks the SetBalance method.
	SetBalanceFunc func(address common.Address, intMoqParam *big.Int)

	// SetCodeFunc mocks the SetCode method.
	SetCodeFunc func(address common.Address, bytes []byte)

	// SetNonceFunc mocks the SetNonce method.
	SetNonceFunc func(address common.Address, v uint64)

	// SetStateFunc mocks the SetState method.
	SetStateFunc func(address common.Address, hash1 common.Hash, hash2 common.Hash)

	// SetStorageFunc mocks the SetStorage method.
	SetStorageFunc func(addr common.Address, storage map[common.Hash]common.Hash)

	// SnapshotFunc mocks the Snapshot method.
	SnapshotFunc func() int

	// SubBalanceFunc mocks the SubBalance method.
	SubBalanceFunc func(address common.Address, intMoqParam *big.Int)

	// TransferBalanceFunc mocks the TransferBalance method.
	TransferBalanceFunc func(address1 common.Address, address2 common.Address, intMoqParam *big.Int)

	// calls tracks calls to the methods.
	calls struct {
		// AddBalance holds details about calls to the AddBalance method.
		AddBalance []struct {
			// Address is the address argument value.
			Address common.Address
			// IntMoqParam is the intMoqParam argument value.
			IntMoqParam *big.Int
		}
		// CreateAccount holds details about calls to the CreateAccount method.
		CreateAccount []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// DeleteSuicides holds details about calls to the DeleteSuicides method.
		DeleteSuicides []struct {
			// Addresss is the addresss argument value.
			Addresss []common.Address
		}
		// Exist holds details about calls to the Exist method.
		Exist []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// Finalize holds details about calls to the Finalize method.
		Finalize []struct {
		}
		// ForEachStorage holds details about calls to the ForEachStorage method.
		ForEachStorage []struct {
			// Address is the address argument value.
			Address common.Address
			// Fn is the fn argument value.
			Fn func(common.Hash, common.Hash) bool
		}
		// GetBalance holds details about calls to the GetBalance method.
		GetBalance []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// GetCode holds details about calls to the GetCode method.
		GetCode []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// GetCodeHash holds details about calls to the GetCodeHash method.
		GetCodeHash []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// GetCodeSize holds details about calls to the GetCodeSize method.
		GetCodeSize []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// GetCommittedState holds details about calls to the GetCommittedState method.
		GetCommittedState []struct {
			// Address is the address argument value.
			Address common.Address
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// GetNonce holds details about calls to the GetNonce method.
		GetNonce []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// GetState holds details about calls to the GetState method.
		GetState []struct {
			// Address is the address argument value.
			Address common.Address
			// Hash is the hash argument value.
			Hash common.Hash
		}
		// RegistryKey holds details about calls to the RegistryKey method.
		RegistryKey []struct {
		}
		// Reset holds details about calls to the Reset method.
		Reset []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
		}
		// RevertToSnapshot holds details about calls to the RevertToSnapshot method.
		RevertToSnapshot []struct {
			// N is the n argument value.
			N int
		}
		// SetBalance holds details about calls to the SetBalance method.
		SetBalance []struct {
			// Address is the address argument value.
			Address common.Address
			// IntMoqParam is the intMoqParam argument value.
			IntMoqParam *big.Int
		}
		// SetCode holds details about calls to the SetCode method.
		SetCode []struct {
			// Address is the address argument value.
			Address common.Address
			// Bytes is the bytes argument value.
			Bytes []byte
		}
		// SetNonce holds details about calls to the SetNonce method.
		SetNonce []struct {
			// Address is the address argument value.
			Address common.Address
			// V is the v argument value.
			V uint64
		}
		// SetState holds details about calls to the SetState method.
		SetState []struct {
			// Address is the address argument value.
			Address common.Address
			// Hash1 is the hash1 argument value.
			Hash1 common.Hash
			// Hash2 is the hash2 argument value.
			Hash2 common.Hash
		}
		// SetStorage holds details about calls to the SetStorage method.
		SetStorage []struct {
			// Addr is the addr argument value.
			Addr common.Address
			// Storage is the storage argument value.
			Storage map[common.Hash]common.Hash
		}
		// Snapshot holds details about calls to the Snapshot method.
		Snapshot []struct {
		}
		// SubBalance holds details about calls to the SubBalance method.
		SubBalance []struct {
			// Address is the address argument value.
			Address common.Address
			// IntMoqParam is the intMoqParam argument value.
			IntMoqParam *big.Int
		}
		// TransferBalance holds details about calls to the TransferBalance method.
		TransferBalance []struct {
			// Address1 is the address1 argument value.
			Address1 common.Address
			// Address2 is the address2 argument value.
			Address2 common.Address
			// IntMoqParam is the intMoqParam argument value.
			IntMoqParam *big.Int
		}
	}
	lockAddBalance        sync.RWMutex
	lockCreateAccount     sync.RWMutex
	lockDeleteSuicides    sync.RWMutex
	lockExist             sync.RWMutex
	lockFinalize          sync.RWMutex
	lockForEachStorage    sync.RWMutex
	lockGetBalance        sync.RWMutex
	lockGetCode           sync.RWMutex
	lockGetCodeHash       sync.RWMutex
	lockGetCodeSize       sync.RWMutex
	lockGetCommittedState sync.RWMutex
	lockGetNonce          sync.RWMutex
	lockGetState          sync.RWMutex
	lockRegistryKey       sync.RWMutex
	lockReset             sync.RWMutex
	lockRevertToSnapshot  sync.RWMutex
	lockSetBalance        sync.RWMutex
	lockSetCode           sync.RWMutex
	lockSetNonce          sync.RWMutex
	lockSetState          sync.RWMutex
	lockSetStorage        sync.RWMutex
	lockSnapshot          sync.RWMutex
	lockSubBalance        sync.RWMutex
	lockTransferBalance   sync.RWMutex
}

// AddBalance calls AddBalanceFunc.
func (mock *PluginMock) AddBalance(address common.Address, intMoqParam *big.Int) {
	if mock.AddBalanceFunc == nil {
		panic("PluginMock.AddBalanceFunc: method is nil but Plugin.AddBalance was just called")
	}
	callInfo := struct {
		Address     common.Address
		IntMoqParam *big.Int
	}{
		Address:     address,
		IntMoqParam: intMoqParam,
	}
	mock.lockAddBalance.Lock()
	mock.calls.AddBalance = append(mock.calls.AddBalance, callInfo)
	mock.lockAddBalance.Unlock()
	mock.AddBalanceFunc(address, intMoqParam)
}

// AddBalanceCalls gets all the calls that were made to AddBalance.
// Check the length with:
//
//	len(mockedPlugin.AddBalanceCalls())
func (mock *PluginMock) AddBalanceCalls() []struct {
	Address     common.Address
	IntMoqParam *big.Int
} {
	var calls []struct {
		Address     common.Address
		IntMoqParam *big.Int
	}
	mock.lockAddBalance.RLock()
	calls = mock.calls.AddBalance
	mock.lockAddBalance.RUnlock()
	return calls
}

// CreateAccount calls CreateAccountFunc.
func (mock *PluginMock) CreateAccount(address common.Address) {
	if mock.CreateAccountFunc == nil {
		panic("PluginMock.CreateAccountFunc: method is nil but Plugin.CreateAccount was just called")
	}
	callInfo := struct {
		Address common.Address
	}{
		Address: address,
	}
	mock.lockCreateAccount.Lock()
	mock.calls.CreateAccount = append(mock.calls.CreateAccount, callInfo)
	mock.lockCreateAccount.Unlock()
	mock.CreateAccountFunc(address)
}

// CreateAccountCalls gets all the calls that were made to CreateAccount.
// Check the length with:
//
//	len(mockedPlugin.CreateAccountCalls())
func (mock *PluginMock) CreateAccountCalls() []struct {
	Address common.Address
} {
	var calls []struct {
		Address common.Address
	}
	mock.lockCreateAccount.RLock()
	calls = mock.calls.CreateAccount
	mock.lockCreateAccount.RUnlock()
	return calls
}

// DeleteSuicides calls DeleteSuicidesFunc.
func (mock *PluginMock) DeleteSuicides(addresss []common.Address) {
	if mock.DeleteSuicidesFunc == nil {
		panic("PluginMock.DeleteSuicidesFunc: method is nil but Plugin.DeleteSuicides was just called")
	}
	callInfo := struct {
		Addresss []common.Address
	}{
		Addresss: addresss,
	}
	mock.lockDeleteSuicides.Lock()
	mock.calls.DeleteSuicides = append(mock.calls.DeleteSuicides, callInfo)
	mock.lockDeleteSuicides.Unlock()
	mock.DeleteSuicidesFunc(addresss)
}

// DeleteSuicidesCalls gets all the calls that were made to DeleteSuicides.
// Check the length with:
//
//	len(mockedPlugin.DeleteSuicidesCalls())
func (mock *PluginMock) DeleteSuicidesCalls() []struct {
	Addresss []common.Address
} {
	var calls []struct {
		Addresss []common.Address
	}
	mock.lockDeleteSuicides.RLock()
	calls = mock.calls.DeleteSuicides
	mock.lockDeleteSuicides.RUnlock()
	return calls
}

// Exist calls ExistFunc.
func (mock *PluginMock) Exist(address common.Address) bool {
	if mock.ExistFunc == nil {
		panic("PluginMock.ExistFunc: method is nil but Plugin.Exist was just called")
	}
	callInfo := struct {
		Address common.Address
	}{
		Address: address,
	}
	mock.lockExist.Lock()
	mock.calls.Exist = append(mock.calls.Exist, callInfo)
	mock.lockExist.Unlock()
	return mock.ExistFunc(address)
}

// ExistCalls gets all the calls that were made to Exist.
// Check the length with:
//
//	len(mockedPlugin.ExistCalls())
func (mock *PluginMock) ExistCalls() []struct {
	Address common.Address
} {
	var calls []struct {
		Address common.Address
	}
	mock.lockExist.RLock()
	calls = mock.calls.Exist
	mock.lockExist.RUnlock()
	return calls
}

// Finalize calls FinalizeFunc.
func (mock *PluginMock) Finalize() {
	if mock.FinalizeFunc == nil {
		panic("PluginMock.FinalizeFunc: method is nil but Plugin.Finalize was just called")
	}
	callInfo := struct {
	}{}
	mock.lockFinalize.Lock()
	mock.calls.Finalize = append(mock.calls.Finalize, callInfo)
	mock.lockFinalize.Unlock()
	mock.FinalizeFunc()
}

// FinalizeCalls gets all the calls that were made to Finalize.
// Check the length with:
//
//	len(mockedPlugin.FinalizeCalls())
func (mock *PluginMock) FinalizeCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockFinalize.RLock()
	calls = mock.calls.Finalize
	mock.lockFinalize.RUnlock()
	return calls
}

// ForEachStorage calls ForEachStorageFunc.
func (mock *PluginMock) ForEachStorage(address common.Address, fn func(common.Hash, common.Hash) bool) error {
	if mock.ForEachStorageFunc == nil {
		panic("PluginMock.ForEachStorageFunc: method is nil but Plugin.ForEachStorage was just called")
	}
	callInfo := struct {
		Address common.Address
		Fn      func(common.Hash, common.Hash) bool
	}{
		Address: address,
		Fn:      fn,
	}
	mock.lockForEachStorage.Lock()
	mock.calls.ForEachStorage = append(mock.calls.ForEachStorage, callInfo)
	mock.lockForEachStorage.Unlock()
	return mock.ForEachStorageFunc(address, fn)
}

// ForEachStorageCalls gets all the calls that were made to ForEachStorage.
// Check the length with:
//
//	len(mockedPlugin.ForEachStorageCalls())
func (mock *PluginMock) ForEachStorageCalls() []struct {
	Address common.Address
	Fn      func(common.Hash, common.Hash) bool
} {
	var calls []struct {
		Address common.Address
		Fn      func(common.Hash, common.Hash) bool
	}
	mock.lockForEachStorage.RLock()
	calls = mock.calls.ForEachStorage
	mock.lockForEachStorage.RUnlock()
	return calls
}

// GetBalance calls GetBalanceFunc.
func (mock *PluginMock) GetBalance(address common.Address) *big.Int {
	if mock.GetBalanceFunc == nil {
		panic("PluginMock.GetBalanceFunc: method is nil but Plugin.GetBalance was just called")
	}
	callInfo := struct {
		Address common.Address
	}{
		Address: address,
	}
	mock.lockGetBalance.Lock()
	mock.calls.GetBalance = append(mock.calls.GetBalance, callInfo)
	mock.lockGetBalance.Unlock()
	return mock.GetBalanceFunc(address)
}

// GetBalanceCalls gets all the calls that were made to GetBalance.
// Check the length with:
//
//	len(mockedPlugin.GetBalanceCalls())
func (mock *PluginMock) GetBalanceCalls() []struct {
	Address common.Address
} {
	var calls []struct {
		Address common.Address
	}
	mock.lockGetBalance.RLock()
	calls = mock.calls.GetBalance
	mock.lockGetBalance.RUnlock()
	return calls
}

// GetCode calls GetCodeFunc.
func (mock *PluginMock) GetCode(address common.Address) []byte {
	if mock.GetCodeFunc == nil {
		panic("PluginMock.GetCodeFunc: method is nil but Plugin.GetCode was just called")
	}
	callInfo := struct {
		Address common.Address
	}{
		Address: address,
	}
	mock.lockGetCode.Lock()
	mock.calls.GetCode = append(mock.calls.GetCode, callInfo)
	mock.lockGetCode.Unlock()
	return mock.GetCodeFunc(address)
}

// GetCodeCalls gets all the calls that were made to GetCode.
// Check the length with:
//
//	len(mockedPlugin.GetCodeCalls())
func (mock *PluginMock) GetCodeCalls() []struct {
	Address common.Address
} {
	var calls []struct {
		Address common.Address
	}
	mock.lockGetCode.RLock()
	calls = mock.calls.GetCode
	mock.lockGetCode.RUnlock()
	return calls
}

// GetCodeHash calls GetCodeHashFunc.
func (mock *PluginMock) GetCodeHash(address common.Address) common.Hash {
	if mock.GetCodeHashFunc == nil {
		panic("PluginMock.GetCodeHashFunc: method is nil but Plugin.GetCodeHash was just called")
	}
	callInfo := struct {
		Address common.Address
	}{
		Address: address,
	}
	mock.lockGetCodeHash.Lock()
	mock.calls.GetCodeHash = append(mock.calls.GetCodeHash, callInfo)
	mock.lockGetCodeHash.Unlock()
	return mock.GetCodeHashFunc(address)
}

// GetCodeHashCalls gets all the calls that were made to GetCodeHash.
// Check the length with:
//
//	len(mockedPlugin.GetCodeHashCalls())
func (mock *PluginMock) GetCodeHashCalls() []struct {
	Address common.Address
} {
	var calls []struct {
		Address common.Address
	}
	mock.lockGetCodeHash.RLock()
	calls = mock.calls.GetCodeHash
	mock.lockGetCodeHash.RUnlock()
	return calls
}

// GetCodeSize calls GetCodeSizeFunc.
func (mock *PluginMock) GetCodeSize(address common.Address) int {
	if mock.GetCodeSizeFunc == nil {
		panic("PluginMock.GetCodeSizeFunc: method is nil but Plugin.GetCodeSize was just called")
	}
	callInfo := struct {
		Address common.Address
	}{
		Address: address,
	}
	mock.lockGetCodeSize.Lock()
	mock.calls.GetCodeSize = append(mock.calls.GetCodeSize, callInfo)
	mock.lockGetCodeSize.Unlock()
	return mock.GetCodeSizeFunc(address)
}

// GetCodeSizeCalls gets all the calls that were made to GetCodeSize.
// Check the length with:
//
//	len(mockedPlugin.GetCodeSizeCalls())
func (mock *PluginMock) GetCodeSizeCalls() []struct {
	Address common.Address
} {
	var calls []struct {
		Address common.Address
	}
	mock.lockGetCodeSize.RLock()
	calls = mock.calls.GetCodeSize
	mock.lockGetCodeSize.RUnlock()
	return calls
}

// GetCommittedState calls GetCommittedStateFunc.
func (mock *PluginMock) GetCommittedState(address common.Address, hash common.Hash) common.Hash {
	if mock.GetCommittedStateFunc == nil {
		panic("PluginMock.GetCommittedStateFunc: method is nil but Plugin.GetCommittedState was just called")
	}
	callInfo := struct {
		Address common.Address
		Hash    common.Hash
	}{
		Address: address,
		Hash:    hash,
	}
	mock.lockGetCommittedState.Lock()
	mock.calls.GetCommittedState = append(mock.calls.GetCommittedState, callInfo)
	mock.lockGetCommittedState.Unlock()
	return mock.GetCommittedStateFunc(address, hash)
}

// GetCommittedStateCalls gets all the calls that were made to GetCommittedState.
// Check the length with:
//
//	len(mockedPlugin.GetCommittedStateCalls())
func (mock *PluginMock) GetCommittedStateCalls() []struct {
	Address common.Address
	Hash    common.Hash
} {
	var calls []struct {
		Address common.Address
		Hash    common.Hash
	}
	mock.lockGetCommittedState.RLock()
	calls = mock.calls.GetCommittedState
	mock.lockGetCommittedState.RUnlock()
	return calls
}

// GetNonce calls GetNonceFunc.
func (mock *PluginMock) GetNonce(address common.Address) uint64 {
	if mock.GetNonceFunc == nil {
		panic("PluginMock.GetNonceFunc: method is nil but Plugin.GetNonce was just called")
	}
	callInfo := struct {
		Address common.Address
	}{
		Address: address,
	}
	mock.lockGetNonce.Lock()
	mock.calls.GetNonce = append(mock.calls.GetNonce, callInfo)
	mock.lockGetNonce.Unlock()
	return mock.GetNonceFunc(address)
}

// GetNonceCalls gets all the calls that were made to GetNonce.
// Check the length with:
//
//	len(mockedPlugin.GetNonceCalls())
func (mock *PluginMock) GetNonceCalls() []struct {
	Address common.Address
} {
	var calls []struct {
		Address common.Address
	}
	mock.lockGetNonce.RLock()
	calls = mock.calls.GetNonce
	mock.lockGetNonce.RUnlock()
	return calls
}

// GetState calls GetStateFunc.
func (mock *PluginMock) GetState(address common.Address, hash common.Hash) common.Hash {
	if mock.GetStateFunc == nil {
		panic("PluginMock.GetStateFunc: method is nil but Plugin.GetState was just called")
	}
	callInfo := struct {
		Address common.Address
		Hash    common.Hash
	}{
		Address: address,
		Hash:    hash,
	}
	mock.lockGetState.Lock()
	mock.calls.GetState = append(mock.calls.GetState, callInfo)
	mock.lockGetState.Unlock()
	return mock.GetStateFunc(address, hash)
}

// GetStateCalls gets all the calls that were made to GetState.
// Check the length with:
//
//	len(mockedPlugin.GetStateCalls())
func (mock *PluginMock) GetStateCalls() []struct {
	Address common.Address
	Hash    common.Hash
} {
	var calls []struct {
		Address common.Address
		Hash    common.Hash
	}
	mock.lockGetState.RLock()
	calls = mock.calls.GetState
	mock.lockGetState.RUnlock()
	return calls
}

// RegistryKey calls RegistryKeyFunc.
func (mock *PluginMock) RegistryKey() string {
	if mock.RegistryKeyFunc == nil {
		panic("PluginMock.RegistryKeyFunc: method is nil but Plugin.RegistryKey was just called")
	}
	callInfo := struct {
	}{}
	mock.lockRegistryKey.Lock()
	mock.calls.RegistryKey = append(mock.calls.RegistryKey, callInfo)
	mock.lockRegistryKey.Unlock()
	return mock.RegistryKeyFunc()
}

// RegistryKeyCalls gets all the calls that were made to RegistryKey.
// Check the length with:
//
//	len(mockedPlugin.RegistryKeyCalls())
func (mock *PluginMock) RegistryKeyCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockRegistryKey.RLock()
	calls = mock.calls.RegistryKey
	mock.lockRegistryKey.RUnlock()
	return calls
}

// Reset calls ResetFunc.
func (mock *PluginMock) Reset(contextMoqParam context.Context) {
	if mock.ResetFunc == nil {
		panic("PluginMock.ResetFunc: method is nil but Plugin.Reset was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
	}{
		ContextMoqParam: contextMoqParam,
	}
	mock.lockReset.Lock()
	mock.calls.Reset = append(mock.calls.Reset, callInfo)
	mock.lockReset.Unlock()
	mock.ResetFunc(contextMoqParam)
}

// ResetCalls gets all the calls that were made to Reset.
// Check the length with:
//
//	len(mockedPlugin.ResetCalls())
func (mock *PluginMock) ResetCalls() []struct {
	ContextMoqParam context.Context
} {
	var calls []struct {
		ContextMoqParam context.Context
	}
	mock.lockReset.RLock()
	calls = mock.calls.Reset
	mock.lockReset.RUnlock()
	return calls
}

// RevertToSnapshot calls RevertToSnapshotFunc.
func (mock *PluginMock) RevertToSnapshot(n int) {
	if mock.RevertToSnapshotFunc == nil {
		panic("PluginMock.RevertToSnapshotFunc: method is nil but Plugin.RevertToSnapshot was just called")
	}
	callInfo := struct {
		N int
	}{
		N: n,
	}
	mock.lockRevertToSnapshot.Lock()
	mock.calls.RevertToSnapshot = append(mock.calls.RevertToSnapshot, callInfo)
	mock.lockRevertToSnapshot.Unlock()
	mock.RevertToSnapshotFunc(n)
}

// RevertToSnapshotCalls gets all the calls that were made to RevertToSnapshot.
// Check the length with:
//
//	len(mockedPlugin.RevertToSnapshotCalls())
func (mock *PluginMock) RevertToSnapshotCalls() []struct {
	N int
} {
	var calls []struct {
		N int
	}
	mock.lockRevertToSnapshot.RLock()
	calls = mock.calls.RevertToSnapshot
	mock.lockRevertToSnapshot.RUnlock()
	return calls
}

// SetBalance calls SetBalanceFunc.
func (mock *PluginMock) SetBalance(address common.Address, intMoqParam *big.Int) {
	if mock.SetBalanceFunc == nil {
		panic("PluginMock.SetBalanceFunc: method is nil but Plugin.SetBalance was just called")
	}
	callInfo := struct {
		Address     common.Address
		IntMoqParam *big.Int
	}{
		Address:     address,
		IntMoqParam: intMoqParam,
	}
	mock.lockSetBalance.Lock()
	mock.calls.SetBalance = append(mock.calls.SetBalance, callInfo)
	mock.lockSetBalance.Unlock()
	mock.SetBalanceFunc(address, intMoqParam)
}

// SetBalanceCalls gets all the calls that were made to SetBalance.
// Check the length with:
//
//	len(mockedPlugin.SetBalanceCalls())
func (mock *PluginMock) SetBalanceCalls() []struct {
	Address     common.Address
	IntMoqParam *big.Int
} {
	var calls []struct {
		Address     common.Address
		IntMoqParam *big.Int
	}
	mock.lockSetBalance.RLock()
	calls = mock.calls.SetBalance
	mock.lockSetBalance.RUnlock()
	return calls
}

// SetCode calls SetCodeFunc.
func (mock *PluginMock) SetCode(address common.Address, bytes []byte) {
	if mock.SetCodeFunc == nil {
		panic("PluginMock.SetCodeFunc: method is nil but Plugin.SetCode was just called")
	}
	callInfo := struct {
		Address common.Address
		Bytes   []byte
	}{
		Address: address,
		Bytes:   bytes,
	}
	mock.lockSetCode.Lock()
	mock.calls.SetCode = append(mock.calls.SetCode, callInfo)
	mock.lockSetCode.Unlock()
	mock.SetCodeFunc(address, bytes)
}

// SetCodeCalls gets all the calls that were made to SetCode.
// Check the length with:
//
//	len(mockedPlugin.SetCodeCalls())
func (mock *PluginMock) SetCodeCalls() []struct {
	Address common.Address
	Bytes   []byte
} {
	var calls []struct {
		Address common.Address
		Bytes   []byte
	}
	mock.lockSetCode.RLock()
	calls = mock.calls.SetCode
	mock.lockSetCode.RUnlock()
	return calls
}

// SetNonce calls SetNonceFunc.
func (mock *PluginMock) SetNonce(address common.Address, v uint64) {
	if mock.SetNonceFunc == nil {
		panic("PluginMock.SetNonceFunc: method is nil but Plugin.SetNonce was just called")
	}
	callInfo := struct {
		Address common.Address
		V       uint64
	}{
		Address: address,
		V:       v,
	}
	mock.lockSetNonce.Lock()
	mock.calls.SetNonce = append(mock.calls.SetNonce, callInfo)
	mock.lockSetNonce.Unlock()
	mock.SetNonceFunc(address, v)
}

// SetNonceCalls gets all the calls that were made to SetNonce.
// Check the length with:
//
//	len(mockedPlugin.SetNonceCalls())
func (mock *PluginMock) SetNonceCalls() []struct {
	Address common.Address
	V       uint64
} {
	var calls []struct {
		Address common.Address
		V       uint64
	}
	mock.lockSetNonce.RLock()
	calls = mock.calls.SetNonce
	mock.lockSetNonce.RUnlock()
	return calls
}

// SetState calls SetStateFunc.
func (mock *PluginMock) SetState(address common.Address, hash1 common.Hash, hash2 common.Hash) {
	if mock.SetStateFunc == nil {
		panic("PluginMock.SetStateFunc: method is nil but Plugin.SetState was just called")
	}
	callInfo := struct {
		Address common.Address
		Hash1   common.Hash
		Hash2   common.Hash
	}{
		Address: address,
		Hash1:   hash1,
		Hash2:   hash2,
	}
	mock.lockSetState.Lock()
	mock.calls.SetState = append(mock.calls.SetState, callInfo)
	mock.lockSetState.Unlock()
	mock.SetStateFunc(address, hash1, hash2)
}

// SetStateCalls gets all the calls that were made to SetState.
// Check the length with:
//
//	len(mockedPlugin.SetStateCalls())
func (mock *PluginMock) SetStateCalls() []struct {
	Address common.Address
	Hash1   common.Hash
	Hash2   common.Hash
} {
	var calls []struct {
		Address common.Address
		Hash1   common.Hash
		Hash2   common.Hash
	}
	mock.lockSetState.RLock()
	calls = mock.calls.SetState
	mock.lockSetState.RUnlock()
	return calls
}

// SetStorage calls SetStorageFunc.
func (mock *PluginMock) SetStorage(addr common.Address, storage map[common.Hash]common.Hash) {
	if mock.SetStorageFunc == nil {
		panic("PluginMock.SetStorageFunc: method is nil but Plugin.SetStorage was just called")
	}
	callInfo := struct {
		Addr    common.Address
		Storage map[common.Hash]common.Hash
	}{
		Addr:    addr,
		Storage: storage,
	}
	mock.lockSetStorage.Lock()
	mock.calls.SetStorage = append(mock.calls.SetStorage, callInfo)
	mock.lockSetStorage.Unlock()
	mock.SetStorageFunc(addr, storage)
}

// SetStorageCalls gets all the calls that were made to SetStorage.
// Check the length with:
//
//	len(mockedPlugin.SetStorageCalls())
func (mock *PluginMock) SetStorageCalls() []struct {
	Addr    common.Address
	Storage map[common.Hash]common.Hash
} {
	var calls []struct {
		Addr    common.Address
		Storage map[common.Hash]common.Hash
	}
	mock.lockSetStorage.RLock()
	calls = mock.calls.SetStorage
	mock.lockSetStorage.RUnlock()
	return calls
}

// Snapshot calls SnapshotFunc.
func (mock *PluginMock) Snapshot() int {
	if mock.SnapshotFunc == nil {
		panic("PluginMock.SnapshotFunc: method is nil but Plugin.Snapshot was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSnapshot.Lock()
	mock.calls.Snapshot = append(mock.calls.Snapshot, callInfo)
	mock.lockSnapshot.Unlock()
	return mock.SnapshotFunc()
}

// SnapshotCalls gets all the calls that were made to Snapshot.
// Check the length with:
//
//	len(mockedPlugin.SnapshotCalls())
func (mock *PluginMock) SnapshotCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSnapshot.RLock()
	calls = mock.calls.Snapshot
	mock.lockSnapshot.RUnlock()
	return calls
}

// SubBalance calls SubBalanceFunc.
func (mock *PluginMock) SubBalance(address common.Address, intMoqParam *big.Int) {
	if mock.SubBalanceFunc == nil {
		panic("PluginMock.SubBalanceFunc: method is nil but Plugin.SubBalance was just called")
	}
	callInfo := struct {
		Address     common.Address
		IntMoqParam *big.Int
	}{
		Address:     address,
		IntMoqParam: intMoqParam,
	}
	mock.lockSubBalance.Lock()
	mock.calls.SubBalance = append(mock.calls.SubBalance, callInfo)
	mock.lockSubBalance.Unlock()
	mock.SubBalanceFunc(address, intMoqParam)
}

// SubBalanceCalls gets all the calls that were made to SubBalance.
// Check the length with:
//
//	len(mockedPlugin.SubBalanceCalls())
func (mock *PluginMock) SubBalanceCalls() []struct {
	Address     common.Address
	IntMoqParam *big.Int
} {
	var calls []struct {
		Address     common.Address
		IntMoqParam *big.Int
	}
	mock.lockSubBalance.RLock()
	calls = mock.calls.SubBalance
	mock.lockSubBalance.RUnlock()
	return calls
}

// TransferBalance calls TransferBalanceFunc.
func (mock *PluginMock) TransferBalance(address1 common.Address, address2 common.Address, intMoqParam *big.Int) {
	if mock.TransferBalanceFunc == nil {
		panic("PluginMock.TransferBalanceFunc: method is nil but Plugin.TransferBalance was just called")
	}
	callInfo := struct {
		Address1    common.Address
		Address2    common.Address
		IntMoqParam *big.Int
	}{
		Address1:    address1,
		Address2:    address2,
		IntMoqParam: intMoqParam,
	}
	mock.lockTransferBalance.Lock()
	mock.calls.TransferBalance = append(mock.calls.TransferBalance, callInfo)
	mock.lockTransferBalance.Unlock()
	mock.TransferBalanceFunc(address1, address2, intMoqParam)
}

// TransferBalanceCalls gets all the calls that were made to TransferBalance.
// Check the length with:
//
//	len(mockedPlugin.TransferBalanceCalls())
func (mock *PluginMock) TransferBalanceCalls() []struct {
	Address1    common.Address
	Address2    common.Address
	IntMoqParam *big.Int
} {
	var calls []struct {
		Address1    common.Address
		Address2    common.Address
		IntMoqParam *big.Int
	}
	mock.lockTransferBalance.RLock()
	calls = mock.calls.TransferBalance
	mock.lockTransferBalance.RUnlock()
	return calls
}
