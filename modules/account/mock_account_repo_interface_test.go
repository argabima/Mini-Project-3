// Code generated by mockery v2.20.0. DO NOT EDIT.

package account

import mock "github.com/stretchr/testify/mock"

// MockAccountRepoInterface is an autogenerated mock type for the AccountRepoInterface type
type MockAccountRepoInterface struct {
	mock.Mock
}

type MockAccountRepoInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAccountRepoInterface) EXPECT() *MockAccountRepoInterface_Expecter {
	return &MockAccountRepoInterface_Expecter{mock: &_m.Mock}
}

// FindAll provides a mock function with given fields:
func (_m *MockAccountRepoInterface) FindAll() ([]Actors, error) {
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

// MockAccountRepoInterface_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockAccountRepoInterface_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockAccountRepoInterface_Expecter) FindAll() *MockAccountRepoInterface_FindAll_Call {
	return &MockAccountRepoInterface_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockAccountRepoInterface_FindAll_Call) Run(run func()) *MockAccountRepoInterface_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockAccountRepoInterface_FindAll_Call) Return(_a0 []Actors, _a1 error) *MockAccountRepoInterface_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepoInterface_FindAll_Call) RunAndReturn(run func() ([]Actors, error)) *MockAccountRepoInterface_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindByID provides a mock function with given fields: id
func (_m *MockAccountRepoInterface) FindByID(id interface{}) (Actors, error) {
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

// MockAccountRepoInterface_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type MockAccountRepoInterface_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - id interface{}
func (_e *MockAccountRepoInterface_Expecter) FindByID(id interface{}) *MockAccountRepoInterface_FindByID_Call {
	return &MockAccountRepoInterface_FindByID_Call{Call: _e.mock.On("FindByID", id)}
}

func (_c *MockAccountRepoInterface_FindByID_Call) Run(run func(id interface{})) *MockAccountRepoInterface_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockAccountRepoInterface_FindByID_Call) Return(_a0 Actors, _a1 error) *MockAccountRepoInterface_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepoInterface_FindByID_Call) RunAndReturn(run func(interface{}) (Actors, error)) *MockAccountRepoInterface_FindByID_Call {
	_c.Call.Return(run)
	return _c
}

// FindByUsername provides a mock function with given fields: username
func (_m *MockAccountRepoInterface) FindByUsername(username string) (Actors, error) {
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

// MockAccountRepoInterface_FindByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByUsername'
type MockAccountRepoInterface_FindByUsername_Call struct {
	*mock.Call
}

// FindByUsername is a helper method to define mock.On call
//   - username string
func (_e *MockAccountRepoInterface_Expecter) FindByUsername(username interface{}) *MockAccountRepoInterface_FindByUsername_Call {
	return &MockAccountRepoInterface_FindByUsername_Call{Call: _e.mock.On("FindByUsername", username)}
}

func (_c *MockAccountRepoInterface_FindByUsername_Call) Run(run func(username string)) *MockAccountRepoInterface_FindByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockAccountRepoInterface_FindByUsername_Call) Return(_a0 Actors, _a1 error) *MockAccountRepoInterface_FindByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockAccountRepoInterface_FindByUsername_Call) RunAndReturn(run func(string) (Actors, error)) *MockAccountRepoInterface_FindByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: account
func (_m *MockAccountRepoInterface) Save(account *Actors) error {
	ret := _m.Called(account)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Actors) error); ok {
		r0 = rf(account)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountRepoInterface_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockAccountRepoInterface_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - account *Actors
func (_e *MockAccountRepoInterface_Expecter) Save(account interface{}) *MockAccountRepoInterface_Save_Call {
	return &MockAccountRepoInterface_Save_Call{Call: _e.mock.On("Save", account)}
}

func (_c *MockAccountRepoInterface_Save_Call) Run(run func(account *Actors)) *MockAccountRepoInterface_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Actors))
	})
	return _c
}

func (_c *MockAccountRepoInterface_Save_Call) Return(_a0 error) *MockAccountRepoInterface_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountRepoInterface_Save_Call) RunAndReturn(run func(*Actors) error) *MockAccountRepoInterface_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: actor
func (_m *MockAccountRepoInterface) Update(actor *Actors) error {
	ret := _m.Called(actor)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Actors) error); ok {
		r0 = rf(actor)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAccountRepoInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockAccountRepoInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - actor *Actors
func (_e *MockAccountRepoInterface_Expecter) Update(actor interface{}) *MockAccountRepoInterface_Update_Call {
	return &MockAccountRepoInterface_Update_Call{Call: _e.mock.On("Update", actor)}
}

func (_c *MockAccountRepoInterface_Update_Call) Run(run func(actor *Actors)) *MockAccountRepoInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Actors))
	})
	return _c
}

func (_c *MockAccountRepoInterface_Update_Call) Return(_a0 error) *MockAccountRepoInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAccountRepoInterface_Update_Call) RunAndReturn(run func(*Actors) error) *MockAccountRepoInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockAccountRepoInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockAccountRepoInterface creates a new instance of MockAccountRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockAccountRepoInterface(t mockConstructorTestingTNewMockAccountRepoInterface) *MockAccountRepoInterface {
	mock := &MockAccountRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
