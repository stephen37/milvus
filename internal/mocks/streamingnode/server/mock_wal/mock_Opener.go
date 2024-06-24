// Code generated by mockery v2.32.4. DO NOT EDIT.

package mock_wal

import (
	context "context"

	wal "github.com/milvus-io/milvus/internal/streamingnode/server/wal"
	mock "github.com/stretchr/testify/mock"
)

// MockOpener is an autogenerated mock type for the Opener type
type MockOpener struct {
	mock.Mock
}

type MockOpener_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOpener) EXPECT() *MockOpener_Expecter {
	return &MockOpener_Expecter{mock: &_m.Mock}
}

// Close provides a mock function with given fields:
func (_m *MockOpener) Close() {
	_m.Called()
}

// MockOpener_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockOpener_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockOpener_Expecter) Close() *MockOpener_Close_Call {
	return &MockOpener_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockOpener_Close_Call) Run(run func()) *MockOpener_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockOpener_Close_Call) Return() *MockOpener_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockOpener_Close_Call) RunAndReturn(run func()) *MockOpener_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Open provides a mock function with given fields: ctx, opt
func (_m *MockOpener) Open(ctx context.Context, opt *wal.OpenOption) (wal.WAL, error) {
	ret := _m.Called(ctx, opt)

	var r0 wal.WAL
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *wal.OpenOption) (wal.WAL, error)); ok {
		return rf(ctx, opt)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *wal.OpenOption) wal.WAL); ok {
		r0 = rf(ctx, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(wal.WAL)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *wal.OpenOption) error); ok {
		r1 = rf(ctx, opt)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOpener_Open_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Open'
type MockOpener_Open_Call struct {
	*mock.Call
}

// Open is a helper method to define mock.On call
//   - ctx context.Context
//   - opt *wal.OpenOption
func (_e *MockOpener_Expecter) Open(ctx interface{}, opt interface{}) *MockOpener_Open_Call {
	return &MockOpener_Open_Call{Call: _e.mock.On("Open", ctx, opt)}
}

func (_c *MockOpener_Open_Call) Run(run func(ctx context.Context, opt *wal.OpenOption)) *MockOpener_Open_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*wal.OpenOption))
	})
	return _c
}

func (_c *MockOpener_Open_Call) Return(_a0 wal.WAL, _a1 error) *MockOpener_Open_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOpener_Open_Call) RunAndReturn(run func(context.Context, *wal.OpenOption) (wal.WAL, error)) *MockOpener_Open_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOpener creates a new instance of MockOpener. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOpener(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOpener {
	mock := &MockOpener{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
