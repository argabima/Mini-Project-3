// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	account "crm_service/modules/account"

	mock "github.com/stretchr/testify/mock"
)

// AccountRepoInterface is an autogenerated mock type for the AccountRepoInterface type
type AccountRepoInterface struct {
	mock.Mock
}

type AccountRepoInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountRepoInterface) EXPECT() *AccountRepoInterface_Expecter {
	return &AccountRepoInterface_Expecter{mock: &_m.Mock}
}

// FindByUsername provides a mock function with given fields: username
func (_m *AccountRepoInterface) FindByUsername(username string) (account.Actors, error) {
	ret := _m.Called(username)

	var r0 account.Actors
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (account.Actors, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) account.Actors); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(account.Actors)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AccountRepoInterface_FindByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByUsername'
type AccountRepoInterface_FindByUsername_Call struct {
	*mock.Call
}

// FindByUsername is a helper method to define mock.On call
//   - username string
func (_e *AccountRepoInterface_Expecter) FindByUsername(username interface{}) *AccountRepoInterface_FindByUsername_Call {
	return &AccountRepoInterface_FindByUsername_Call{Call: _e.mock.On("FindByUsername", username)}
}

func (_c *AccountRepoInterface_FindByUsername_Call) Run(run func(username string)) *AccountRepoInterface_FindByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *AccountRepoInterface_FindByUsername_Call) Return(_a0 account.Actors, _a1 error) *AccountRepoInterface_FindByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *AccountRepoInterface_FindByUsername_Call) RunAndReturn(run func(string) (account.Actors, error)) *AccountRepoInterface_FindByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: _a0
func (_m *AccountRepoInterface) Save(_a0 *account.Actors) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Actors) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountRepoInterface_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type AccountRepoInterface_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - _a0 *account.Actors
func (_e *AccountRepoInterface_Expecter) Save(_a0 interface{}) *AccountRepoInterface_Save_Call {
	return &AccountRepoInterface_Save_Call{Call: _e.mock.On("Save", _a0)}
}

func (_c *AccountRepoInterface_Save_Call) Run(run func(_a0 *account.Actors)) *AccountRepoInterface_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*account.Actors))
	})
	return _c
}

func (_c *AccountRepoInterface_Save_Call) Return(_a0 error) *AccountRepoInterface_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *AccountRepoInterface_Save_Call) RunAndReturn(run func(*account.Actors) error) *AccountRepoInterface_Save_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAccountRepoInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountRepoInterface creates a new instance of AccountRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountRepoInterface(t mockConstructorTestingTNewAccountRepoInterface) *AccountRepoInterface {
	mock := &AccountRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
