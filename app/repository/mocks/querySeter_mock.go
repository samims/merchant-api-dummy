// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	orm "github.com/astaxie/beego/orm"
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// QuerySeterMock is an autogenerated mock type for the QuerySeter type
type QuerySeterMock struct {
	mock.Mock
}

// All provides a mock function with given fields: container, cols
func (_m *QuerySeterMock) All(container interface{}, cols ...string) (int64, error) {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, container)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(interface{}, ...string) int64); ok {
		r0 = rf(container, cols...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...string) error); ok {
		r1 = rf(container, cols...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Count provides a mock function with given fields:
func (_m *QuerySeterMock) Count() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields:
func (_m *QuerySeterMock) Delete() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Distinct provides a mock function with given fields:
func (_m *QuerySeterMock) Distinct() orm.QuerySeter {
	ret := _m.Called()

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func() orm.QuerySeter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// Exclude provides a mock function with given fields: _a0, _a1
func (_m *QuerySeterMock) Exclude(_a0 string, _a1 ...interface{}) orm.QuerySeter {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.QuerySeter); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// Exist provides a mock function with given fields:
func (_m *QuerySeterMock) Exist() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Filter provides a mock function with given fields: _a0, _a1
func (_m *QuerySeterMock) Filter(_a0 string, _a1 ...interface{}) orm.QuerySeter {
	var _ca []interface{}
	_ca = append(_ca, _a0)
	_ca = append(_ca, _a1...)
	ret := _m.Called(_ca...)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.QuerySeter); ok {
		r0 = rf(_a0, _a1...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// FilterRaw provides a mock function with given fields: _a0, _a1
func (_m *QuerySeterMock) FilterRaw(_a0 string, _a1 string) orm.QuerySeter {
	ret := _m.Called(_a0, _a1)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(string, string) orm.QuerySeter); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// ForUpdate provides a mock function with given fields:
func (_m *QuerySeterMock) ForUpdate() orm.QuerySeter {
	ret := _m.Called()

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func() orm.QuerySeter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// GetCond provides a mock function with given fields:
func (_m *QuerySeterMock) GetCond() *orm.Condition {
	ret := _m.Called()

	var r0 *orm.Condition
	if rf, ok := ret.Get(0).(func() *orm.Condition); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*orm.Condition)
		}
	}

	return r0
}

// GroupBy provides a mock function with given fields: exprs
func (_m *QuerySeterMock) GroupBy(exprs ...string) orm.QuerySeter {
	_va := make([]interface{}, len(exprs))
	for _i := range exprs {
		_va[_i] = exprs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(...string) orm.QuerySeter); ok {
		r0 = rf(exprs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// Limit provides a mock function with given fields: limit, args
func (_m *QuerySeterMock) Limit(limit interface{}, args ...interface{}) orm.QuerySeter {
	var _ca []interface{}
	_ca = append(_ca, limit)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) orm.QuerySeter); ok {
		r0 = rf(limit, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// Offset provides a mock function with given fields: offset
func (_m *QuerySeterMock) Offset(offset interface{}) orm.QuerySeter {
	ret := _m.Called(offset)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(interface{}) orm.QuerySeter); ok {
		r0 = rf(offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// One provides a mock function with given fields: container, cols
func (_m *QuerySeterMock) One(container interface{}, cols ...string) error {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, container)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...string) error); ok {
		r0 = rf(container, cols...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// OrderBy provides a mock function with given fields: exprs
func (_m *QuerySeterMock) OrderBy(exprs ...string) orm.QuerySeter {
	_va := make([]interface{}, len(exprs))
	for _i := range exprs {
		_va[_i] = exprs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(...string) orm.QuerySeter); ok {
		r0 = rf(exprs...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// PrepareInsert provides a mock function with given fields:
func (_m *QuerySeterMock) PrepareInsert() (orm.Inserter, error) {
	ret := _m.Called()

	var r0 orm.Inserter
	if rf, ok := ret.Get(0).(func() orm.Inserter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Inserter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RelatedSel provides a mock function with given fields: params
func (_m *QuerySeterMock) RelatedSel(params ...interface{}) orm.QuerySeter {
	var _ca []interface{}
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(...interface{}) orm.QuerySeter); ok {
		r0 = rf(params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// RowsToMap provides a mock function with given fields: result, keyCol, valueCol
func (_m *QuerySeterMock) RowsToMap(result *orm.Params, keyCol string, valueCol string) (int64, error) {
	ret := _m.Called(result, keyCol, valueCol)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*orm.Params, string, string) int64); ok {
		r0 = rf(result, keyCol, valueCol)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*orm.Params, string, string) error); ok {
		r1 = rf(result, keyCol, valueCol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RowsToStruct provides a mock function with given fields: ptrStruct, keyCol, valueCol
func (_m *QuerySeterMock) RowsToStruct(ptrStruct interface{}, keyCol string, valueCol string) (int64, error) {
	ret := _m.Called(ptrStruct, keyCol, valueCol)

	var r0 int64
	if rf, ok := ret.Get(0).(func(interface{}, string, string) int64); ok {
		r0 = rf(ptrStruct, keyCol, valueCol)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, string) error); ok {
		r1 = rf(ptrStruct, keyCol, valueCol)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetCond provides a mock function with given fields: _a0
func (_m *QuerySeterMock) SetCond(_a0 *orm.Condition) orm.QuerySeter {
	ret := _m.Called(_a0)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(*orm.Condition) orm.QuerySeter); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// Update provides a mock function with given fields: values
func (_m *QuerySeterMock) Update(values orm.Params) (int64, error) {
	ret := _m.Called(values)

	var r0 int64
	if rf, ok := ret.Get(0).(func(orm.Params) int64); ok {
		r0 = rf(values)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(orm.Params) error); ok {
		r1 = rf(values)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Values provides a mock function with given fields: results, exprs
func (_m *QuerySeterMock) Values(results *[]orm.Params, exprs ...string) (int64, error) {
	_va := make([]interface{}, len(exprs))
	for _i := range exprs {
		_va[_i] = exprs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, results)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*[]orm.Params, ...string) int64); ok {
		r0 = rf(results, exprs...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*[]orm.Params, ...string) error); ok {
		r1 = rf(results, exprs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValuesFlat provides a mock function with given fields: result, expr
func (_m *QuerySeterMock) ValuesFlat(result *orm.ParamsList, expr string) (int64, error) {
	ret := _m.Called(result, expr)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*orm.ParamsList, string) int64); ok {
		r0 = rf(result, expr)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*orm.ParamsList, string) error); ok {
		r1 = rf(result, expr)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValuesList provides a mock function with given fields: results, exprs
func (_m *QuerySeterMock) ValuesList(results *[]orm.ParamsList, exprs ...string) (int64, error) {
	_va := make([]interface{}, len(exprs))
	for _i := range exprs {
		_va[_i] = exprs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, results)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*[]orm.ParamsList, ...string) int64); ok {
		r0 = rf(results, exprs...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*[]orm.ParamsList, ...string) error); ok {
		r1 = rf(results, exprs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewQuerySeterMock creates a new instance of QuerySeterMock. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewQuerySeterMock(t testing.TB) *QuerySeterMock {
	mock := &QuerySeterMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
