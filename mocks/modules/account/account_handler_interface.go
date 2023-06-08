// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// AccountHandlerInterface is an autogenerated mock type for the AccountHandlerInterface type
type AccountHandlerInterface struct {
	mock.Mock
}

type AccountHandlerInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *AccountHandlerInterface) EXPECT() *AccountHandlerInterface_Expecter {
	return &AccountHandlerInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: c
func (_m *AccountHandlerInterface) Create(c *gin.Context) {
	_m.Called(c)
}

// AccountHandlerInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type AccountHandlerInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - c *gin.Context
func (_e *AccountHandlerInterface_Expecter) Create(c interface{}) *AccountHandlerInterface_Create_Call {
	return &AccountHandlerInterface_Create_Call{Call: _e.mock.On("Create", c)}
}

func (_c *AccountHandlerInterface_Create_Call) Run(run func(c *gin.Context)) *AccountHandlerInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *AccountHandlerInterface_Create_Call) Return() *AccountHandlerInterface_Create_Call {
	_c.Call.Return()
	return _c
}

func (_c *AccountHandlerInterface_Create_Call) RunAndReturn(run func(*gin.Context)) *AccountHandlerInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: c
func (_m *AccountHandlerInterface) Login(c *gin.Context) {
	_m.Called(c)
}

// AccountHandlerInterface_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type AccountHandlerInterface_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - c *gin.Context
func (_e *AccountHandlerInterface_Expecter) Login(c interface{}) *AccountHandlerInterface_Login_Call {
	return &AccountHandlerInterface_Login_Call{Call: _e.mock.On("Login", c)}
}

func (_c *AccountHandlerInterface_Login_Call) Run(run func(c *gin.Context)) *AccountHandlerInterface_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *AccountHandlerInterface_Login_Call) Return() *AccountHandlerInterface_Login_Call {
	_c.Call.Return()
	return _c
}

func (_c *AccountHandlerInterface_Login_Call) RunAndReturn(run func(*gin.Context)) *AccountHandlerInterface_Login_Call {
	_c.Call.Return(run)
	return _c
}

// ReadByUsername provides a mock function with given fields: c
func (_m *AccountHandlerInterface) ReadByUsername(c *gin.Context) {
	_m.Called(c)
}

// AccountHandlerInterface_ReadByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadByUsername'
type AccountHandlerInterface_ReadByUsername_Call struct {
	*mock.Call
}

// ReadByUsername is a helper method to define mock.On call
//   - c *gin.Context
func (_e *AccountHandlerInterface_Expecter) ReadByUsername(c interface{}) *AccountHandlerInterface_ReadByUsername_Call {
	return &AccountHandlerInterface_ReadByUsername_Call{Call: _e.mock.On("ReadByUsername", c)}
}

func (_c *AccountHandlerInterface_ReadByUsername_Call) Run(run func(c *gin.Context)) *AccountHandlerInterface_ReadByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gin.Context))
	})
	return _c
}

func (_c *AccountHandlerInterface_ReadByUsername_Call) Return() *AccountHandlerInterface_ReadByUsername_Call {
	_c.Call.Return()
	return _c
}

func (_c *AccountHandlerInterface_ReadByUsername_Call) RunAndReturn(run func(*gin.Context)) *AccountHandlerInterface_ReadByUsername_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAccountHandlerInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewAccountHandlerInterface creates a new instance of AccountHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAccountHandlerInterface(t mockConstructorTestingTNewAccountHandlerInterface) *AccountHandlerInterface {
	mock := &AccountHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}