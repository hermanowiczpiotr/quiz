// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	entity "quiz/internal/domain/entity"

	mock "github.com/stretchr/testify/mock"
)

// MockUserGameRepository is an autogenerated mock type for the UserGameRepository type
type MockUserGameRepository struct {
	mock.Mock
}

type MockUserGameRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserGameRepository) EXPECT() *MockUserGameRepository_Expecter {
	return &MockUserGameRepository_Expecter{mock: &_m.Mock}
}

// FindAll provides a mock function with given fields:
func (_m *MockUserGameRepository) FindAll() map[string]*entity.UserGame {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 map[string]*entity.UserGame
	if rf, ok := ret.Get(0).(func() map[string]*entity.UserGame); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]*entity.UserGame)
		}
	}

	return r0
}

// MockUserGameRepository_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockUserGameRepository_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockUserGameRepository_Expecter) FindAll() *MockUserGameRepository_FindAll_Call {
	return &MockUserGameRepository_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockUserGameRepository_FindAll_Call) Run(run func()) *MockUserGameRepository_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUserGameRepository_FindAll_Call) Return(_a0 map[string]*entity.UserGame) *MockUserGameRepository_FindAll_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserGameRepository_FindAll_Call) RunAndReturn(run func() map[string]*entity.UserGame) *MockUserGameRepository_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// GetByUsername provides a mock function with given fields: username
func (_m *MockUserGameRepository) GetByUsername(username string) (*entity.UserGame, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetByUsername")
	}

	var r0 *entity.UserGame
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entity.UserGame, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *entity.UserGame); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.UserGame)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUserGameRepository_GetByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByUsername'
type MockUserGameRepository_GetByUsername_Call struct {
	*mock.Call
}

// GetByUsername is a helper method to define mock.On call
//   - username string
func (_e *MockUserGameRepository_Expecter) GetByUsername(username interface{}) *MockUserGameRepository_GetByUsername_Call {
	return &MockUserGameRepository_GetByUsername_Call{Call: _e.mock.On("GetByUsername", username)}
}

func (_c *MockUserGameRepository_GetByUsername_Call) Run(run func(username string)) *MockUserGameRepository_GetByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUserGameRepository_GetByUsername_Call) Return(_a0 *entity.UserGame, _a1 error) *MockUserGameRepository_GetByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUserGameRepository_GetByUsername_Call) RunAndReturn(run func(string) (*entity.UserGame, error)) *MockUserGameRepository_GetByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: userGame
func (_m *MockUserGameRepository) Save(userGame *entity.UserGame) {
	_m.Called(userGame)
}

// MockUserGameRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockUserGameRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - userGame *entity.UserGame
func (_e *MockUserGameRepository_Expecter) Save(userGame interface{}) *MockUserGameRepository_Save_Call {
	return &MockUserGameRepository_Save_Call{Call: _e.mock.On("Save", userGame)}
}

func (_c *MockUserGameRepository_Save_Call) Run(run func(userGame *entity.UserGame)) *MockUserGameRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.UserGame))
	})
	return _c
}

func (_c *MockUserGameRepository_Save_Call) Return() *MockUserGameRepository_Save_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockUserGameRepository_Save_Call) RunAndReturn(run func(*entity.UserGame)) *MockUserGameRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserGameRepository creates a new instance of MockUserGameRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserGameRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserGameRepository {
	mock := &MockUserGameRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
