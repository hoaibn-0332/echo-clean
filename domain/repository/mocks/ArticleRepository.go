// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	entity "echo-clean/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// ArticleRepository is an autogenerated mock type for the ArticleRepository type
type ArticleRepository struct {
	mock.Mock
}

type ArticleRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *ArticleRepository) EXPECT() *ArticleRepository_Expecter {
	return &ArticleRepository_Expecter{mock: &_m.Mock}
}

// Delete provides a mock function with given fields: id
func (_m *ArticleRepository) Delete(id int64) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ArticleRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ArticleRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id int64
func (_e *ArticleRepository_Expecter) Delete(id interface{}) *ArticleRepository_Delete_Call {
	return &ArticleRepository_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *ArticleRepository_Delete_Call) Run(run func(id int64)) *ArticleRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *ArticleRepository_Delete_Call) Return(_a0 error) *ArticleRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ArticleRepository_Delete_Call) RunAndReturn(run func(int64) error) *ArticleRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// Fetch provides a mock function with given fields:
func (_m *ArticleRepository) Fetch() ([]*entity.Article, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []*entity.Article
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entity.Article, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entity.Article); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ArticleRepository_Fetch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fetch'
type ArticleRepository_Fetch_Call struct {
	*mock.Call
}

// Fetch is a helper method to define mock.On call
func (_e *ArticleRepository_Expecter) Fetch() *ArticleRepository_Fetch_Call {
	return &ArticleRepository_Fetch_Call{Call: _e.mock.On("Fetch")}
}

func (_c *ArticleRepository_Fetch_Call) Run(run func()) *ArticleRepository_Fetch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ArticleRepository_Fetch_Call) Return(_a0 []*entity.Article, _a1 error) *ArticleRepository_Fetch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ArticleRepository_Fetch_Call) RunAndReturn(run func() ([]*entity.Article, error)) *ArticleRepository_Fetch_Call {
	_c.Call.Return(run)
	return _c
}

// GetByID provides a mock function with given fields: id
func (_m *ArticleRepository) GetByID(id int64) (*entity.Article, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *entity.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (*entity.Article, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) *entity.Article); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ArticleRepository_GetByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByID'
type ArticleRepository_GetByID_Call struct {
	*mock.Call
}

// GetByID is a helper method to define mock.On call
//   - id int64
func (_e *ArticleRepository_Expecter) GetByID(id interface{}) *ArticleRepository_GetByID_Call {
	return &ArticleRepository_GetByID_Call{Call: _e.mock.On("GetByID", id)}
}

func (_c *ArticleRepository_GetByID_Call) Run(run func(id int64)) *ArticleRepository_GetByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *ArticleRepository_GetByID_Call) Return(_a0 *entity.Article, _a1 error) *ArticleRepository_GetByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ArticleRepository_GetByID_Call) RunAndReturn(run func(int64) (*entity.Article, error)) *ArticleRepository_GetByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetByTitle provides a mock function with given fields: title
func (_m *ArticleRepository) GetByTitle(title string) (*entity.Article, error) {
	ret := _m.Called(title)

	if len(ret) == 0 {
		panic("no return value specified for GetByTitle")
	}

	var r0 *entity.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.Article, error)); ok {
		return rf(title)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.Article); ok {
		r0 = rf(title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ArticleRepository_GetByTitle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByTitle'
type ArticleRepository_GetByTitle_Call struct {
	*mock.Call
}

// GetByTitle is a helper method to define mock.On call
//   - title string
func (_e *ArticleRepository_Expecter) GetByTitle(title interface{}) *ArticleRepository_GetByTitle_Call {
	return &ArticleRepository_GetByTitle_Call{Call: _e.mock.On("GetByTitle", title)}
}

func (_c *ArticleRepository_GetByTitle_Call) Run(run func(title string)) *ArticleRepository_GetByTitle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ArticleRepository_GetByTitle_Call) Return(_a0 *entity.Article, _a1 error) *ArticleRepository_GetByTitle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ArticleRepository_GetByTitle_Call) RunAndReturn(run func(string) (*entity.Article, error)) *ArticleRepository_GetByTitle_Call {
	_c.Call.Return(run)
	return _c
}

// Store provides a mock function with given fields: article, authorId
func (_m *ArticleRepository) Store(article *entity.Article, authorId int64) (*entity.Article, error) {
	ret := _m.Called(article, authorId)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 *entity.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Article, int64) (*entity.Article, error)); ok {
		return rf(article, authorId)
	}
	if rf, ok := ret.Get(0).(func(*entity.Article, int64) *entity.Article); ok {
		r0 = rf(article, authorId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Article, int64) error); ok {
		r1 = rf(article, authorId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ArticleRepository_Store_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Store'
type ArticleRepository_Store_Call struct {
	*mock.Call
}

// Store is a helper method to define mock.On call
//   - article *entity.Article
//   - authorId int64
func (_e *ArticleRepository_Expecter) Store(article interface{}, authorId interface{}) *ArticleRepository_Store_Call {
	return &ArticleRepository_Store_Call{Call: _e.mock.On("Store", article, authorId)}
}

func (_c *ArticleRepository_Store_Call) Run(run func(article *entity.Article, authorId int64)) *ArticleRepository_Store_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Article), args[1].(int64))
	})
	return _c
}

func (_c *ArticleRepository_Store_Call) Return(_a0 *entity.Article, _a1 error) *ArticleRepository_Store_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ArticleRepository_Store_Call) RunAndReturn(run func(*entity.Article, int64) (*entity.Article, error)) *ArticleRepository_Store_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: article
func (_m *ArticleRepository) Update(article *entity.Article) (*entity.Article, error) {
	ret := _m.Called(article)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *entity.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(*entity.Article) (*entity.Article, error)); ok {
		return rf(article)
	}
	if rf, ok := ret.Get(0).(func(*entity.Article) *entity.Article); ok {
		r0 = rf(article)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Article) error); ok {
		r1 = rf(article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ArticleRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ArticleRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - article *entity.Article
func (_e *ArticleRepository_Expecter) Update(article interface{}) *ArticleRepository_Update_Call {
	return &ArticleRepository_Update_Call{Call: _e.mock.On("Update", article)}
}

func (_c *ArticleRepository_Update_Call) Run(run func(article *entity.Article)) *ArticleRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Article))
	})
	return _c
}

func (_c *ArticleRepository_Update_Call) Return(_a0 *entity.Article, _a1 error) *ArticleRepository_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ArticleRepository_Update_Call) RunAndReturn(run func(*entity.Article) (*entity.Article, error)) *ArticleRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewArticleRepository creates a new instance of ArticleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArticleRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArticleRepository {
	mock := &ArticleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
