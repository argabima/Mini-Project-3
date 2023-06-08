// Code generated by mockery v2.20.0. DO NOT EDIT.

package customers

import mock "github.com/stretchr/testify/mock"

// MockRepositoryInterface is an autogenerated mock type for the RepositoryInterface type
type MockRepositoryInterface struct {
	mock.Mock
}

type MockRepositoryInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRepositoryInterface) EXPECT() *MockRepositoryInterface_Expecter {
	return &MockRepositoryInterface_Expecter{mock: &_m.Mock}
}

// DeleteByID provides a mock function with given fields: customers
func (_m *MockRepositoryInterface) DeleteByID(customers *Customers) error {
	ret := _m.Called(customers)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Customers) error); ok {
		r0 = rf(customers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepositoryInterface_DeleteByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByID'
type MockRepositoryInterface_DeleteByID_Call struct {
	*mock.Call
}

// DeleteByID is a helper method to define mock.On call
//   - customers *Customers
func (_e *MockRepositoryInterface_Expecter) DeleteByID(customers interface{}) *MockRepositoryInterface_DeleteByID_Call {
	return &MockRepositoryInterface_DeleteByID_Call{Call: _e.mock.On("DeleteByID", customers)}
}

func (_c *MockRepositoryInterface_DeleteByID_Call) Run(run func(customers *Customers)) *MockRepositoryInterface_DeleteByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Customers))
	})
	return _c
}

func (_c *MockRepositoryInterface_DeleteByID_Call) Return(_a0 error) *MockRepositoryInterface_DeleteByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepositoryInterface_DeleteByID_Call) RunAndReturn(run func(*Customers) error) *MockRepositoryInterface_DeleteByID_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *MockRepositoryInterface) FindAll() ([]Customers, error) {
	ret := _m.Called()

	var r0 []Customers
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]Customers, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []Customers); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]Customers)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockRepositoryInterface_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockRepositoryInterface_Expecter) FindAll() *MockRepositoryInterface_FindAll_Call {
	return &MockRepositoryInterface_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockRepositoryInterface_FindAll_Call) Run(run func()) *MockRepositoryInterface_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRepositoryInterface_FindAll_Call) Return(_a0 []Customers, _a1 error) *MockRepositoryInterface_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_FindAll_Call) RunAndReturn(run func() ([]Customers, error)) *MockRepositoryInterface_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// FindByEmail provides a mock function with given fields: email
func (_m *MockRepositoryInterface) FindByEmail(email string) (Customers, error) {
	ret := _m.Called(email)

	var r0 Customers
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (Customers, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) Customers); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(Customers)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_FindByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByEmail'
type MockRepositoryInterface_FindByEmail_Call struct {
	*mock.Call
}

// FindByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockRepositoryInterface_Expecter) FindByEmail(email interface{}) *MockRepositoryInterface_FindByEmail_Call {
	return &MockRepositoryInterface_FindByEmail_Call{Call: _e.mock.On("FindByEmail", email)}
}

func (_c *MockRepositoryInterface_FindByEmail_Call) Run(run func(email string)) *MockRepositoryInterface_FindByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockRepositoryInterface_FindByEmail_Call) Return(_a0 Customers, _a1 error) *MockRepositoryInterface_FindByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_FindByEmail_Call) RunAndReturn(run func(string) (Customers, error)) *MockRepositoryInterface_FindByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// FindById provides a mock function with given fields: id
func (_m *MockRepositoryInterface) FindById(id interface{}) (Customers, error) {
	ret := _m.Called(id)

	var r0 Customers
	var r1 error
	if rf, ok := ret.Get(0).(func(interface{}) (Customers, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(interface{}) Customers); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(Customers)
	}

	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockRepositoryInterface_FindById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindById'
type MockRepositoryInterface_FindById_Call struct {
	*mock.Call
}

// FindById is a helper method to define mock.On call
//   - id interface{}
func (_e *MockRepositoryInterface_Expecter) FindById(id interface{}) *MockRepositoryInterface_FindById_Call {
	return &MockRepositoryInterface_FindById_Call{Call: _e.mock.On("FindById", id)}
}

func (_c *MockRepositoryInterface_FindById_Call) Run(run func(id interface{})) *MockRepositoryInterface_FindById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockRepositoryInterface_FindById_Call) Return(_a0 Customers, _a1 error) *MockRepositoryInterface_FindById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockRepositoryInterface_FindById_Call) RunAndReturn(run func(interface{}) (Customers, error)) *MockRepositoryInterface_FindById_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: customers
func (_m *MockRepositoryInterface) Save(customers *Customers) error {
	ret := _m.Called(customers)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Customers) error); ok {
		r0 = rf(customers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepositoryInterface_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockRepositoryInterface_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - customers *Customers
func (_e *MockRepositoryInterface_Expecter) Save(customers interface{}) *MockRepositoryInterface_Save_Call {
	return &MockRepositoryInterface_Save_Call{Call: _e.mock.On("Save", customers)}
}

func (_c *MockRepositoryInterface_Save_Call) Run(run func(customers *Customers)) *MockRepositoryInterface_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Customers))
	})
	return _c
}

func (_c *MockRepositoryInterface_Save_Call) Return(_a0 error) *MockRepositoryInterface_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepositoryInterface_Save_Call) RunAndReturn(run func(*Customers) error) *MockRepositoryInterface_Save_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateByID provides a mock function with given fields: customers
func (_m *MockRepositoryInterface) UpdateByID(customers *Customers) error {
	ret := _m.Called(customers)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Customers) error); ok {
		r0 = rf(customers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRepositoryInterface_UpdateByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateByID'
type MockRepositoryInterface_UpdateByID_Call struct {
	*mock.Call
}

// UpdateByID is a helper method to define mock.On call
//   - customers *Customers
func (_e *MockRepositoryInterface_Expecter) UpdateByID(customers interface{}) *MockRepositoryInterface_UpdateByID_Call {
	return &MockRepositoryInterface_UpdateByID_Call{Call: _e.mock.On("UpdateByID", customers)}
}

func (_c *MockRepositoryInterface_UpdateByID_Call) Run(run func(customers *Customers)) *MockRepositoryInterface_UpdateByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Customers))
	})
	return _c
}

func (_c *MockRepositoryInterface_UpdateByID_Call) Return(_a0 error) *MockRepositoryInterface_UpdateByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRepositoryInterface_UpdateByID_Call) RunAndReturn(run func(*Customers) error) *MockRepositoryInterface_UpdateByID_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockRepositoryInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockRepositoryInterface creates a new instance of MockRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockRepositoryInterface(t mockConstructorTestingTNewMockRepositoryInterface) *MockRepositoryInterface {
	mock := &MockRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
