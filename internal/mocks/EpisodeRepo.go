// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	episode "github.com/bangumi/server/internal/episode"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"

	query "github.com/bangumi/server/internal/dal/query"
)

// EpisodeRepo is an autogenerated mock type for the Repo type
type EpisodeRepo struct {
	mock.Mock
}

type EpisodeRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *EpisodeRepo) EXPECT() *EpisodeRepo_Expecter {
	return &EpisodeRepo_Expecter{mock: &_m.Mock}
}

// Count provides a mock function with given fields: ctx, subjectID, filter
func (_m *EpisodeRepo) Count(ctx context.Context, subjectID model.SubjectID, filter episode.Filter) (int64, error) {
	ret := _m.Called(ctx, subjectID, filter)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID, episode.Filter) int64); ok {
		r0 = rf(ctx, subjectID, filter)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID, episode.Filter) error); ok {
		r1 = rf(ctx, subjectID, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EpisodeRepo_Count_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Count'
type EpisodeRepo_Count_Call struct {
	*mock.Call
}

// Count is a helper method to define mock.On call
//   - ctx context.Context
//   - subjectID model.SubjectID
//   - filter episode.Filter
func (_e *EpisodeRepo_Expecter) Count(ctx interface{}, subjectID interface{}, filter interface{}) *EpisodeRepo_Count_Call {
	return &EpisodeRepo_Count_Call{Call: _e.mock.On("Count", ctx, subjectID, filter)}
}

func (_c *EpisodeRepo_Count_Call) Run(run func(ctx context.Context, subjectID model.SubjectID, filter episode.Filter)) *EpisodeRepo_Count_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID), args[2].(episode.Filter))
	})
	return _c
}

func (_c *EpisodeRepo_Count_Call) Return(_a0 int64, _a1 error) *EpisodeRepo_Count_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Get provides a mock function with given fields: ctx, episodeID
func (_m *EpisodeRepo) Get(ctx context.Context, episodeID model.EpisodeID) (model.Episode, error) {
	ret := _m.Called(ctx, episodeID)

	var r0 model.Episode
	if rf, ok := ret.Get(0).(func(context.Context, model.EpisodeID) model.Episode); ok {
		r0 = rf(ctx, episodeID)
	} else {
		r0 = ret.Get(0).(model.Episode)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.EpisodeID) error); ok {
		r1 = rf(ctx, episodeID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EpisodeRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type EpisodeRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - episodeID model.EpisodeID
func (_e *EpisodeRepo_Expecter) Get(ctx interface{}, episodeID interface{}) *EpisodeRepo_Get_Call {
	return &EpisodeRepo_Get_Call{Call: _e.mock.On("Get", ctx, episodeID)}
}

func (_c *EpisodeRepo_Get_Call) Run(run func(ctx context.Context, episodeID model.EpisodeID)) *EpisodeRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.EpisodeID))
	})
	return _c
}

func (_c *EpisodeRepo_Get_Call) Return(_a0 model.Episode, _a1 error) *EpisodeRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// List provides a mock function with given fields: ctx, subjectID, filter, limit, offset
func (_m *EpisodeRepo) List(ctx context.Context, subjectID model.SubjectID, filter episode.Filter, limit int, offset int) ([]model.Episode, error) {
	ret := _m.Called(ctx, subjectID, filter, limit, offset)

	var r0 []model.Episode
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID, episode.Filter, int, int) []model.Episode); ok {
		r0 = rf(ctx, subjectID, filter, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Episode)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID, episode.Filter, int, int) error); ok {
		r1 = rf(ctx, subjectID, filter, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EpisodeRepo_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type EpisodeRepo_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - subjectID model.SubjectID
//   - filter episode.Filter
//   - limit int
//   - offset int
func (_e *EpisodeRepo_Expecter) List(ctx interface{}, subjectID interface{}, filter interface{}, limit interface{}, offset interface{}) *EpisodeRepo_List_Call {
	return &EpisodeRepo_List_Call{Call: _e.mock.On("List", ctx, subjectID, filter, limit, offset)}
}

func (_c *EpisodeRepo_List_Call) Run(run func(ctx context.Context, subjectID model.SubjectID, filter episode.Filter, limit int, offset int)) *EpisodeRepo_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID), args[2].(episode.Filter), args[3].(int), args[4].(int))
	})
	return _c
}

func (_c *EpisodeRepo_List_Call) Return(_a0 []model.Episode, _a1 error) *EpisodeRepo_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// WithQuery provides a mock function with given fields: _a0
func (_m *EpisodeRepo) WithQuery(_a0 *query.Query) episode.Repo {
	ret := _m.Called(_a0)

	var r0 episode.Repo
	if rf, ok := ret.Get(0).(func(*query.Query) episode.Repo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(episode.Repo)
		}
	}

	return r0
}

// EpisodeRepo_WithQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithQuery'
type EpisodeRepo_WithQuery_Call struct {
	*mock.Call
}

// WithQuery is a helper method to define mock.On call
//   - _a0 *query.Query
func (_e *EpisodeRepo_Expecter) WithQuery(_a0 interface{}) *EpisodeRepo_WithQuery_Call {
	return &EpisodeRepo_WithQuery_Call{Call: _e.mock.On("WithQuery", _a0)}
}

func (_c *EpisodeRepo_WithQuery_Call) Run(run func(_a0 *query.Query)) *EpisodeRepo_WithQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*query.Query))
	})
	return _c
}

func (_c *EpisodeRepo_WithQuery_Call) Return(_a0 episode.Repo) *EpisodeRepo_WithQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewEpisodeRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewEpisodeRepo creates a new instance of EpisodeRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEpisodeRepo(t mockConstructorTestingTNewEpisodeRepo) *EpisodeRepo {
	mock := &EpisodeRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
