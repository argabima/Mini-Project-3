// Code generated by mockery v2.20.0. DO NOT EDIT.

package customers

import mock "github.com/stretchr/testify/mock"

// MockUsecaseInterface is an autogenerated mock type for the UsecaseInterface type
type MockUsecaseInterface struct {
	mock.Mock
}

type MockUsecaseInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUsecaseInterface) EXPECT() *MockUsecaseInterface_Expecter {
	return &MockUsecaseInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: customers
func (_m *MockUsecaseInterface) Create(customers *Customers) error {
	ret := _m.Called(customers)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Customers) error); ok {
		r0 = rf(customers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsecaseInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockUsecaseInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - customers *Customers
func (_e *MockUsecaseInterface_Expecter) Create(customers interface{}) *MockUsecaseInterface_Create_Call {
	return &MockUsecaseInterface_Create_Call{Call: _e.mock.On("Create", customers)}
}

func (_c *MockUsecaseInterface_Create_Call) Run(run func(customers *Customers)) *MockUsecaseInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Customers))
	})
	return _c
}

func (_c *MockUsecaseInterface_Create_Call) Return(_a0 error) *MockUsecaseInterface_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsecaseInterface_Create_Call) RunAndReturn(run func(*Customers) error) *MockUsecaseInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByID provides a mock function with given fields: customers
func (_m *MockUsecaseInterface) DeleteByID(customers *Customers) error {
	ret := _m.Called(customers)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Customers) error); ok {
		r0 = rf(customers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsecaseInterface_DeleteByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByID'
type MockUsecaseInterface_DeleteByID_Call struct {
	*mock.Call
}

// DeleteByID is a helper method to define mock.On call
//   - customers *Customers
func (_e *MockUsecaseInterface_Expecter) DeleteByID(customers interface{}) *MockUsecaseInterface_DeleteByID_Call {
	return &MockUsecaseInterface_DeleteByID_Call{Call: _e.mock.On("DeleteByID", customers)}
}

func (_c *MockUsecaseInterface_DeleteByID_Call) Run(run func(customers *Customers)) *MockUsecaseInterface_DeleteByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Customers))
	})
	return _c
}

func (_c *MockUsecaseInterface_DeleteByID_Call) Return(_a0 error) *MockUsecaseInterface_DeleteByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsecaseInterface_DeleteByID_Call) RunAndReturn(run func(*Customers) error) *MockUsecaseInterface_DeleteByID_Call {
	_c.Call.Return(run)
	return _c
}

// FindAll provides a mock function with given fields:
func (_m *MockUsecaseInterface) FindAll() ([]Customers, error) {
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

// MockUsecaseInterface_FindAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindAll'
type MockUsecaseInterface_FindAll_Call struct {
	*mock.Call
}

// FindAll is a helper method to define mock.On call
func (_e *MockUsecaseInterface_Expecter) FindAll() *MockUsecaseInterface_FindAll_Call {
	return &MockUsecaseInterface_FindAll_Call{Call: _e.mock.On("FindAll")}
}

func (_c *MockUsecaseInterface_FindAll_Call) Run(run func()) *MockUsecaseInterface_FindAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUsecaseInterface_FindAll_Call) Return(_a0 []Customers, _a1 error) *MockUsecaseInterface_FindAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsecaseInterface_FindAll_Call) RunAndReturn(run func() ([]Customers, error)) *MockUsecaseInterface_FindAll_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields:
func (_m *MockUsecaseInterface) Read() ([]Customers, error) {
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

// MockUsecaseInterface_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type MockUsecaseInterface_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
func (_e *MockUsecaseInterface_Expecter) Read() *MockUsecaseInterface_Read_Call {
	return &MockUsecaseInterface_Read_Call{Call: _e.mock.On("Read")}
}

func (_c *MockUsecaseInterface_Read_Call) Run(run func()) *MockUsecaseInterface_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockUsecaseInterface_Read_Call) Return(_a0 []Customers, _a1 error) *MockUsecaseInterface_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsecaseInterface_Read_Call) RunAndReturn(run func() ([]Customers, error)) *MockUsecaseInterface_Read_Call {
	_c.Call.Return(run)
	return _c
}

// ReadByID provides a mock function with given fields: id
func (_m *MockUsecaseInterface) ReadByID(id interface{}) (Customers, error) {
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

// MockUsecaseInterface_ReadByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadByID'
type MockUsecaseInterface_ReadByID_Call struct {
	*mock.Call
}

// ReadByID is a helper method to define mock.On call
//   - id interface{}
func (_e *MockUsecaseInterface_Expecter) ReadByID(id interface{}) *MockUsecaseInterface_ReadByID_Call {
	return &MockUsecaseInterface_ReadByID_Call{Call: _e.mock.On("ReadByID", id)}
}

func (_c *MockUsecaseInterface_ReadByID_Call) Run(run func(id interface{})) *MockUsecaseInterface_ReadByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(interface{}))
	})
	return _c
}

func (_c *MockUsecaseInterface_ReadByID_Call) Return(_a0 Customers, _a1 error) *MockUsecaseInterface_ReadByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsecaseInterface_ReadByID_Call) RunAndReturn(run func(interface{}) (Customers, error)) *MockUsecaseInterface_ReadByID_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: customers
func (_m *MockUsecaseInterface) Save(customers *Customers) error {
	ret := _m.Called(customers)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Customers) error); ok {
		r0 = rf(customers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsecaseInterface_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockUsecaseInterface_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - customers *Customers
func (_e *MockUsecaseInterface_Expecter) Save(customers interface{}) *MockUsecaseInterface_Save_Call {
	return &MockUsecaseInterface_Save_Call{Call: _e.mock.On("Save", customers)}
}

func (_c *MockUsecaseInterface_Save_Call) Run(run func(customers *Customers)) *MockUsecaseInterface_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Customers))
	})
	return _c
}

func (_c *MockUsecaseInterface_Save_Call) Return(_a0 error) *MockUsecaseInterface_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsecaseInterface_Save_Call) RunAndReturn(run func(*Customers) error) *MockUsecaseInterface_Save_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateByID provides a mock function with given fields: customers
func (_m *MockUsecaseInterface) UpdateByID(customers *Customers) error {
	ret := _m.Called(customers)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Customers) error); ok {
		r0 = rf(customers)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUsecaseInterface_UpdateByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateByID'
type MockUsecaseInterface_UpdateByID_Call struct {
	*mock.Call
}

// UpdateByID is a helper method to define mock.On call
//   - customers *Customers
func (_e *MockUsecaseInterface_Expecter) UpdateByID(customers interface{}) *MockUsecaseInterface_UpdateByID_Call {
	return &MockUsecaseInterface_UpdateByID_Call{Call: _e.mock.On("UpdateByID", customers)}
}

func (_c *MockUsecaseInterface_UpdateByID_Call) Run(run func(customers *Customers)) *MockUsecaseInterface_UpdateByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*Customers))
	})
	return _c
}

func (_c *MockUsecaseInterface_UpdateByID_Call) Return(_a0 error) *MockUsecaseInterface_UpdateByID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUsecaseInterface_UpdateByID_Call) RunAndReturn(run func(*Customers) error) *MockUsecaseInterface_UpdateByID_Call {
	_c.Call.Return(run)
	return _c
}

// getByEmail provides a mock function with given fields: email
func (_m *MockUsecaseInterface) getByEmail(email string) (Customers, error) {
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

// MockUsecaseInterface_getByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'getByEmail'
type MockUsecaseInterface_getByEmail_Call struct {
	*mock.Call
}

// getByEmail is a helper method to define mock.On call
//   - email string
func (_e *MockUsecaseInterface_Expecter) getByEmail(email interface{}) *MockUsecaseInterface_getByEmail_Call {
	return &MockUsecaseInterface_getByEmail_Call{Call: _e.mock.On("getByEmail", email)}
}

func (_c *MockUsecaseInterface_getByEmail_Call) Run(run func(email string)) *MockUsecaseInterface_getByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockUsecaseInterface_getByEmail_Call) Return(_a0 Customers, _a1 error) *MockUsecaseInterface_getByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUsecaseInterface_getByEmail_Call) RunAndReturn(run func(string) (Customers, error)) *MockUsecaseInterface_getByEmail_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockUsecaseInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUsecaseInterface creates a new instance of MockUsecaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUsecaseInterface(t mockConstructorTestingTNewMockUsecaseInterface) *MockUsecaseInterface {
	mock := &MockUsecaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}