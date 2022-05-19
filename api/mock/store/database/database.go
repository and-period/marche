// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	database "github.com/and-period/furumaru/api/internal/store/database"
	entity "github.com/and-period/furumaru/api/internal/store/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockStaff is a mock of Staff interface.
type MockStaff struct {
	ctrl     *gomock.Controller
	recorder *MockStaffMockRecorder
}

// MockStaffMockRecorder is the mock recorder for MockStaff.
type MockStaffMockRecorder struct {
	mock *MockStaff
}

// NewMockStaff creates a new mock instance.
func NewMockStaff(ctrl *gomock.Controller) *MockStaff {
	mock := &MockStaff{ctrl: ctrl}
	mock.recorder = &MockStaffMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStaff) EXPECT() *MockStaffMockRecorder {
	return m.recorder
}

// ListByStoreID mocks base method.
func (m *MockStaff) ListByStoreID(ctx context.Context, storeID int64, fields ...string) (entity.Staffs, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, storeID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListByStoreID", varargs...)
	ret0, _ := ret[0].(entity.Staffs)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByStoreID indicates an expected call of ListByStoreID.
func (mr *MockStaffMockRecorder) ListByStoreID(ctx, storeID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, storeID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByStoreID", reflect.TypeOf((*MockStaff)(nil).ListByStoreID), varargs...)
}

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStore) Create(ctx context.Context, store *entity.Store) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, store)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStoreMockRecorder) Create(ctx, store interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStore)(nil).Create), ctx, store)
}

// Get mocks base method.
func (m *MockStore) Get(ctx context.Context, storeID int64, fields ...string) (*entity.Store, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, storeID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStoreMockRecorder) Get(ctx, storeID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, storeID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStore)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockStore) List(ctx context.Context, params *database.ListStoresParams, fields ...string) (entity.Stores, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Stores)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockStoreMockRecorder) List(ctx, params interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStore)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockStore) Update(ctx context.Context, storeID int64, name, thumbnailURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, storeID, name, thumbnailURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockStoreMockRecorder) Update(ctx, storeID, name, thumbnailURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockStore)(nil).Update), ctx, storeID, name, thumbnailURL)
}
