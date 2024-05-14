// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	user "github.com/bangumi/server/internal/user"
	mock "github.com/stretchr/testify/mock"
)

// UserRepo is an autogenerated mock type for the Repo type
type UserRepo struct {
	mock.Mock
}

type UserRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepo) EXPECT() *UserRepo_Expecter {
	return &UserRepo_Expecter{mock: &_m.Mock}
}

// CheckIsFriendToOthers provides a mock function with given fields: ctx, selfID, otherIDs
func (_m *UserRepo) CheckIsFriendToOthers(ctx context.Context, selfID uint32, otherIDs ...uint32) (bool, error) {
	_va := make([]interface{}, len(otherIDs))
	for _i := range otherIDs {
		_va[_i] = otherIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, selfID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CheckIsFriendToOthers")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, ...uint32) (bool, error)); ok {
		return rf(ctx, selfID, otherIDs...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, ...uint32) bool); ok {
		r0 = rf(ctx, selfID, otherIDs...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, ...uint32) error); ok {
		r1 = rf(ctx, selfID, otherIDs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_CheckIsFriendToOthers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CheckIsFriendToOthers'
type UserRepo_CheckIsFriendToOthers_Call struct {
	*mock.Call
}

// CheckIsFriendToOthers is a helper method to define mock.On call
//   - ctx context.Context
//   - selfID uint32
//   - otherIDs ...uint32
func (_e *UserRepo_Expecter) CheckIsFriendToOthers(ctx interface{}, selfID interface{}, otherIDs ...interface{}) *UserRepo_CheckIsFriendToOthers_Call {
	return &UserRepo_CheckIsFriendToOthers_Call{Call: _e.mock.On("CheckIsFriendToOthers",
		append([]interface{}{ctx, selfID}, otherIDs...)...)}
}

func (_c *UserRepo_CheckIsFriendToOthers_Call) Run(run func(ctx context.Context, selfID uint32, otherIDs ...uint32)) *UserRepo_CheckIsFriendToOthers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]uint32, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(uint32)
			}
		}
		run(args[0].(context.Context), args[1].(uint32), variadicArgs...)
	})
	return _c
}

func (_c *UserRepo_CheckIsFriendToOthers_Call) Return(_a0 bool, _a1 error) *UserRepo_CheckIsFriendToOthers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_CheckIsFriendToOthers_Call) RunAndReturn(run func(context.Context, uint32, ...uint32) (bool, error)) *UserRepo_CheckIsFriendToOthers_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: ctx, userID
func (_m *UserRepo) GetByID(ctx context.Context, userID uint32) (user.User, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (user.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) user.User); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type UserRepo_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
func (_e *UserRepo_Expecter) GetByID(ctx interface{}, userID interface{}) *UserRepo_GetByID_Call {
	return &UserRepo_GetByID_Call{Call: _e.mock.On("GetByID", ctx, userID)}
}

func (_c *UserRepo_GetByID_Call) Run(run func(ctx context.Context, userID uint32)) *UserRepo_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *UserRepo_GetByID_Call) Return(_a0 user.User, _a1 error) *UserRepo_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_GetByID_Call) RunAndReturn(run func(context.Context, uint32) (user.User, error)) *UserRepo_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, ids
func (_m *UserRepo) GetByIDs(ctx context.Context, ids []uint32) (map[uint32]user.User, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for GetByIDs")
	}

	var r0 map[uint32]user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) (map[uint32]user.User, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) map[uint32]user.User); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type UserRepo_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []uint32
func (_e *UserRepo_Expecter) GetByIDs(ctx interface{}, ids interface{}) *UserRepo_GetByIDs_Call {
	return &UserRepo_GetByIDs_Call{Call: _e.mock.On("GetByIDs", ctx, ids)}
}

func (_c *UserRepo_GetByIDs_Call) Run(run func(ctx context.Context, ids []uint32)) *UserRepo_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint32))
	})
	return _c
}

func (_c *UserRepo_GetByIDs_Call) Return(_a0 map[uint32]user.User, _a1 error) *UserRepo_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_GetByIDs_Call) RunAndReturn(run func(context.Context, []uint32) (map[uint32]user.User, error)) *UserRepo_GetByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetByName provides a mock function with given fields: ctx, username
func (_m *UserRepo) GetByName(ctx context.Context, username string) (user.User, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetByName")
	}

	var r0 user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (user.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) user.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(user.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetByName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByName'
type UserRepo_GetByName_Call struct {
	*mock.Call
}

// GetByName is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *UserRepo_Expecter) GetByName(ctx interface{}, username interface{}) *UserRepo_GetByName_Call {
	return &UserRepo_GetByName_Call{Call: _e.mock.On("GetByName", ctx, username)}
}

func (_c *UserRepo_GetByName_Call) Run(run func(ctx context.Context, username string)) *UserRepo_GetByName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepo_GetByName_Call) Return(_a0 user.User, _a1 error) *UserRepo_GetByName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_GetByName_Call) RunAndReturn(run func(context.Context, string) (user.User, error)) *UserRepo_GetByName_Call {
	_c.Call.Return(run)
	return _c
}

// GetFieldsByIDs provides a mock function with given fields: ctx, ids
func (_m *UserRepo) GetFieldsByIDs(ctx context.Context, ids []uint32) (map[uint32]user.Fields, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for GetFieldsByIDs")
	}

	var r0 map[uint32]user.Fields
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) (map[uint32]user.Fields, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uint32) map[uint32]user.Fields); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]user.Fields)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uint32) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetFieldsByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFieldsByIDs'
type UserRepo_GetFieldsByIDs_Call struct {
	*mock.Call
}

// GetFieldsByIDs is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []uint32
func (_e *UserRepo_Expecter) GetFieldsByIDs(ctx interface{}, ids interface{}) *UserRepo_GetFieldsByIDs_Call {
	return &UserRepo_GetFieldsByIDs_Call{Call: _e.mock.On("GetFieldsByIDs", ctx, ids)}
}

func (_c *UserRepo_GetFieldsByIDs_Call) Run(run func(ctx context.Context, ids []uint32)) *UserRepo_GetFieldsByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uint32))
	})
	return _c
}

func (_c *UserRepo_GetFieldsByIDs_Call) Return(_a0 map[uint32]user.Fields, _a1 error) *UserRepo_GetFieldsByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_GetFieldsByIDs_Call) RunAndReturn(run func(context.Context, []uint32) (map[uint32]user.Fields, error)) *UserRepo_GetFieldsByIDs_Call {
	_c.Call.Return(run)
	return _c
}

// GetFriends provides a mock function with given fields: ctx, userID
func (_m *UserRepo) GetFriends(ctx context.Context, userID uint32) (map[uint32]user.FriendItem, error) {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for GetFriends")
	}

	var r0 map[uint32]user.FriendItem
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32) (map[uint32]user.FriendItem, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32) map[uint32]user.FriendItem); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[uint32]user.FriendItem)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepo_GetFriends_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFriends'
type UserRepo_GetFriends_Call struct {
	*mock.Call
}

// GetFriends is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
func (_e *UserRepo_Expecter) GetFriends(ctx interface{}, userID interface{}) *UserRepo_GetFriends_Call {
	return &UserRepo_GetFriends_Call{Call: _e.mock.On("GetFriends", ctx, userID)}
}

func (_c *UserRepo_GetFriends_Call) Run(run func(ctx context.Context, userID uint32)) *UserRepo_GetFriends_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *UserRepo_GetFriends_Call) Return(_a0 map[uint32]user.FriendItem, _a1 error) *UserRepo_GetFriends_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepo_GetFriends_Call) RunAndReturn(run func(context.Context, uint32) (map[uint32]user.FriendItem, error)) *UserRepo_GetFriends_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepo creates a new instance of UserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepo {
	mock := &UserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
