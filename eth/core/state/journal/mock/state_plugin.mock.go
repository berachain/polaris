// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"sync"
)

// selfDestructStatePluginMock is a mock implementation of journal.selfDestructStatePlugin.
//
//	func TestSomethingThatUsesselfDestructStatePlugin(t *testing.T) {
//
//		// make and configure a mocked journal.selfDestructStatePlugin
//		mockedselfDestructStatePlugin := &selfDestructStatePluginMock{
//			GetBalanceFunc: func(address common.Address) *big.Int {
//				panic("mock out the GetBalance method")
//			},
//			GetCodeHashFunc: func(address common.Address) common.Hash {
//				panic("mock out the GetCodeHash method")
//			},
//			SubBalanceFunc: func(address common.Address, intMoqParam *big.Int)  {
//				panic("mock out the SubBalance method")
//			},
//		}
//
//		// use mockedselfDestructStatePlugin in code that requires journal.selfDestructStatePlugin
//		// and then make assertions.
//
//	}
type selfDestructStatePluginMock struct {
	// GetBalanceFunc mocks the GetBalance method.
	GetBalanceFunc func(address common.Address) *big.Int

	// GetCodeHashFunc mocks the GetCodeHash method.
	GetCodeHashFunc func(address common.Address) common.Hash

	// SubBalanceFunc mocks the SubBalance method.
	SubBalanceFunc func(address common.Address, intMoqParam *big.Int)

	// calls tracks calls to the methods.
	calls struct {
		// GetBalance holds details about calls to the GetBalance method.
		GetBalance []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// GetCodeHash holds details about calls to the GetCodeHash method.
		GetCodeHash []struct {
			// Address is the address argument value.
			Address common.Address
		}
		// SubBalance holds details about calls to the SubBalance method.
		SubBalance []struct {
			// Address is the address argument value.
			Address common.Address
			// IntMoqParam is the intMoqParam argument value.
			IntMoqParam *big.Int
		}
	}
	lockGetBalance  sync.RWMutex
	lockGetCodeHash sync.RWMutex
	lockSubBalance  sync.RWMutex
}

// GetBalance calls GetBalanceFunc.
func (mock *selfDestructStatePluginMock) GetBalance(address common.Address) *big.Int {
	if mock.GetBalanceFunc == nil {
		panic("selfDestructStatePluginMock.GetBalanceFunc: method is nil but selfDestructStatePlugin.GetBalance was just called")
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
//	len(mockedselfDestructStatePlugin.GetBalanceCalls())
func (mock *selfDestructStatePluginMock) GetBalanceCalls() []struct {
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

// GetCodeHash calls GetCodeHashFunc.
func (mock *selfDestructStatePluginMock) GetCodeHash(address common.Address) common.Hash {
	if mock.GetCodeHashFunc == nil {
		panic("selfDestructStatePluginMock.GetCodeHashFunc: method is nil but selfDestructStatePlugin.GetCodeHash was just called")
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
//	len(mockedselfDestructStatePlugin.GetCodeHashCalls())
func (mock *selfDestructStatePluginMock) GetCodeHashCalls() []struct {
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

// SubBalance calls SubBalanceFunc.
func (mock *selfDestructStatePluginMock) SubBalance(address common.Address, intMoqParam *big.Int) {
	if mock.SubBalanceFunc == nil {
		panic("selfDestructStatePluginMock.SubBalanceFunc: method is nil but selfDestructStatePlugin.SubBalance was just called")
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
//	len(mockedselfDestructStatePlugin.SubBalanceCalls())
func (mock *selfDestructStatePluginMock) SubBalanceCalls() []struct {
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
