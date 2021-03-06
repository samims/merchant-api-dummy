// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"

	orm "github.com/astaxie/beego/orm"
	mock "github.com/stretchr/testify/mock"

	sql "database/sql"

	testing "testing"
)

// OrmerMock is an autogenerated mock type for the Ormer type
type OrmerMock struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *OrmerMock) Begin() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BeginTx provides a mock function with given fields: ctx, opts
func (_m *OrmerMock) BeginTx(ctx context.Context, opts *sql.TxOptions) error {
	ret := _m.Called(ctx, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *sql.TxOptions) error); ok {
		r0 = rf(ctx, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Commit provides a mock function with given fields:
func (_m *OrmerMock) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DBStats provides a mock function with given fields:
func (_m *OrmerMock) DBStats() *sql.DBStats {
	ret := _m.Called()

	var r0 *sql.DBStats
	if rf, ok := ret.Get(0).(func() *sql.DBStats); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.DBStats)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: md, cols
func (_m *OrmerMock) Delete(md interface{}, cols ...string) (int64, error) {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, md)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(interface{}, ...string) int64); ok {
		r0 = rf(md, cols...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...string) error); ok {
		r1 = rf(md, cols...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Driver provides a mock function with given fields:
func (_m *OrmerMock) Driver() orm.Driver {
	ret := _m.Called()

	var r0 orm.Driver
	if rf, ok := ret.Get(0).(func() orm.Driver); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Driver)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: _a0
func (_m *OrmerMock) Insert(_a0 interface{}) (int64, error) {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(interface{}) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertMulti provides a mock function with given fields: bulk, mds
func (_m *OrmerMock) InsertMulti(bulk int, mds interface{}) (int64, error) {
	ret := _m.Called(bulk, mds)

	var r0 int64
	if rf, ok := ret.Get(0).(func(int, interface{}) int64); ok {
		r0 = rf(bulk, mds)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, interface{}) error); ok {
		r1 = rf(bulk, mds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOrUpdate provides a mock function with given fields: md, colConflitAndArgs
func (_m *OrmerMock) InsertOrUpdate(md interface{}, colConflitAndArgs ...string) (int64, error) {
	_va := make([]interface{}, len(colConflitAndArgs))
	for _i := range colConflitAndArgs {
		_va[_i] = colConflitAndArgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, md)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(interface{}, ...string) int64); ok {
		r0 = rf(md, colConflitAndArgs...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...string) error); ok {
		r1 = rf(md, colConflitAndArgs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoadRelated provides a mock function with given fields: md, name, args
func (_m *OrmerMock) LoadRelated(md interface{}, name string, args ...interface{}) (int64, error) {
	var _ca []interface{}
	_ca = append(_ca, md, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(interface{}, string, ...interface{}) int64); ok {
		r0 = rf(md, name, args...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string, ...interface{}) error); ok {
		r1 = rf(md, name, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryM2M provides a mock function with given fields: md, name
func (_m *OrmerMock) QueryM2M(md interface{}, name string) orm.QueryM2Mer {
	ret := _m.Called(md, name)

	var r0 orm.QueryM2Mer
	if rf, ok := ret.Get(0).(func(interface{}, string) orm.QueryM2Mer); ok {
		r0 = rf(md, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QueryM2Mer)
		}
	}

	return r0
}

// QueryTable provides a mock function with given fields: ptrStructOrTableName
func (_m *OrmerMock) QueryTable(ptrStructOrTableName interface{}) orm.QuerySeter {
	ret := _m.Called(ptrStructOrTableName)

	var r0 orm.QuerySeter
	if rf, ok := ret.Get(0).(func(interface{}) orm.QuerySeter); ok {
		r0 = rf(ptrStructOrTableName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.QuerySeter)
		}
	}

	return r0
}

// Raw provides a mock function with given fields: query, args
func (_m *OrmerMock) Raw(query string, args ...interface{}) orm.RawSeter {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 orm.RawSeter
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.RawSeter); ok {
		r0 = rf(query, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.RawSeter)
		}
	}

	return r0
}

// Read provides a mock function with given fields: md, cols
func (_m *OrmerMock) Read(md interface{}, cols ...string) error {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, md)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...string) error); ok {
		r0 = rf(md, cols...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadForUpdate provides a mock function with given fields: md, cols
func (_m *OrmerMock) ReadForUpdate(md interface{}, cols ...string) error {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, md)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, ...string) error); ok {
		r0 = rf(md, cols...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReadOrCreate provides a mock function with given fields: md, col1, cols
func (_m *OrmerMock) ReadOrCreate(md interface{}, col1 string, cols ...string) (bool, int64, error) {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, md, col1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(interface{}, string, ...string) bool); ok {
		r0 = rf(md, col1, cols...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(interface{}, string, ...string) int64); ok {
		r1 = rf(md, col1, cols...)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(interface{}, string, ...string) error); ok {
		r2 = rf(md, col1, cols...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Rollback provides a mock function with given fields:
func (_m *OrmerMock) Rollback() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: md, cols
func (_m *OrmerMock) Update(md interface{}, cols ...string) (int64, error) {
	_va := make([]interface{}, len(cols))
	for _i := range cols {
		_va[_i] = cols[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, md)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 int64
	if rf, ok := ret.Get(0).(func(interface{}, ...string) int64); ok {
		r0 = rf(md, cols...)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...string) error); ok {
		r1 = rf(md, cols...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Using provides a mock function with given fields: name
func (_m *OrmerMock) Using(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOrmerMock creates a new instance of OrmerMock. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewOrmerMock(t testing.TB) *OrmerMock {
	mock := &OrmerMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
