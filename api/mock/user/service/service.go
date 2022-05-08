// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	entity "github.com/and-period/marche/api/internal/user/entity"
	service "github.com/and-period/marche/api/internal/user/service"
	gomock "github.com/golang/mock/gomock"
)

// MockUserService is a mock of UserService interface.
type MockUserService struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceMockRecorder
}

// MockUserServiceMockRecorder is the mock recorder for MockUserService.
type MockUserServiceMockRecorder struct {
	mock *MockUserService
}

// NewMockUserService creates a new mock instance.
func NewMockUserService(ctrl *gomock.Controller) *MockUserService {
	mock := &MockUserService{ctrl: ctrl}
	mock.recorder = &MockUserServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserService) EXPECT() *MockUserServiceMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserService) CreateUser(ctx context.Context, in *service.CreateUserInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, in)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserServiceMockRecorder) CreateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserService)(nil).CreateUser), ctx, in)
}

// CreateUserWithOAuth mocks base method.
func (m *MockUserService) CreateUserWithOAuth(ctx context.Context, in *service.CreateUserWithOAuthInput) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserWithOAuth", ctx, in)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserWithOAuth indicates an expected call of CreateUserWithOAuth.
func (mr *MockUserServiceMockRecorder) CreateUserWithOAuth(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserWithOAuth", reflect.TypeOf((*MockUserService)(nil).CreateUserWithOAuth), ctx, in)
}

// DeleteUser mocks base method.
func (m *MockUserService) DeleteUser(ctx context.Context, in *service.DeleteUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockUserServiceMockRecorder) DeleteUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserService)(nil).DeleteUser), ctx, in)
}

// ForgotUserPassword mocks base method.
func (m *MockUserService) ForgotUserPassword(ctx context.Context, in *service.ForgotUserPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotUserPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotUserPassword indicates an expected call of ForgotUserPassword.
func (mr *MockUserServiceMockRecorder) ForgotUserPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotUserPassword", reflect.TypeOf((*MockUserService)(nil).ForgotUserPassword), ctx, in)
}

// GetAdmin mocks base method.
func (m *MockUserService) GetAdmin(ctx context.Context, in *service.GetAdminInput) (*entity.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", ctx, in)
	ret0, _ := ret[0].(*entity.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockUserServiceMockRecorder) GetAdmin(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockUserService)(nil).GetAdmin), ctx, in)
}

// GetAdminAuth mocks base method.
func (m *MockUserService) GetAdminAuth(ctx context.Context, in *service.GetAdminAuthInput) (*entity.AdminAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminAuth", ctx, in)
	ret0, _ := ret[0].(*entity.AdminAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminAuth indicates an expected call of GetAdminAuth.
func (mr *MockUserServiceMockRecorder) GetAdminAuth(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminAuth", reflect.TypeOf((*MockUserService)(nil).GetAdminAuth), ctx, in)
}

// GetShop mocks base method.
func (m *MockUserService) GetShop(ctx context.Context, in *service.GetShopInput) (*entity.Shop, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShop", ctx, in)
	ret0, _ := ret[0].(*entity.Shop)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShop indicates an expected call of GetShop.
func (mr *MockUserServiceMockRecorder) GetShop(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShop", reflect.TypeOf((*MockUserService)(nil).GetShop), ctx, in)
}

// GetUser mocks base method.
func (m *MockUserService) GetUser(ctx context.Context, in *service.GetUserInput) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, in)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserServiceMockRecorder) GetUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserService)(nil).GetUser), ctx, in)
}

// GetUserAuth mocks base method.
func (m *MockUserService) GetUserAuth(ctx context.Context, in *service.GetUserAuthInput) (*entity.UserAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAuth", ctx, in)
	ret0, _ := ret[0].(*entity.UserAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAuth indicates an expected call of GetUserAuth.
func (mr *MockUserServiceMockRecorder) GetUserAuth(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAuth", reflect.TypeOf((*MockUserService)(nil).GetUserAuth), ctx, in)
}

// InitializeUser mocks base method.
func (m *MockUserService) InitializeUser(ctx context.Context, in *service.InitializeUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeUser", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeUser indicates an expected call of InitializeUser.
func (mr *MockUserServiceMockRecorder) InitializeUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeUser", reflect.TypeOf((*MockUserService)(nil).InitializeUser), ctx, in)
}

// MultiGetShops mocks base method.
func (m *MockUserService) MultiGetShops(ctx context.Context, in *service.MultiGetShopsInput) (entity.Shops, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetShops", ctx, in)
	ret0, _ := ret[0].(entity.Shops)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetShops indicates an expected call of MultiGetShops.
func (mr *MockUserServiceMockRecorder) MultiGetShops(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetShops", reflect.TypeOf((*MockUserService)(nil).MultiGetShops), ctx, in)
}

// RefreshAdminToken mocks base method.
func (m *MockUserService) RefreshAdminToken(ctx context.Context, in *service.RefreshAdminTokenInput) (*entity.AdminAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshAdminToken", ctx, in)
	ret0, _ := ret[0].(*entity.AdminAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshAdminToken indicates an expected call of RefreshAdminToken.
func (mr *MockUserServiceMockRecorder) RefreshAdminToken(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshAdminToken", reflect.TypeOf((*MockUserService)(nil).RefreshAdminToken), ctx, in)
}

// RefreshUserToken mocks base method.
func (m *MockUserService) RefreshUserToken(ctx context.Context, in *service.RefreshUserTokenInput) (*entity.UserAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshUserToken", ctx, in)
	ret0, _ := ret[0].(*entity.UserAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshUserToken indicates an expected call of RefreshUserToken.
func (mr *MockUserServiceMockRecorder) RefreshUserToken(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshUserToken", reflect.TypeOf((*MockUserService)(nil).RefreshUserToken), ctx, in)
}

// SignInAdmin mocks base method.
func (m *MockUserService) SignInAdmin(ctx context.Context, in *service.SignInAdminInput) (*entity.AdminAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInAdmin", ctx, in)
	ret0, _ := ret[0].(*entity.AdminAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInAdmin indicates an expected call of SignInAdmin.
func (mr *MockUserServiceMockRecorder) SignInAdmin(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInAdmin", reflect.TypeOf((*MockUserService)(nil).SignInAdmin), ctx, in)
}

// SignInUser mocks base method.
func (m *MockUserService) SignInUser(ctx context.Context, in *service.SignInUserInput) (*entity.UserAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInUser", ctx, in)
	ret0, _ := ret[0].(*entity.UserAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInUser indicates an expected call of SignInUser.
func (mr *MockUserServiceMockRecorder) SignInUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInUser", reflect.TypeOf((*MockUserService)(nil).SignInUser), ctx, in)
}

// SignOutAdmin mocks base method.
func (m *MockUserService) SignOutAdmin(ctx context.Context, in *service.SignOutAdminInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignOutAdmin", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignOutAdmin indicates an expected call of SignOutAdmin.
func (mr *MockUserServiceMockRecorder) SignOutAdmin(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignOutAdmin", reflect.TypeOf((*MockUserService)(nil).SignOutAdmin), ctx, in)
}

// SignOutUser mocks base method.
func (m *MockUserService) SignOutUser(ctx context.Context, in *service.SignOutUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignOutUser", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignOutUser indicates an expected call of SignOutUser.
func (mr *MockUserServiceMockRecorder) SignOutUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignOutUser", reflect.TypeOf((*MockUserService)(nil).SignOutUser), ctx, in)
}

// UpdateAdminEmail mocks base method.
func (m *MockUserService) UpdateAdminEmail(ctx context.Context, in *service.UpdateAdminEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminEmail indicates an expected call of UpdateAdminEmail.
func (mr *MockUserServiceMockRecorder) UpdateAdminEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminEmail", reflect.TypeOf((*MockUserService)(nil).UpdateAdminEmail), ctx, in)
}

// UpdateAdminPassword mocks base method.
func (m *MockUserService) UpdateAdminPassword(ctx context.Context, in *service.UpdateAdminPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminPassword indicates an expected call of UpdateAdminPassword.
func (mr *MockUserServiceMockRecorder) UpdateAdminPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminPassword", reflect.TypeOf((*MockUserService)(nil).UpdateAdminPassword), ctx, in)
}

// UpdateUserEmail mocks base method.
func (m *MockUserService) UpdateUserEmail(ctx context.Context, in *service.UpdateUserEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserEmail indicates an expected call of UpdateUserEmail.
func (mr *MockUserServiceMockRecorder) UpdateUserEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserEmail", reflect.TypeOf((*MockUserService)(nil).UpdateUserEmail), ctx, in)
}

// UpdateUserPassword mocks base method.
func (m *MockUserService) UpdateUserPassword(ctx context.Context, in *service.UpdateUserPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockUserServiceMockRecorder) UpdateUserPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockUserService)(nil).UpdateUserPassword), ctx, in)
}

// VerifyAdminEmail mocks base method.
func (m *MockUserService) VerifyAdminEmail(ctx context.Context, in *service.VerifyAdminEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAdminEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyAdminEmail indicates an expected call of VerifyAdminEmail.
func (mr *MockUserServiceMockRecorder) VerifyAdminEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAdminEmail", reflect.TypeOf((*MockUserService)(nil).VerifyAdminEmail), ctx, in)
}

// VerifyUser mocks base method.
func (m *MockUserService) VerifyUser(ctx context.Context, in *service.VerifyUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUser indicates an expected call of VerifyUser.
func (mr *MockUserServiceMockRecorder) VerifyUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockUserService)(nil).VerifyUser), ctx, in)
}

// VerifyUserEmail mocks base method.
func (m *MockUserService) VerifyUserEmail(ctx context.Context, in *service.VerifyUserEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUserEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUserEmail indicates an expected call of VerifyUserEmail.
func (mr *MockUserServiceMockRecorder) VerifyUserEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUserEmail", reflect.TypeOf((*MockUserService)(nil).VerifyUserEmail), ctx, in)
}

// VerifyUserPassword mocks base method.
func (m *MockUserService) VerifyUserPassword(ctx context.Context, in *service.VerifyUserPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUserPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUserPassword indicates an expected call of VerifyUserPassword.
func (mr *MockUserServiceMockRecorder) VerifyUserPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUserPassword", reflect.TypeOf((*MockUserService)(nil).VerifyUserPassword), ctx, in)
}
