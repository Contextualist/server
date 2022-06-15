// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// CaptchaManager is an autogenerated mock type for the Manager type
type CaptchaManager struct {
	mock.Mock
}

type CaptchaManager_Expecter struct {
	mock *mock.Mock
}

func (_m *CaptchaManager) EXPECT() *CaptchaManager_Expecter {
	return &CaptchaManager_Expecter{mock: &_m.Mock}
}

// Verify provides a mock function with given fields: ctx, response
func (_m *CaptchaManager) Verify(ctx context.Context, response string) (bool, error) {
	ret := _m.Called(ctx, response)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, response)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, response)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CaptchaManager_Verify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Verify'
type CaptchaManager_Verify_Call struct {
	*mock.Call
}

// Verify is a helper method to define mock.On call
//  - ctx context.Context
//  - response string
func (_e *CaptchaManager_Expecter) Verify(ctx interface{}, response interface{}) *CaptchaManager_Verify_Call {
	return &CaptchaManager_Verify_Call{Call: _e.mock.On("Verify", ctx, response)}
}

func (_c *CaptchaManager_Verify_Call) Run(run func(ctx context.Context, response string)) *CaptchaManager_Verify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *CaptchaManager_Verify_Call) Return(_a0 bool, _a1 error) *CaptchaManager_Verify_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type NewCaptchaManagerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewCaptchaManager creates a new instance of CaptchaManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCaptchaManager(t NewCaptchaManagerT) *CaptchaManager {
	mock := &CaptchaManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
