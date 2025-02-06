// Code generated by mockery v2.52.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "finance-operations-service/internal/finance/models"

	types "finance-operations-service/internal/finance/types"

	uuid "github.com/google/uuid"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

type Repository_Expecter struct {
	mock *mock.Mock
}

func (_m *Repository) EXPECT() *Repository_Expecter {
	return &Repository_Expecter{mock: &_m.Mock}
}

// CreateOperation provides a mock function with given fields: _a0, _a1, _a2
func (_m *Repository) CreateOperation(_a0 context.Context, _a1 *models.DetailsOperation, _a2 types.OperationType) (uuid.UUID, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CreateOperation")
	}

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.DetailsOperation, types.OperationType) (uuid.UUID, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.DetailsOperation, types.OperationType) uuid.UUID); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.DetailsOperation, types.OperationType) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_CreateOperation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOperation'
type Repository_CreateOperation_Call struct {
	*mock.Call
}

// CreateOperation is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.DetailsOperation
//   - _a2 types.OperationType
func (_e *Repository_Expecter) CreateOperation(_a0 interface{}, _a1 interface{}, _a2 interface{}) *Repository_CreateOperation_Call {
	return &Repository_CreateOperation_Call{Call: _e.mock.On("CreateOperation", _a0, _a1, _a2)}
}

func (_c *Repository_CreateOperation_Call) Run(run func(_a0 context.Context, _a1 *models.DetailsOperation, _a2 types.OperationType)) *Repository_CreateOperation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.DetailsOperation), args[2].(types.OperationType))
	})
	return _c
}

func (_c *Repository_CreateOperation_Call) Return(_a0 uuid.UUID, _a1 error) *Repository_CreateOperation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_CreateOperation_Call) RunAndReturn(run func(context.Context, *models.DetailsOperation, types.OperationType) (uuid.UUID, error)) *Repository_CreateOperation_Call {
	_c.Call.Return(run)
	return _c
}

// Decrease provides a mock function with given fields: _a0, _a1
func (_m *Repository) Decrease(_a0 context.Context, _a1 *models.Funds) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Decrease")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Funds) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Decrease_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decrease'
type Repository_Decrease_Call struct {
	*mock.Call
}

// Decrease is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.Funds
func (_e *Repository_Expecter) Decrease(_a0 interface{}, _a1 interface{}) *Repository_Decrease_Call {
	return &Repository_Decrease_Call{Call: _e.mock.On("Decrease", _a0, _a1)}
}

func (_c *Repository_Decrease_Call) Run(run func(_a0 context.Context, _a1 *models.Funds)) *Repository_Decrease_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Funds))
	})
	return _c
}

func (_c *Repository_Decrease_Call) Return(_a0 error) *Repository_Decrease_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Decrease_Call) RunAndReturn(run func(context.Context, *models.Funds) error) *Repository_Decrease_Call {
	_c.Call.Return(run)
	return _c
}

// GetAccountByUserID provides a mock function with given fields: _a0, _a1
func (_m *Repository) GetAccountByUserID(_a0 context.Context, _a1 uuid.UUID) (*models.Account, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetAccountByUserID")
	}

	var r0 *models.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*models.Account, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *models.Account); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_GetAccountByUserID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAccountByUserID'
type Repository_GetAccountByUserID_Call struct {
	*mock.Call
}

// GetAccountByUserID is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uuid.UUID
func (_e *Repository_Expecter) GetAccountByUserID(_a0 interface{}, _a1 interface{}) *Repository_GetAccountByUserID_Call {
	return &Repository_GetAccountByUserID_Call{Call: _e.mock.On("GetAccountByUserID", _a0, _a1)}
}

func (_c *Repository_GetAccountByUserID_Call) Run(run func(_a0 context.Context, _a1 uuid.UUID)) *Repository_GetAccountByUserID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Repository_GetAccountByUserID_Call) Return(_a0 *models.Account, _a1 error) *Repository_GetAccountByUserID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_GetAccountByUserID_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*models.Account, error)) *Repository_GetAccountByUserID_Call {
	_c.Call.Return(run)
	return _c
}

// LastOperations provides a mock function with given fields: _a0, _a1
func (_m *Repository) LastOperations(_a0 context.Context, _a1 uuid.UUID) (models.Operations, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LastOperations")
	}

	var r0 models.Operations
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (models.Operations, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) models.Operations); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.Operations)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repository_LastOperations_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LastOperations'
type Repository_LastOperations_Call struct {
	*mock.Call
}

// LastOperations is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 uuid.UUID
func (_e *Repository_Expecter) LastOperations(_a0 interface{}, _a1 interface{}) *Repository_LastOperations_Call {
	return &Repository_LastOperations_Call{Call: _e.mock.On("LastOperations", _a0, _a1)}
}

func (_c *Repository_LastOperations_Call) Run(run func(_a0 context.Context, _a1 uuid.UUID)) *Repository_LastOperations_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Repository_LastOperations_Call) Return(_a0 models.Operations, _a1 error) *Repository_LastOperations_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Repository_LastOperations_Call) RunAndReturn(run func(context.Context, uuid.UUID) (models.Operations, error)) *Repository_LastOperations_Call {
	_c.Call.Return(run)
	return _c
}

// Replenish provides a mock function with given fields: _a0, _a1
func (_m *Repository) Replenish(_a0 context.Context, _a1 *models.Funds) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Replenish")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Funds) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Repository_Replenish_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Replenish'
type Repository_Replenish_Call struct {
	*mock.Call
}

// Replenish is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *models.Funds
func (_e *Repository_Expecter) Replenish(_a0 interface{}, _a1 interface{}) *Repository_Replenish_Call {
	return &Repository_Replenish_Call{Call: _e.mock.On("Replenish", _a0, _a1)}
}

func (_c *Repository_Replenish_Call) Run(run func(_a0 context.Context, _a1 *models.Funds)) *Repository_Replenish_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Funds))
	})
	return _c
}

func (_c *Repository_Replenish_Call) Return(_a0 error) *Repository_Replenish_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Repository_Replenish_Call) RunAndReturn(run func(context.Context, *models.Funds) error) *Repository_Replenish_Call {
	_c.Call.Return(run)
	return _c
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
