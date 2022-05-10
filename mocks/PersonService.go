// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/model"

	testing "testing"
)

// PersonService is an autogenerated mock type for the PersonService type
type PersonService struct {
	mock.Mock
}

type PersonService_Expecter struct {
	mock *mock.Mock
}

func (_m *PersonService) EXPECT() *PersonService_Expecter {
	return &PersonService_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *PersonService) Get(ctx context.Context, id uint32) (model.Person, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Person
	if rf, ok := ret.Get(0).(func(context.Context, uint32) model.Person); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Person)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonService_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PersonService_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - id uint32
func (_e *PersonService_Expecter) Get(ctx interface{}, id interface{}) *PersonService_Get_Call {
	return &PersonService_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *PersonService_Get_Call) Run(run func(ctx context.Context, id uint32)) *PersonService_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *PersonService_Get_Call) Return(_a0 model.Person, _a1 error) *PersonService_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCharacterRelated provides a mock function with given fields: ctx, characterID
func (_m *PersonService) GetCharacterRelated(ctx context.Context, characterID uint32) ([]model.PersonCharacterRelation, error) {
	ret := _m.Called(ctx, characterID)

	var r0 []model.PersonCharacterRelation
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []model.PersonCharacterRelation); ok {
		r0 = rf(ctx, characterID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PersonCharacterRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, characterID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonService_GetCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCharacterRelated'
type PersonService_GetCharacterRelated_Call struct {
	*mock.Call
}

// GetCharacterRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - characterID uint32
func (_e *PersonService_Expecter) GetCharacterRelated(ctx interface{}, characterID interface{}) *PersonService_GetCharacterRelated_Call {
	return &PersonService_GetCharacterRelated_Call{Call: _e.mock.On("GetCharacterRelated", ctx, characterID)}
}

func (_c *PersonService_GetCharacterRelated_Call) Run(run func(ctx context.Context, characterID uint32)) *PersonService_GetCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *PersonService_GetCharacterRelated_Call) Return(_a0 []model.PersonCharacterRelation, _a1 error) *PersonService_GetCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectRelated provides a mock function with given fields: ctx, subjectID
func (_m *PersonService) GetSubjectRelated(ctx context.Context, subjectID uint32) ([]model.SubjectPersonRelation, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 []model.SubjectPersonRelation
	if rf, ok := ret.Get(0).(func(context.Context, uint32) []model.SubjectPersonRelation); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.SubjectPersonRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint32) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonService_GetSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectRelated'
type PersonService_GetSubjectRelated_Call struct {
	*mock.Call
}

// GetSubjectRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID uint32
func (_e *PersonService_Expecter) GetSubjectRelated(ctx interface{}, subjectID interface{}) *PersonService_GetSubjectRelated_Call {
	return &PersonService_GetSubjectRelated_Call{Call: _e.mock.On("GetSubjectRelated", ctx, subjectID)}
}

func (_c *PersonService_GetSubjectRelated_Call) Run(run func(ctx context.Context, subjectID uint32)) *PersonService_GetSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32))
	})
	return _c
}

func (_c *PersonService_GetSubjectRelated_Call) Return(_a0 []model.SubjectPersonRelation, _a1 error) *PersonService_GetSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// NewPersonService creates a new instance of PersonService. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersonService(t testing.TB) *PersonService {
	mock := &PersonService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
