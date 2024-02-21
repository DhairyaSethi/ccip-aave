// Code generated by mockery v2.38.0. DO NOT EDIT.

package prices

import (
	context "context"
	big "math/big"

	cciptypes "github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/cciptypes"

	mock "github.com/stretchr/testify/mock"
)

// MockGasPriceEstimator is an autogenerated mock type for the GasPriceEstimator type
type MockGasPriceEstimator struct {
	mock.Mock
}

// DenoteInUSD provides a mock function with given fields: p, wrappedNativePrice
func (_m *MockGasPriceEstimator) DenoteInUSD(p *big.Int, wrappedNativePrice *big.Int) (*big.Int, error) {
	ret := _m.Called(p, wrappedNativePrice)

	if len(ret) == 0 {
		panic("no return value specified for DenoteInUSD")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) (*big.Int, error)); ok {
		return rf(p, wrappedNativePrice)
	}
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) *big.Int); ok {
		r0 = rf(p, wrappedNativePrice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int) error); ok {
		r1 = rf(p, wrappedNativePrice)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deviates provides a mock function with given fields: p1, p2
func (_m *MockGasPriceEstimator) Deviates(p1 *big.Int, p2 *big.Int) (bool, error) {
	ret := _m.Called(p1, p2)

	if len(ret) == 0 {
		panic("no return value specified for Deviates")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) (bool, error)); ok {
		return rf(p1, p2)
	}
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) bool); ok {
		r0 = rf(p1, p2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int) error); ok {
		r1 = rf(p1, p2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateMsgCostUSD provides a mock function with given fields: p, wrappedNativePrice, msg
func (_m *MockGasPriceEstimator) EstimateMsgCostUSD(p *big.Int, wrappedNativePrice *big.Int, msg cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta) (*big.Int, error) {
	ret := _m.Called(p, wrappedNativePrice, msg)

	if len(ret) == 0 {
		panic("no return value specified for EstimateMsgCostUSD")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta) (*big.Int, error)); ok {
		return rf(p, wrappedNativePrice, msg)
	}
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta) *big.Int); ok {
		r0 = rf(p, wrappedNativePrice, msg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int, cciptypes.EVM2EVMOnRampCCIPSendRequestedWithMeta) error); ok {
		r1 = rf(p, wrappedNativePrice, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasPrice provides a mock function with given fields: ctx
func (_m *MockGasPriceEstimator) GetGasPrice(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetGasPrice")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Median provides a mock function with given fields: gasPrices
func (_m *MockGasPriceEstimator) Median(gasPrices []*big.Int) (*big.Int, error) {
	ret := _m.Called(gasPrices)

	if len(ret) == 0 {
		panic("no return value specified for Median")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func([]*big.Int) (*big.Int, error)); ok {
		return rf(gasPrices)
	}
	if rf, ok := ret.Get(0).(func([]*big.Int) *big.Int); ok {
		r0 = rf(gasPrices)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func([]*big.Int) error); ok {
		r1 = rf(gasPrices)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockGasPriceEstimator creates a new instance of MockGasPriceEstimator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGasPriceEstimator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGasPriceEstimator {
	mock := &MockGasPriceEstimator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
