// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	entity "echo-clean/domain/entity"
	error "echo-clean/domain/error"

	mock "github.com/stretchr/testify/mock"
)

// ArticleService is an autogenerated mock type for the ArticleService type
type ArticleService struct {
	mock.Mock
}

type ArticleService_Expecter struct {
	mock *mock.Mock
}

func (_m *ArticleService) EXPECT() *ArticleService_Expecter {
	return &ArticleService_Expecter{mock: &_m.Mock}
}

// Fetch provides a mock function with given fields:
func (_m *ArticleService) Fetch() ([]*entity.Article, []error.Error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Fetch")
	}

	var r0 []*entity.Article
	var r1 []error.Error
	if rf, ok := ret.Get(0).(func() ([]*entity.Article, []error.Error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entity.Article); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func() []error.Error); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error.Error)
		}
	}

	return r0, r1
}

// ArticleService_Fetch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Fetch'
type ArticleService_Fetch_Call struct {
	*mock.Call
}

// Fetch is a helper method to define mock.On call
func (_e *ArticleService_Expecter) Fetch() *ArticleService_Fetch_Call {
	return &ArticleService_Fetch_Call{Call: _e.mock.On("Fetch")}
}

func (_c *ArticleService_Fetch_Call) Run(run func()) *ArticleService_Fetch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ArticleService_Fetch_Call) Return(_a0 []*entity.Article, _a1 []error.Error) *ArticleService_Fetch_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ArticleService_Fetch_Call) RunAndReturn(run func() ([]*entity.Article, []error.Error)) *ArticleService_Fetch_Call {
	_c.Call.Return(run)
	return _c
}

// Store provides a mock function with given fields: article
func (_m *ArticleService) Store(article *entity.Article) (*entity.Article, []error.Error) {
	ret := _m.Called(article)

	if len(ret) == 0 {
		panic("no return value specified for Store")
	}

	var r0 *entity.Article
	var r1 []error.Error
	if rf, ok := ret.Get(0).(func(*entity.Article) (*entity.Article, []error.Error)); ok {
		return rf(article)
	}
	if rf, ok := ret.Get(0).(func(*entity.Article) *entity.Article); ok {
		r0 = rf(article)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Article) []error.Error); ok {
		r1 = rf(article)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]error.Error)
		}
	}

	return r0, r1
}

// ArticleService_Store_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Store'
type ArticleService_Store_Call struct {
	*mock.Call
}

// Store is a helper method to define mock.On call
//   - article *entity.Article
func (_e *ArticleService_Expecter) Store(article interface{}) *ArticleService_Store_Call {
	return &ArticleService_Store_Call{Call: _e.mock.On("Store", article)}
}

func (_c *ArticleService_Store_Call) Run(run func(article *entity.Article)) *ArticleService_Store_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Article))
	})
	return _c
}

func (_c *ArticleService_Store_Call) Return(_a0 *entity.Article, _a1 []error.Error) *ArticleService_Store_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ArticleService_Store_Call) RunAndReturn(run func(*entity.Article) (*entity.Article, []error.Error)) *ArticleService_Store_Call {
	_c.Call.Return(run)
	return _c
}

// NewArticleService creates a new instance of ArticleService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewArticleService(t interface {
	mock.TestingT
	Cleanup(func())
}) *ArticleService {
	mock := &ArticleService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}