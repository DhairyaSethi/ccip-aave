// Code generated by mockery v2.38.0. DO NOT EDIT.

package mock_arbitrum_inbox

import (
	big "math/big"

	arbitrum_inbox "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/rebalancer/generated/arbitrum_inbox"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	event "github.com/ethereum/go-ethereum/event"

	generated "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// ArbitrumInboxInterface is an autogenerated mock type for the ArbitrumInboxInterface type
type ArbitrumInboxInterface struct {
	mock.Mock
}

// Address provides a mock function with given fields:
func (_m *ArbitrumInboxInterface) Address() common.Address {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Address")
	}

	var r0 common.Address
	if rf, ok := ret.Get(0).(func() common.Address); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	return r0
}

// AllowListEnabled provides a mock function with given fields: opts
func (_m *ArbitrumInboxInterface) AllowListEnabled(opts *bind.CallOpts) (bool, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for AllowListEnabled")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (bool, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) bool); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Bridge provides a mock function with given fields: opts
func (_m *ArbitrumInboxInterface) Bridge(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for Bridge")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (common.Address, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateRetryableSubmissionFee provides a mock function with given fields: opts, dataLength, baseFee
func (_m *ArbitrumInboxInterface) CalculateRetryableSubmissionFee(opts *bind.CallOpts, dataLength *big.Int, baseFee *big.Int) (*big.Int, error) {
	ret := _m.Called(opts, dataLength, baseFee)

	if len(ret) == 0 {
		panic("no return value specified for CalculateRetryableSubmissionFee")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, *big.Int) (*big.Int, error)); ok {
		return rf(opts, dataLength, baseFee)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(opts, dataLength, baseFee)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int, *big.Int) error); ok {
		r1 = rf(opts, dataLength, baseFee)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInboxMessageDelivered provides a mock function with given fields: opts, messageNum
func (_m *ArbitrumInboxInterface) FilterInboxMessageDelivered(opts *bind.FilterOpts, messageNum []*big.Int) (*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredIterator, error) {
	ret := _m.Called(opts, messageNum)

	if len(ret) == 0 {
		panic("no return value specified for FilterInboxMessageDelivered")
	}

	var r0 *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []*big.Int) (*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredIterator, error)); ok {
		return rf(opts, messageNum)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []*big.Int) *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredIterator); ok {
		r0 = rf(opts, messageNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []*big.Int) error); ok {
		r1 = rf(opts, messageNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterInboxMessageDeliveredFromOrigin provides a mock function with given fields: opts, messageNum
func (_m *ArbitrumInboxInterface) FilterInboxMessageDeliveredFromOrigin(opts *bind.FilterOpts, messageNum []*big.Int) (*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOriginIterator, error) {
	ret := _m.Called(opts, messageNum)

	if len(ret) == 0 {
		panic("no return value specified for FilterInboxMessageDeliveredFromOrigin")
	}

	var r0 *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOriginIterator
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []*big.Int) (*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOriginIterator, error)); ok {
		return rf(opts, messageNum)
	}
	if rf, ok := ret.Get(0).(func(*bind.FilterOpts, []*big.Int) *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOriginIterator); ok {
		r0 = rf(opts, messageNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOriginIterator)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.FilterOpts, []*big.Int) error); ok {
		r1 = rf(opts, messageNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProxyAdmin provides a mock function with given fields: opts
func (_m *ArbitrumInboxInterface) GetProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for GetProxyAdmin")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (common.Address, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Initialize provides a mock function with given fields: opts, _bridge, _sequencerInbox
func (_m *ArbitrumInboxInterface) Initialize(opts *bind.TransactOpts, _bridge common.Address, _sequencerInbox common.Address) (*types.Transaction, error) {
	ret := _m.Called(opts, _bridge, _sequencerInbox)

	if len(ret) == 0 {
		panic("no return value specified for Initialize")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, common.Address) (*types.Transaction, error)); ok {
		return rf(opts, _bridge, _sequencerInbox)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, common.Address, common.Address) *types.Transaction); ok {
		r0 = rf(opts, _bridge, _sequencerInbox)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, common.Address, common.Address) error); ok {
		r1 = rf(opts, _bridge, _sequencerInbox)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAllowed provides a mock function with given fields: opts, user
func (_m *ArbitrumInboxInterface) IsAllowed(opts *bind.CallOpts, user common.Address) (bool, error) {
	ret := _m.Called(opts, user)

	if len(ret) == 0 {
		panic("no return value specified for IsAllowed")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) (bool, error)); ok {
		return rf(opts, user)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, common.Address) bool); ok {
		r0 = rf(opts, user)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, common.Address) error); ok {
		r1 = rf(opts, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaxDataSize provides a mock function with given fields: opts
func (_m *ArbitrumInboxInterface) MaxDataSize(opts *bind.CallOpts) (*big.Int, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for MaxDataSize")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (*big.Int, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) *big.Int); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInboxMessageDelivered provides a mock function with given fields: log
func (_m *ArbitrumInboxInterface) ParseInboxMessageDelivered(log types.Log) (*arbitrum_inbox.ArbitrumInboxInboxMessageDelivered, error) {
	ret := _m.Called(log)

	if len(ret) == 0 {
		panic("no return value specified for ParseInboxMessageDelivered")
	}

	var r0 *arbitrum_inbox.ArbitrumInboxInboxMessageDelivered
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*arbitrum_inbox.ArbitrumInboxInboxMessageDelivered, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *arbitrum_inbox.ArbitrumInboxInboxMessageDelivered); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*arbitrum_inbox.ArbitrumInboxInboxMessageDelivered)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInboxMessageDeliveredFromOrigin provides a mock function with given fields: log
func (_m *ArbitrumInboxInterface) ParseInboxMessageDeliveredFromOrigin(log types.Log) (*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin, error) {
	ret := _m.Called(log)

	if len(ret) == 0 {
		panic("no return value specified for ParseInboxMessageDeliveredFromOrigin")
	}

	var r0 *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseLog provides a mock function with given fields: log
func (_m *ArbitrumInboxInterface) ParseLog(log types.Log) (generated.AbigenLog, error) {
	ret := _m.Called(log)

	if len(ret) == 0 {
		panic("no return value specified for ParseLog")
	}

	var r0 generated.AbigenLog
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (generated.AbigenLog, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) generated.AbigenLog); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.AbigenLog)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Pause provides a mock function with given fields: opts
func (_m *ArbitrumInboxInterface) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for Pause")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) (*types.Transaction, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendContractTransaction provides a mock function with given fields: opts, gasLimit, maxFeePerGas, to, value, data
func (_m *ArbitrumInboxInterface) SendContractTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, gasLimit, maxFeePerGas, to, value, data)

	if len(ret) == 0 {
		panic("no return value specified for SendContractTransaction")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, *big.Int, common.Address, *big.Int, []byte) (*types.Transaction, error)); ok {
		return rf(opts, gasLimit, maxFeePerGas, to, value, data)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, *big.Int, common.Address, *big.Int, []byte) *types.Transaction); ok {
		r0 = rf(opts, gasLimit, maxFeePerGas, to, value, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int, *big.Int, common.Address, *big.Int, []byte) error); ok {
		r1 = rf(opts, gasLimit, maxFeePerGas, to, value, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendL2Message provides a mock function with given fields: opts, messageData
func (_m *ArbitrumInboxInterface) SendL2Message(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, messageData)

	if len(ret) == 0 {
		panic("no return value specified for SendL2Message")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte) (*types.Transaction, error)); ok {
		return rf(opts, messageData)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte) *types.Transaction); ok {
		r0 = rf(opts, messageData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []byte) error); ok {
		r1 = rf(opts, messageData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendL2MessageFromOrigin provides a mock function with given fields: opts, messageData
func (_m *ArbitrumInboxInterface) SendL2MessageFromOrigin(opts *bind.TransactOpts, messageData []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, messageData)

	if len(ret) == 0 {
		panic("no return value specified for SendL2MessageFromOrigin")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte) (*types.Transaction, error)); ok {
		return rf(opts, messageData)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []byte) *types.Transaction); ok {
		r0 = rf(opts, messageData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []byte) error); ok {
		r1 = rf(opts, messageData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendUnsignedTransaction provides a mock function with given fields: opts, gasLimit, maxFeePerGas, nonce, to, value, data
func (_m *ArbitrumInboxInterface) SendUnsignedTransaction(opts *bind.TransactOpts, gasLimit *big.Int, maxFeePerGas *big.Int, nonce *big.Int, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	ret := _m.Called(opts, gasLimit, maxFeePerGas, nonce, to, value, data)

	if len(ret) == 0 {
		panic("no return value specified for SendUnsignedTransaction")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, []byte) (*types.Transaction, error)); ok {
		return rf(opts, gasLimit, maxFeePerGas, nonce, to, value, data)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, []byte) *types.Transaction); ok {
		r0 = rf(opts, gasLimit, maxFeePerGas, nonce, to, value, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, *big.Int, *big.Int, *big.Int, common.Address, *big.Int, []byte) error); ok {
		r1 = rf(opts, gasLimit, maxFeePerGas, nonce, to, value, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SequencerInbox provides a mock function with given fields: opts
func (_m *ArbitrumInboxInterface) SequencerInbox(opts *bind.CallOpts) (common.Address, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for SequencerInbox")
	}

	var r0 common.Address
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (common.Address, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) common.Address); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Address)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAllowList provides a mock function with given fields: opts, user, val
func (_m *ArbitrumInboxInterface) SetAllowList(opts *bind.TransactOpts, user []common.Address, val []bool) (*types.Transaction, error) {
	ret := _m.Called(opts, user, val)

	if len(ret) == 0 {
		panic("no return value specified for SetAllowList")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address, []bool) (*types.Transaction, error)); ok {
		return rf(opts, user, val)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, []common.Address, []bool) *types.Transaction); ok {
		r0 = rf(opts, user, val)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, []common.Address, []bool) error); ok {
		r1 = rf(opts, user, val)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAllowListEnabled provides a mock function with given fields: opts, _allowListEnabled
func (_m *ArbitrumInboxInterface) SetAllowListEnabled(opts *bind.TransactOpts, _allowListEnabled bool) (*types.Transaction, error) {
	ret := _m.Called(opts, _allowListEnabled)

	if len(ret) == 0 {
		panic("no return value specified for SetAllowListEnabled")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, bool) (*types.Transaction, error)); ok {
		return rf(opts, _allowListEnabled)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts, bool) *types.Transaction); ok {
		r0 = rf(opts, _allowListEnabled)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts, bool) error); ok {
		r1 = rf(opts, _allowListEnabled)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unpause provides a mock function with given fields: opts
func (_m *ArbitrumInboxInterface) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for Unpause")
	}

	var r0 *types.Transaction
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) (*types.Transaction, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.TransactOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInboxMessageDelivered provides a mock function with given fields: opts, sink, messageNum
func (_m *ArbitrumInboxInterface) WatchInboxMessageDelivered(opts *bind.WatchOpts, sink chan<- *arbitrum_inbox.ArbitrumInboxInboxMessageDelivered, messageNum []*big.Int) (event.Subscription, error) {
	ret := _m.Called(opts, sink, messageNum)

	if len(ret) == 0 {
		panic("no return value specified for WatchInboxMessageDelivered")
	}

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *arbitrum_inbox.ArbitrumInboxInboxMessageDelivered, []*big.Int) (event.Subscription, error)); ok {
		return rf(opts, sink, messageNum)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *arbitrum_inbox.ArbitrumInboxInboxMessageDelivered, []*big.Int) event.Subscription); ok {
		r0 = rf(opts, sink, messageNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *arbitrum_inbox.ArbitrumInboxInboxMessageDelivered, []*big.Int) error); ok {
		r1 = rf(opts, sink, messageNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchInboxMessageDeliveredFromOrigin provides a mock function with given fields: opts, sink, messageNum
func (_m *ArbitrumInboxInterface) WatchInboxMessageDeliveredFromOrigin(opts *bind.WatchOpts, sink chan<- *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin, messageNum []*big.Int) (event.Subscription, error) {
	ret := _m.Called(opts, sink, messageNum)

	if len(ret) == 0 {
		panic("no return value specified for WatchInboxMessageDeliveredFromOrigin")
	}

	var r0 event.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin, []*big.Int) (event.Subscription, error)); ok {
		return rf(opts, sink, messageNum)
	}
	if rf, ok := ret.Get(0).(func(*bind.WatchOpts, chan<- *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin, []*big.Int) event.Subscription); ok {
		r0 = rf(opts, sink, messageNum)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(event.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.WatchOpts, chan<- *arbitrum_inbox.ArbitrumInboxInboxMessageDeliveredFromOrigin, []*big.Int) error); ok {
		r1 = rf(opts, sink, messageNum)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewArbitrumInboxInterface creates a new instance of ArbitrumInboxInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArbitrumInboxInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArbitrumInboxInterface {
	mock := &ArbitrumInboxInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
