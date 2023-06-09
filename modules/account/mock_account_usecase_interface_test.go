// Code generated by mockery v2.20.0. DO NOT EDIT.

package account

import mock "github.com/stretchr/testify/mock"

// MockAccountUsecaseInterface is an autogenerated mock type for the AccountUsecaseInterface type
type MockAccountUsecaseInterface struct {
	mock.Mock
}

type MockAccountUsecaseInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountUsecaseInterface) EXPECT() *MockAccountUsecaseInterface_Expecter {
	return &MockAccountUsecaseInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: account
func (_m *MockAccountUsecaseInterface) Create(account *Actors) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Actors) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountUsecaseInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockAccountUsecaseInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - account *Actors
func (_e *MockAccountUsecaseInterface_Expecter) Create(account interface{}) *MockAccountUsecaseInterface_Create_Call {
	return &MockAccountUsecaseInterface_Create_Call{Call: _e.mock.On("Create", account)}
}

func (_c *MockAccountUsecaseInterface_Create_Call) Run(run func(account *Actors)) *MockAccountUsecaseInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Actors))
	})
	return _c
}

func (_c *MockAccountUsecaseInterface_Create_Call) Return(_a0 error) *MockAccountUsecaseInterface_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsecaseInterface_Create_Call) RunAndReturn(run func(*Actors) error) *MockAccountUsecaseInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *MockAccountUsecaseInterface) FindAll() ([]Actors, error) {
	ret := _m.Called()

	var r0 []Actors
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Actors, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Actors); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Actors)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountUsecaseInterface_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockAccountUsecaseInterface_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockAccountUsecaseInterface_Expecter) FindAll() *MockAccountUsecaseInterface_FindAll_Call {
	return &MockAccountUsecaseInterface_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockAccountUsecaseInterface_FindAll_Call) Run(run func()) *MockAccountUsecaseInterface_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountUsecaseInterface_FindAll_Call) Return(_a0 []Actors, _a1 error) *MockAccountUsecaseInterface_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsecaseInterface_FindAll_Call) RunAndReturn(run func() ([]Actors, error)) *MockAccountUsecaseInterface_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindByID provides a mock function with given fields: id
func (_m *MockAccountUsecaseInterface) FindByID(id interface{}) (Actors, error) {
	ret := _m.Called(id)

	var r0 Actors
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (Actors, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(interface{}) Actors); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(Actors)
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountUsecaseInterface_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type MockAccountUsecaseInterface_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - id interface{}
func (_e *MockAccountUsecaseInterface_Expecter) FindByID(id interface{}) *MockAccountUsecaseInterface_FindByID_Call {
	return &MockAccountUsecaseInterface_FindByID_Call{Call: _e.mock.On("FindByID", id)}
}

func (_c *MockAccountUsecaseInterface_FindByID_Call) Run(run func(id interface{})) *MockAccountUsecaseInterface_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockAccountUsecaseInterface_FindByID_Call) Return(_a0 Actors, _a1 error) *MockAccountUsecaseInterface_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsecaseInterface_FindByID_Call) RunAndReturn(run func(interface{}) (Actors, error)) *MockAccountUsecaseInterface_FindByID_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: actor
func (_m *MockAccountUsecaseInterface) Update(actor *Actors) error {
	ret := _m.Called(actor)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Actors) error); ok {
		r0 = rf(actor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountUsecaseInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountUsecaseInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - actor *Actors
func (_e *MockAccountUsecaseInterface_Expecter) Update(actor interface{}) *MockAccountUsecaseInterface_Update_Call {
	return &MockAccountUsecaseInterface_Update_Call{Call: _e.mock.On("Update", actor)}
}

func (_c *MockAccountUsecaseInterface_Update_Call) Run(run func(actor *Actors)) *MockAccountUsecaseInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Actors))
	})
	return _c
}

func (_c *MockAccountUsecaseInterface_Update_Call) Return(_a0 error) *MockAccountUsecaseInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountUsecaseInterface_Update_Call) RunAndReturn(run func(*Actors) error) *MockAccountUsecaseInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// getByUsername provides a mock function with given fields: username
func (_m *MockAccountUsecaseInterface) getByUsername(username string) (Actors, error) {
	ret := _m.Called(username)

	var r0 Actors
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (Actors, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) Actors); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(Actors)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockAccountUsecaseInterface_getByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getByUsername'
type MockAccountUsecaseInterface_getByUsername_Call struct {
	*mock.Call
}

// getByUsername is a helper method to define mock.On call
//   - username string
func (_e *MockAccountUsecaseInterface_Expecter) getByUsername(username interface{}) *MockAccountUsecaseInterface_getByUsername_Call {
	return &MockAccountUsecaseInterface_getByUsername_Call{Call: _e.mock.On("getByUsername", username)}
}

func (_c *MockAccountUsecaseInterface_getByUsername_Call) Run(run func(username string)) *MockAccountUsecaseInterface_getByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAccountUsecaseInterface_getByUsername_Call) Return(_a0 Actors, _a1 error) *MockAccountUsecaseInterface_getByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountUsecaseInterface_getByUsername_Call) RunAndReturn(run func(string) (Actors, error)) *MockAccountUsecaseInterface_getByUsername_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAccountUsecaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAccountUsecaseInterface creates a new instance of MockAccountUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAccountUsecaseInterface(t mockConstructorTestingTNewMockAccountUsecaseInterface) *MockAccountUsecaseInterface {
	mock := &MockAccountUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
