// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	database "github.com/and-period/furumaru/api/internal/user/database"
	entity "github.com/and-period/furumaru/api/internal/user/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAdmin) Create(ctx context.Context, admin *entity.Admin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, admin)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAdminMockRecorder) Create(ctx, admin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdmin)(nil).Create), ctx, admin)
}

// Get mocks base method.
func (m *MockAdmin) Get(ctx context.Context, adminID string, fields ...string) (*entity.Admin, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, adminID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockAdminMockRecorder) Get(ctx, adminID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, adminID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAdmin)(nil).Get), varargs...)
}

// GetByCognitoID mocks base method.
func (m *MockAdmin) GetByCognitoID(ctx context.Context, cognitoID string, fields ...string) (*entity.Admin, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, cognitoID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByCognitoID", varargs...)
	ret0, _ := ret[0].(*entity.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCognitoID indicates an expected call of GetByCognitoID.
func (mr *MockAdminMockRecorder) GetByCognitoID(ctx, cognitoID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, cognitoID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCognitoID", reflect.TypeOf((*MockAdmin)(nil).GetByCognitoID), varargs...)
}

// List mocks base method.
func (m *MockAdmin) List(ctx context.Context, params *database.ListAdminsParams, fields ...string) (entity.Admins, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Admins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockAdminMockRecorder) List(ctx, params interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockAdmin)(nil).List), varargs...)
}

// MultiGet mocks base method.
func (m *MockAdmin) MultiGet(ctx context.Context, adminIDs []string, fields ...string) (entity.Admins, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, adminIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.Admins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockAdminMockRecorder) MultiGet(ctx, adminIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, adminIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockAdmin)(nil).MultiGet), varargs...)
}

// UpdateEmail mocks base method.
func (m *MockAdmin) UpdateEmail(ctx context.Context, adminID, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmail", ctx, adminID, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmail indicates an expected call of UpdateEmail.
func (mr *MockAdminMockRecorder) UpdateEmail(ctx, adminID, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockAdmin)(nil).UpdateEmail), ctx, adminID, email)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUser) Create(ctx context.Context, user *entity.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUserMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), ctx, user)
}

// Delete mocks base method.
func (m *MockUser) Delete(ctx context.Context, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserMockRecorder) Delete(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUser)(nil).Delete), ctx, userID)
}

// Get mocks base method.
func (m *MockUser) Get(ctx context.Context, userID string, fields ...string) (*entity.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, userID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserMockRecorder) Get(ctx, userID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, userID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUser)(nil).Get), varargs...)
}

// GetByCognitoID mocks base method.
func (m *MockUser) GetByCognitoID(ctx context.Context, cognitoID string, fields ...string) (*entity.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, cognitoID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByCognitoID", varargs...)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCognitoID indicates an expected call of GetByCognitoID.
func (mr *MockUserMockRecorder) GetByCognitoID(ctx, cognitoID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, cognitoID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCognitoID", reflect.TypeOf((*MockUser)(nil).GetByCognitoID), varargs...)
}

// GetByEmail mocks base method.
func (m *MockUser) GetByEmail(ctx context.Context, email string, fields ...string) (*entity.User, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, email}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetByEmail", varargs...)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockUserMockRecorder) GetByEmail(ctx, email interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, email}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockUser)(nil).GetByEmail), varargs...)
}

// InitializeUser mocks base method.
func (m *MockUser) InitializeUser(ctx context.Context, userID, accountID, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeUser", ctx, userID, accountID, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeUser indicates an expected call of InitializeUser.
func (mr *MockUserMockRecorder) InitializeUser(ctx, userID, accountID, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeUser", reflect.TypeOf((*MockUser)(nil).InitializeUser), ctx, userID, accountID, username)
}

// UpdateEmail mocks base method.
func (m *MockUser) UpdateEmail(ctx context.Context, userID, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmail", ctx, userID, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmail indicates an expected call of UpdateEmail.
func (mr *MockUserMockRecorder) UpdateEmail(ctx, userID, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmail", reflect.TypeOf((*MockUser)(nil).UpdateEmail), ctx, userID, email)
}

// UpdateVerified mocks base method.
func (m *MockUser) UpdateVerified(ctx context.Context, userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVerified", ctx, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVerified indicates an expected call of UpdateVerified.
func (mr *MockUserMockRecorder) UpdateVerified(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerified", reflect.TypeOf((*MockUser)(nil).UpdateVerified), ctx, userID)
}
