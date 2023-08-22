// Code generated by mockery v2.33.0. DO NOT EDIT.

package mocks

import (
	context "context"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	mock "github.com/stretchr/testify/mock"
)

// DynamodbAPI is an autogenerated mock type for the DynamodbAPI type
type DynamodbAPI struct {
	mock.Mock
}

type DynamodbAPI_Expecter struct {
	mock *mock.Mock
}

func (_m *DynamodbAPI) EXPECT() *DynamodbAPI_Expecter {
	return &DynamodbAPI_Expecter{mock: &_m.Mock}
}

// CreateTable provides a mock function with given fields: ctx, params, optFns
func (_m *DynamodbAPI) CreateTable(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dynamodb.CreateTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.CreateTableInput, ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.CreateTableInput, ...func(*dynamodb.Options)) *dynamodb.CreateTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.CreateTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.CreateTableInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamodbAPI_CreateTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTable'
type DynamodbAPI_CreateTable_Call struct {
	*mock.Call
}

// CreateTable is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.CreateTableInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamodbAPI_Expecter) CreateTable(ctx interface{}, params interface{}, optFns ...interface{}) *DynamodbAPI_CreateTable_Call {
	return &DynamodbAPI_CreateTable_Call{Call: _e.mock.On("CreateTable",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamodbAPI_CreateTable_Call) Run(run func(ctx context.Context, params *dynamodb.CreateTableInput, optFns ...func(*dynamodb.Options))) *DynamodbAPI_CreateTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.CreateTableInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamodbAPI_CreateTable_Call) Return(_a0 *dynamodb.CreateTableOutput, _a1 error) *DynamodbAPI_CreateTable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamodbAPI_CreateTable_Call) RunAndReturn(run func(context.Context, *dynamodb.CreateTableInput, ...func(*dynamodb.Options)) (*dynamodb.CreateTableOutput, error)) *DynamodbAPI_CreateTable_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamodbAPI) DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dynamodb.DeleteItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) *dynamodb.DeleteItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.DeleteItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamodbAPI_DeleteItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteItem'
type DynamodbAPI_DeleteItem_Call struct {
	*mock.Call
}

// DeleteItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.DeleteItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamodbAPI_Expecter) DeleteItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamodbAPI_DeleteItem_Call {
	return &DynamodbAPI_DeleteItem_Call{Call: _e.mock.On("DeleteItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamodbAPI_DeleteItem_Call) Run(run func(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options))) *DynamodbAPI_DeleteItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.DeleteItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamodbAPI_DeleteItem_Call) Return(_a0 *dynamodb.DeleteItemOutput, _a1 error) *DynamodbAPI_DeleteItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamodbAPI_DeleteItem_Call) RunAndReturn(run func(context.Context, *dynamodb.DeleteItemInput, ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)) *DynamodbAPI_DeleteItem_Call {
	_c.Call.Return(run)
	return _c
}

// DescribeTable provides a mock function with given fields: ctx, params, optFns
func (_m *DynamodbAPI) DescribeTable(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dynamodb.DescribeTableOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) *dynamodb.DescribeTableOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.DescribeTableOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamodbAPI_DescribeTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DescribeTable'
type DynamodbAPI_DescribeTable_Call struct {
	*mock.Call
}

// DescribeTable is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.DescribeTableInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamodbAPI_Expecter) DescribeTable(ctx interface{}, params interface{}, optFns ...interface{}) *DynamodbAPI_DescribeTable_Call {
	return &DynamodbAPI_DescribeTable_Call{Call: _e.mock.On("DescribeTable",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamodbAPI_DescribeTable_Call) Run(run func(ctx context.Context, params *dynamodb.DescribeTableInput, optFns ...func(*dynamodb.Options))) *DynamodbAPI_DescribeTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.DescribeTableInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamodbAPI_DescribeTable_Call) Return(_a0 *dynamodb.DescribeTableOutput, _a1 error) *DynamodbAPI_DescribeTable_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamodbAPI_DescribeTable_Call) RunAndReturn(run func(context.Context, *dynamodb.DescribeTableInput, ...func(*dynamodb.Options)) (*dynamodb.DescribeTableOutput, error)) *DynamodbAPI_DescribeTable_Call {
	_c.Call.Return(run)
	return _c
}

// GetItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamodbAPI) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dynamodb.GetItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) *dynamodb.GetItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.GetItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamodbAPI_GetItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetItem'
type DynamodbAPI_GetItem_Call struct {
	*mock.Call
}

// GetItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.GetItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamodbAPI_Expecter) GetItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamodbAPI_GetItem_Call {
	return &DynamodbAPI_GetItem_Call{Call: _e.mock.On("GetItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamodbAPI_GetItem_Call) Run(run func(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options))) *DynamodbAPI_GetItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.GetItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamodbAPI_GetItem_Call) Return(_a0 *dynamodb.GetItemOutput, _a1 error) *DynamodbAPI_GetItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamodbAPI_GetItem_Call) RunAndReturn(run func(context.Context, *dynamodb.GetItemInput, ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)) *DynamodbAPI_GetItem_Call {
	_c.Call.Return(run)
	return _c
}

// PutItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamodbAPI) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dynamodb.PutItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) *dynamodb.PutItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.PutItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamodbAPI_PutItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PutItem'
type DynamodbAPI_PutItem_Call struct {
	*mock.Call
}

// PutItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.PutItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamodbAPI_Expecter) PutItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamodbAPI_PutItem_Call {
	return &DynamodbAPI_PutItem_Call{Call: _e.mock.On("PutItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamodbAPI_PutItem_Call) Run(run func(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options))) *DynamodbAPI_PutItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.PutItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamodbAPI_PutItem_Call) Return(_a0 *dynamodb.PutItemOutput, _a1 error) *DynamodbAPI_PutItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamodbAPI_PutItem_Call) RunAndReturn(run func(context.Context, *dynamodb.PutItemInput, ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)) *DynamodbAPI_PutItem_Call {
	_c.Call.Return(run)
	return _c
}

// Query provides a mock function with given fields: ctx, params, optFns
func (_m *DynamodbAPI) Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dynamodb.QueryOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) *dynamodb.QueryOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.QueryOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamodbAPI_Query_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Query'
type DynamodbAPI_Query_Call struct {
	*mock.Call
}

// Query is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.QueryInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamodbAPI_Expecter) Query(ctx interface{}, params interface{}, optFns ...interface{}) *DynamodbAPI_Query_Call {
	return &DynamodbAPI_Query_Call{Call: _e.mock.On("Query",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamodbAPI_Query_Call) Run(run func(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options))) *DynamodbAPI_Query_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.QueryInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamodbAPI_Query_Call) Return(_a0 *dynamodb.QueryOutput, _a1 error) *DynamodbAPI_Query_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamodbAPI_Query_Call) RunAndReturn(run func(context.Context, *dynamodb.QueryInput, ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)) *DynamodbAPI_Query_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateItem provides a mock function with given fields: ctx, params, optFns
func (_m *DynamodbAPI) UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error) {
	_va := make([]interface{}, len(optFns))
	for _i := range optFns {
		_va[_i] = optFns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, params)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *dynamodb.UpdateItemOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)); ok {
		return rf(ctx, params, optFns...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) *dynamodb.UpdateItemOutput); ok {
		r0 = rf(ctx, params, optFns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dynamodb.UpdateItemOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) error); ok {
		r1 = rf(ctx, params, optFns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DynamodbAPI_UpdateItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateItem'
type DynamodbAPI_UpdateItem_Call struct {
	*mock.Call
}

// UpdateItem is a helper method to define mock.On call
//   - ctx context.Context
//   - params *dynamodb.UpdateItemInput
//   - optFns ...func(*dynamodb.Options)
func (_e *DynamodbAPI_Expecter) UpdateItem(ctx interface{}, params interface{}, optFns ...interface{}) *DynamodbAPI_UpdateItem_Call {
	return &DynamodbAPI_UpdateItem_Call{Call: _e.mock.On("UpdateItem",
		append([]interface{}{ctx, params}, optFns...)...)}
}

func (_c *DynamodbAPI_UpdateItem_Call) Run(run func(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options))) *DynamodbAPI_UpdateItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*dynamodb.Options), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*dynamodb.Options))
			}
		}
		run(args[0].(context.Context), args[1].(*dynamodb.UpdateItemInput), variadicArgs...)
	})
	return _c
}

func (_c *DynamodbAPI_UpdateItem_Call) Return(_a0 *dynamodb.UpdateItemOutput, _a1 error) *DynamodbAPI_UpdateItem_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *DynamodbAPI_UpdateItem_Call) RunAndReturn(run func(context.Context, *dynamodb.UpdateItemInput, ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)) *DynamodbAPI_UpdateItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewDynamodbAPI creates a new instance of DynamodbAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDynamodbAPI(t interface {
	mock.TestingT
	Cleanup(func())
}) *DynamodbAPI {
	mock := &DynamodbAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}