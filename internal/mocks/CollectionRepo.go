// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	collections "github.com/bangumi/server/internal/collections"
	collection "github.com/bangumi/server/internal/collections/domain/collection"

	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"

	query "github.com/bangumi/server/dal/query"

	time "time"
)

// CollectionRepo is an autogenerated mock type for the Repo type
type CollectionRepo struct {
	mock.Mock
}

type CollectionRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *CollectionRepo) EXPECT() *CollectionRepo_Expecter {
	return &CollectionRepo_Expecter{mock: &_m.Mock}
}

// CountSubjectCollections provides a mock function with given fields: ctx, userID, subjectType, collectionType, showPrivate
func (_m *CollectionRepo) CountSubjectCollections(ctx context.Context, userID uint32, subjectType uint8, collectionType collection.SubjectCollection, showPrivate bool) (int64, error) {
	ret := _m.Called(ctx, userID, subjectType, collectionType, showPrivate)

	if len(ret) == 0 {
		panic("no return value specified for CountSubjectCollections")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, collection.SubjectCollection, bool) (int64, error)); ok {
		return rf(ctx, userID, subjectType, collectionType, showPrivate)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, collection.SubjectCollection, bool) int64); ok {
		r0 = rf(ctx, userID, subjectType, collectionType, showPrivate)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8, collection.SubjectCollection, bool) error); ok {
		r1 = rf(ctx, userID, subjectType, collectionType, showPrivate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_CountSubjectCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountSubjectCollections'
type CollectionRepo_CountSubjectCollections_Call struct {
	*mock.Call
}

// CountSubjectCollections is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - subjectType uint8
//   - collectionType collection.SubjectCollection
//   - showPrivate bool
func (_e *CollectionRepo_Expecter) CountSubjectCollections(ctx interface{}, userID interface{}, subjectType interface{}, collectionType interface{}, showPrivate interface{}) *CollectionRepo_CountSubjectCollections_Call {
	return &CollectionRepo_CountSubjectCollections_Call{Call: _e.mock.On("CountSubjectCollections", ctx, userID, subjectType, collectionType, showPrivate)}
}

func (_c *CollectionRepo_CountSubjectCollections_Call) Run(run func(ctx context.Context, userID uint32, subjectType uint8, collectionType collection.SubjectCollection, showPrivate bool)) *CollectionRepo_CountSubjectCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8), args[3].(collection.SubjectCollection), args[4].(bool))
	})
	return _c
}

func (_c *CollectionRepo_CountSubjectCollections_Call) Return(_a0 int64, _a1 error) *CollectionRepo_CountSubjectCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CollectionRepo_CountSubjectCollections_Call) RunAndReturn(run func(context.Context, uint32, uint8, collection.SubjectCollection, bool) (int64, error)) *CollectionRepo_CountSubjectCollections_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubjectCollection provides a mock function with given fields: ctx, userID, subjectID
func (_m *CollectionRepo) GetSubjectCollection(ctx context.Context, userID uint32, subjectID uint32) (collection.UserSubjectCollection, error) {
	ret := _m.Called(ctx, userID, subjectID)

	if len(ret) == 0 {
		panic("no return value specified for GetSubjectCollection")
	}

	var r0 collection.UserSubjectCollection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) (collection.UserSubjectCollection, error)); ok {
		return rf(ctx, userID, subjectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) collection.UserSubjectCollection); ok {
		r0 = rf(ctx, userID, subjectID)
	} else {
		r0 = ret.Get(0).(collection.UserSubjectCollection)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint32) error); ok {
		r1 = rf(ctx, userID, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_GetSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectCollection'
type CollectionRepo_GetSubjectCollection_Call struct {
	*mock.Call
}

// GetSubjectCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - subjectID uint32
func (_e *CollectionRepo_Expecter) GetSubjectCollection(ctx interface{}, userID interface{}, subjectID interface{}) *CollectionRepo_GetSubjectCollection_Call {
	return &CollectionRepo_GetSubjectCollection_Call{Call: _e.mock.On("GetSubjectCollection", ctx, userID, subjectID)}
}

func (_c *CollectionRepo_GetSubjectCollection_Call) Run(run func(ctx context.Context, userID uint32, subjectID uint32)) *CollectionRepo_GetSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *CollectionRepo_GetSubjectCollection_Call) Return(_a0 collection.UserSubjectCollection, _a1 error) *CollectionRepo_GetSubjectCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CollectionRepo_GetSubjectCollection_Call) RunAndReturn(run func(context.Context, uint32, uint32) (collection.UserSubjectCollection, error)) *CollectionRepo_GetSubjectCollection_Call {
	_c.Call.Return(run)
	return _c
}

// GetSubjectEpisodesCollection provides a mock function with given fields: ctx, userID, subjectID
func (_m *CollectionRepo) GetSubjectEpisodesCollection(ctx context.Context, userID uint32, subjectID uint32) (collection.UserSubjectEpisodesCollection, error) {
	ret := _m.Called(ctx, userID, subjectID)

	if len(ret) == 0 {
		panic("no return value specified for GetSubjectEpisodesCollection")
	}

	var r0 collection.UserSubjectEpisodesCollection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) (collection.UserSubjectEpisodesCollection, error)); ok {
		return rf(ctx, userID, subjectID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32) collection.UserSubjectEpisodesCollection); ok {
		r0 = rf(ctx, userID, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(collection.UserSubjectEpisodesCollection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint32) error); ok {
		r1 = rf(ctx, userID, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_GetSubjectEpisodesCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectEpisodesCollection'
type CollectionRepo_GetSubjectEpisodesCollection_Call struct {
	*mock.Call
}

// GetSubjectEpisodesCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - subjectID uint32
func (_e *CollectionRepo_Expecter) GetSubjectEpisodesCollection(ctx interface{}, userID interface{}, subjectID interface{}) *CollectionRepo_GetSubjectEpisodesCollection_Call {
	return &CollectionRepo_GetSubjectEpisodesCollection_Call{Call: _e.mock.On("GetSubjectEpisodesCollection", ctx, userID, subjectID)}
}

func (_c *CollectionRepo_GetSubjectEpisodesCollection_Call) Run(run func(ctx context.Context, userID uint32, subjectID uint32)) *CollectionRepo_GetSubjectEpisodesCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32))
	})
	return _c
}

func (_c *CollectionRepo_GetSubjectEpisodesCollection_Call) Return(_a0 collection.UserSubjectEpisodesCollection, _a1 error) *CollectionRepo_GetSubjectEpisodesCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CollectionRepo_GetSubjectEpisodesCollection_Call) RunAndReturn(run func(context.Context, uint32, uint32) (collection.UserSubjectEpisodesCollection, error)) *CollectionRepo_GetSubjectEpisodesCollection_Call {
	_c.Call.Return(run)
	return _c
}

// ListSubjectCollection provides a mock function with given fields: ctx, userID, subjectType, collectionType, showPrivate, limit, offset
func (_m *CollectionRepo) ListSubjectCollection(ctx context.Context, userID uint32, subjectType uint8, collectionType collection.SubjectCollection, showPrivate bool, limit int, offset int) ([]collection.UserSubjectCollection, error) {
	ret := _m.Called(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)

	if len(ret) == 0 {
		panic("no return value specified for ListSubjectCollection")
	}

	var r0 []collection.UserSubjectCollection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, collection.SubjectCollection, bool, int, int) ([]collection.UserSubjectCollection, error)); ok {
		return rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint8, collection.SubjectCollection, bool, int, int) []collection.UserSubjectCollection); ok {
		r0 = rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]collection.UserSubjectCollection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint8, collection.SubjectCollection, bool, int, int) error); ok {
		r1 = rf(ctx, userID, subjectType, collectionType, showPrivate, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_ListSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListSubjectCollection'
type CollectionRepo_ListSubjectCollection_Call struct {
	*mock.Call
}

// ListSubjectCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - subjectType uint8
//   - collectionType collection.SubjectCollection
//   - showPrivate bool
//   - limit int
//   - offset int
func (_e *CollectionRepo_Expecter) ListSubjectCollection(ctx interface{}, userID interface{}, subjectType interface{}, collectionType interface{}, showPrivate interface{}, limit interface{}, offset interface{}) *CollectionRepo_ListSubjectCollection_Call {
	return &CollectionRepo_ListSubjectCollection_Call{Call: _e.mock.On("ListSubjectCollection", ctx, userID, subjectType, collectionType, showPrivate, limit, offset)}
}

func (_c *CollectionRepo_ListSubjectCollection_Call) Run(run func(ctx context.Context, userID uint32, subjectType uint8, collectionType collection.SubjectCollection, showPrivate bool, limit int, offset int)) *CollectionRepo_ListSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint8), args[3].(collection.SubjectCollection), args[4].(bool), args[5].(int), args[6].(int))
	})
	return _c
}

func (_c *CollectionRepo_ListSubjectCollection_Call) Return(_a0 []collection.UserSubjectCollection, _a1 error) *CollectionRepo_ListSubjectCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CollectionRepo_ListSubjectCollection_Call) RunAndReturn(run func(context.Context, uint32, uint8, collection.SubjectCollection, bool, int, int) ([]collection.UserSubjectCollection, error)) *CollectionRepo_ListSubjectCollection_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEpisodeCollection provides a mock function with given fields: ctx, userID, subjectID, episodeIDs, _a4, at
func (_m *CollectionRepo) UpdateEpisodeCollection(ctx context.Context, userID uint32, subjectID uint32, episodeIDs []uint32, _a4 collection.EpisodeCollection, at time.Time) (collection.UserSubjectEpisodesCollection, error) {
	ret := _m.Called(ctx, userID, subjectID, episodeIDs, _a4, at)

	if len(ret) == 0 {
		panic("no return value specified for UpdateEpisodeCollection")
	}

	var r0 collection.UserSubjectEpisodesCollection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32, []uint32, collection.EpisodeCollection, time.Time) (collection.UserSubjectEpisodesCollection, error)); ok {
		return rf(ctx, userID, subjectID, episodeIDs, _a4, at)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint32, uint32, []uint32, collection.EpisodeCollection, time.Time) collection.UserSubjectEpisodesCollection); ok {
		r0 = rf(ctx, userID, subjectID, episodeIDs, _a4, at)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(collection.UserSubjectEpisodesCollection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint32, uint32, []uint32, collection.EpisodeCollection, time.Time) error); ok {
		r1 = rf(ctx, userID, subjectID, episodeIDs, _a4, at)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CollectionRepo_UpdateEpisodeCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEpisodeCollection'
type CollectionRepo_UpdateEpisodeCollection_Call struct {
	*mock.Call
}

// UpdateEpisodeCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - subjectID uint32
//   - episodeIDs []uint32
//   - _a4 collection.EpisodeCollection
//   - at time.Time
func (_e *CollectionRepo_Expecter) UpdateEpisodeCollection(ctx interface{}, userID interface{}, subjectID interface{}, episodeIDs interface{}, _a4 interface{}, at interface{}) *CollectionRepo_UpdateEpisodeCollection_Call {
	return &CollectionRepo_UpdateEpisodeCollection_Call{Call: _e.mock.On("UpdateEpisodeCollection", ctx, userID, subjectID, episodeIDs, _a4, at)}
}

func (_c *CollectionRepo_UpdateEpisodeCollection_Call) Run(run func(ctx context.Context, userID uint32, subjectID uint32, episodeIDs []uint32, _a4 collection.EpisodeCollection, at time.Time)) *CollectionRepo_UpdateEpisodeCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(uint32), args[3].([]uint32), args[4].(collection.EpisodeCollection), args[5].(time.Time))
	})
	return _c
}

func (_c *CollectionRepo_UpdateEpisodeCollection_Call) Return(_a0 collection.UserSubjectEpisodesCollection, _a1 error) *CollectionRepo_UpdateEpisodeCollection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CollectionRepo_UpdateEpisodeCollection_Call) RunAndReturn(run func(context.Context, uint32, uint32, []uint32, collection.EpisodeCollection, time.Time) (collection.UserSubjectEpisodesCollection, error)) *CollectionRepo_UpdateEpisodeCollection_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateOrCreateSubjectCollection provides a mock function with given fields: ctx, userID, subject, at, ip, update
func (_m *CollectionRepo) UpdateOrCreateSubjectCollection(ctx context.Context, userID uint32, subject model.Subject, at time.Time, ip string, update func(context.Context, *collection.Subject) (*collection.Subject, error)) error {
	ret := _m.Called(ctx, userID, subject, at, ip, update)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrCreateSubjectCollection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.Subject, time.Time, string, func(context.Context, *collection.Subject) (*collection.Subject, error)) error); ok {
		r0 = rf(ctx, userID, subject, at, ip, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CollectionRepo_UpdateOrCreateSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateOrCreateSubjectCollection'
type CollectionRepo_UpdateOrCreateSubjectCollection_Call struct {
	*mock.Call
}

// UpdateOrCreateSubjectCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - subject model.Subject
//   - at time.Time
//   - ip string
//   - update func(context.Context , *collection.Subject)(*collection.Subject , error)
func (_e *CollectionRepo_Expecter) UpdateOrCreateSubjectCollection(ctx interface{}, userID interface{}, subject interface{}, at interface{}, ip interface{}, update interface{}) *CollectionRepo_UpdateOrCreateSubjectCollection_Call {
	return &CollectionRepo_UpdateOrCreateSubjectCollection_Call{Call: _e.mock.On("UpdateOrCreateSubjectCollection", ctx, userID, subject, at, ip, update)}
}

func (_c *CollectionRepo_UpdateOrCreateSubjectCollection_Call) Run(run func(ctx context.Context, userID uint32, subject model.Subject, at time.Time, ip string, update func(context.Context, *collection.Subject) (*collection.Subject, error))) *CollectionRepo_UpdateOrCreateSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(model.Subject), args[3].(time.Time), args[4].(string), args[5].(func(context.Context, *collection.Subject) (*collection.Subject, error)))
	})
	return _c
}

func (_c *CollectionRepo_UpdateOrCreateSubjectCollection_Call) Return(_a0 error) *CollectionRepo_UpdateOrCreateSubjectCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CollectionRepo_UpdateOrCreateSubjectCollection_Call) RunAndReturn(run func(context.Context, uint32, model.Subject, time.Time, string, func(context.Context, *collection.Subject) (*collection.Subject, error)) error) *CollectionRepo_UpdateOrCreateSubjectCollection_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateSubjectCollection provides a mock function with given fields: ctx, userID, subject, at, ip, update
func (_m *CollectionRepo) UpdateSubjectCollection(ctx context.Context, userID uint32, subject model.Subject, at time.Time, ip string, update func(context.Context, *collection.Subject) (*collection.Subject, error)) error {
	ret := _m.Called(ctx, userID, subject, at, ip, update)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSubjectCollection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uint32, model.Subject, time.Time, string, func(context.Context, *collection.Subject) (*collection.Subject, error)) error); ok {
		r0 = rf(ctx, userID, subject, at, ip, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CollectionRepo_UpdateSubjectCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateSubjectCollection'
type CollectionRepo_UpdateSubjectCollection_Call struct {
	*mock.Call
}

// UpdateSubjectCollection is a helper method to define mock.On call
//   - ctx context.Context
//   - userID uint32
//   - subject model.Subject
//   - at time.Time
//   - ip string
//   - update func(context.Context , *collection.Subject)(*collection.Subject , error)
func (_e *CollectionRepo_Expecter) UpdateSubjectCollection(ctx interface{}, userID interface{}, subject interface{}, at interface{}, ip interface{}, update interface{}) *CollectionRepo_UpdateSubjectCollection_Call {
	return &CollectionRepo_UpdateSubjectCollection_Call{Call: _e.mock.On("UpdateSubjectCollection", ctx, userID, subject, at, ip, update)}
}

func (_c *CollectionRepo_UpdateSubjectCollection_Call) Run(run func(ctx context.Context, userID uint32, subject model.Subject, at time.Time, ip string, update func(context.Context, *collection.Subject) (*collection.Subject, error))) *CollectionRepo_UpdateSubjectCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uint32), args[2].(model.Subject), args[3].(time.Time), args[4].(string), args[5].(func(context.Context, *collection.Subject) (*collection.Subject, error)))
	})
	return _c
}

func (_c *CollectionRepo_UpdateSubjectCollection_Call) Return(_a0 error) *CollectionRepo_UpdateSubjectCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CollectionRepo_UpdateSubjectCollection_Call) RunAndReturn(run func(context.Context, uint32, model.Subject, time.Time, string, func(context.Context, *collection.Subject) (*collection.Subject, error)) error) *CollectionRepo_UpdateSubjectCollection_Call {
	_c.Call.Return(run)
	return _c
}

// WithQuery provides a mock function with given fields: _a0
func (_m *CollectionRepo) WithQuery(_a0 *query.Query) collections.Repo {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for WithQuery")
	}

	var r0 collections.Repo
	if rf, ok := ret.Get(0).(func(*query.Query) collections.Repo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(collections.Repo)
		}
	}

	return r0
}

// CollectionRepo_WithQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithQuery'
type CollectionRepo_WithQuery_Call struct {
	*mock.Call
}

// WithQuery is a helper method to define mock.On call
//   - _a0 *query.Query
func (_e *CollectionRepo_Expecter) WithQuery(_a0 interface{}) *CollectionRepo_WithQuery_Call {
	return &CollectionRepo_WithQuery_Call{Call: _e.mock.On("WithQuery", _a0)}
}

func (_c *CollectionRepo_WithQuery_Call) Run(run func(_a0 *query.Query)) *CollectionRepo_WithQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*query.Query))
	})
	return _c
}

func (_c *CollectionRepo_WithQuery_Call) Return(_a0 collections.Repo) *CollectionRepo_WithQuery_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CollectionRepo_WithQuery_Call) RunAndReturn(run func(*query.Query) collections.Repo) *CollectionRepo_WithQuery_Call {
	_c.Call.Return(run)
	return _c
}

// NewCollectionRepo creates a new instance of CollectionRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCollectionRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *CollectionRepo {
	mock := &CollectionRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
