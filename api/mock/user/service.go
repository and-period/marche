// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	context "context"
	reflect "reflect"

	user "github.com/and-period/furumaru/api/internal/user"
	entity "github.com/and-period/furumaru/api/internal/user/entity"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AggregateRealatedProducers mocks base method.
func (m *MockService) AggregateRealatedProducers(ctx context.Context, in *user.AggregateRealatedProducersInput) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AggregateRealatedProducers", ctx, in)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AggregateRealatedProducers indicates an expected call of AggregateRealatedProducers.
func (mr *MockServiceMockRecorder) AggregateRealatedProducers(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AggregateRealatedProducers", reflect.TypeOf((*MockService)(nil).AggregateRealatedProducers), ctx, in)
}

// CreateAdministrator mocks base method.
func (m *MockService) CreateAdministrator(ctx context.Context, in *user.CreateAdministratorInput) (*entity.Administrator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdministrator", ctx, in)
	ret0, _ := ret[0].(*entity.Administrator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAdministrator indicates an expected call of CreateAdministrator.
func (mr *MockServiceMockRecorder) CreateAdministrator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdministrator", reflect.TypeOf((*MockService)(nil).CreateAdministrator), ctx, in)
}

// CreateCoordinator mocks base method.
func (m *MockService) CreateCoordinator(ctx context.Context, in *user.CreateCoordinatorInput) (*entity.Coordinator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCoordinator", ctx, in)
	ret0, _ := ret[0].(*entity.Coordinator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCoordinator indicates an expected call of CreateCoordinator.
func (mr *MockServiceMockRecorder) CreateCoordinator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCoordinator", reflect.TypeOf((*MockService)(nil).CreateCoordinator), ctx, in)
}

// CreateProducer mocks base method.
func (m *MockService) CreateProducer(ctx context.Context, in *user.CreateProducerInput) (*entity.Producer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProducer", ctx, in)
	ret0, _ := ret[0].(*entity.Producer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProducer indicates an expected call of CreateProducer.
func (mr *MockServiceMockRecorder) CreateProducer(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProducer", reflect.TypeOf((*MockService)(nil).CreateProducer), ctx, in)
}

// CreateUser mocks base method.
func (m *MockService) CreateUser(ctx context.Context, in *user.CreateUserInput) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, in)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockServiceMockRecorder) CreateUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockService)(nil).CreateUser), ctx, in)
}

// CreateUserWithOAuth mocks base method.
func (m *MockService) CreateUserWithOAuth(ctx context.Context, in *user.CreateUserWithOAuthInput) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUserWithOAuth", ctx, in)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUserWithOAuth indicates an expected call of CreateUserWithOAuth.
func (mr *MockServiceMockRecorder) CreateUserWithOAuth(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUserWithOAuth", reflect.TypeOf((*MockService)(nil).CreateUserWithOAuth), ctx, in)
}

// DeleteAdministrator mocks base method.
func (m *MockService) DeleteAdministrator(ctx context.Context, in *user.DeleteAdministratorInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAdministrator", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAdministrator indicates an expected call of DeleteAdministrator.
func (mr *MockServiceMockRecorder) DeleteAdministrator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdministrator", reflect.TypeOf((*MockService)(nil).DeleteAdministrator), ctx, in)
}

// DeleteCoordinator mocks base method.
func (m *MockService) DeleteCoordinator(ctx context.Context, in *user.DeleteCoordinatorInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCoordinator", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCoordinator indicates an expected call of DeleteCoordinator.
func (mr *MockServiceMockRecorder) DeleteCoordinator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCoordinator", reflect.TypeOf((*MockService)(nil).DeleteCoordinator), ctx, in)
}

// DeleteProducer mocks base method.
func (m *MockService) DeleteProducer(ctx context.Context, in *user.DeleteProducerInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProducer", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProducer indicates an expected call of DeleteProducer.
func (mr *MockServiceMockRecorder) DeleteProducer(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProducer", reflect.TypeOf((*MockService)(nil).DeleteProducer), ctx, in)
}

// DeleteUser mocks base method.
func (m *MockService) DeleteUser(ctx context.Context, in *user.DeleteUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockServiceMockRecorder) DeleteUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockService)(nil).DeleteUser), ctx, in)
}

// ForgotAdminPassword mocks base method.
func (m *MockService) ForgotAdminPassword(ctx context.Context, in *user.ForgotAdminPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotAdminPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotAdminPassword indicates an expected call of ForgotAdminPassword.
func (mr *MockServiceMockRecorder) ForgotAdminPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotAdminPassword", reflect.TypeOf((*MockService)(nil).ForgotAdminPassword), ctx, in)
}

// ForgotUserPassword mocks base method.
func (m *MockService) ForgotUserPassword(ctx context.Context, in *user.ForgotUserPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotUserPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotUserPassword indicates an expected call of ForgotUserPassword.
func (mr *MockServiceMockRecorder) ForgotUserPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotUserPassword", reflect.TypeOf((*MockService)(nil).ForgotUserPassword), ctx, in)
}

// GetAdmin mocks base method.
func (m *MockService) GetAdmin(ctx context.Context, in *user.GetAdminInput) (*entity.Admin, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", ctx, in)
	ret0, _ := ret[0].(*entity.Admin)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockServiceMockRecorder) GetAdmin(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockService)(nil).GetAdmin), ctx, in)
}

// GetAdminAuth mocks base method.
func (m *MockService) GetAdminAuth(ctx context.Context, in *user.GetAdminAuthInput) (*entity.AdminAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminAuth", ctx, in)
	ret0, _ := ret[0].(*entity.AdminAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminAuth indicates an expected call of GetAdminAuth.
func (mr *MockServiceMockRecorder) GetAdminAuth(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminAuth", reflect.TypeOf((*MockService)(nil).GetAdminAuth), ctx, in)
}

// GetAdministrator mocks base method.
func (m *MockService) GetAdministrator(ctx context.Context, in *user.GetAdministratorInput) (*entity.Administrator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdministrator", ctx, in)
	ret0, _ := ret[0].(*entity.Administrator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdministrator indicates an expected call of GetAdministrator.
func (mr *MockServiceMockRecorder) GetAdministrator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdministrator", reflect.TypeOf((*MockService)(nil).GetAdministrator), ctx, in)
}

// GetCoordinator mocks base method.
func (m *MockService) GetCoordinator(ctx context.Context, in *user.GetCoordinatorInput) (*entity.Coordinator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCoordinator", ctx, in)
	ret0, _ := ret[0].(*entity.Coordinator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCoordinator indicates an expected call of GetCoordinator.
func (mr *MockServiceMockRecorder) GetCoordinator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCoordinator", reflect.TypeOf((*MockService)(nil).GetCoordinator), ctx, in)
}

// GetProducer mocks base method.
func (m *MockService) GetProducer(ctx context.Context, in *user.GetProducerInput) (*entity.Producer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducer", ctx, in)
	ret0, _ := ret[0].(*entity.Producer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducer indicates an expected call of GetProducer.
func (mr *MockServiceMockRecorder) GetProducer(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducer", reflect.TypeOf((*MockService)(nil).GetProducer), ctx, in)
}

// GetUser mocks base method.
func (m *MockService) GetUser(ctx context.Context, in *user.GetUserInput) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, in)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockServiceMockRecorder) GetUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), ctx, in)
}

// GetUserAuth mocks base method.
func (m *MockService) GetUserAuth(ctx context.Context, in *user.GetUserAuthInput) (*entity.UserAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserAuth", ctx, in)
	ret0, _ := ret[0].(*entity.UserAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserAuth indicates an expected call of GetUserAuth.
func (mr *MockServiceMockRecorder) GetUserAuth(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserAuth", reflect.TypeOf((*MockService)(nil).GetUserAuth), ctx, in)
}

// InitializeUser mocks base method.
func (m *MockService) InitializeUser(ctx context.Context, in *user.InitializeUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeUser", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeUser indicates an expected call of InitializeUser.
func (mr *MockServiceMockRecorder) InitializeUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeUser", reflect.TypeOf((*MockService)(nil).InitializeUser), ctx, in)
}

// ListAdministrators mocks base method.
func (m *MockService) ListAdministrators(ctx context.Context, in *user.ListAdministratorsInput) (entity.Administrators, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAdministrators", ctx, in)
	ret0, _ := ret[0].(entity.Administrators)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListAdministrators indicates an expected call of ListAdministrators.
func (mr *MockServiceMockRecorder) ListAdministrators(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAdministrators", reflect.TypeOf((*MockService)(nil).ListAdministrators), ctx, in)
}

// ListCoordinators mocks base method.
func (m *MockService) ListCoordinators(ctx context.Context, in *user.ListCoordinatorsInput) (entity.Coordinators, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCoordinators", ctx, in)
	ret0, _ := ret[0].(entity.Coordinators)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListCoordinators indicates an expected call of ListCoordinators.
func (mr *MockServiceMockRecorder) ListCoordinators(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCoordinators", reflect.TypeOf((*MockService)(nil).ListCoordinators), ctx, in)
}

// ListProducers mocks base method.
func (m *MockService) ListProducers(ctx context.Context, in *user.ListProducersInput) (entity.Producers, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducers", ctx, in)
	ret0, _ := ret[0].(entity.Producers)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProducers indicates an expected call of ListProducers.
func (mr *MockServiceMockRecorder) ListProducers(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducers", reflect.TypeOf((*MockService)(nil).ListProducers), ctx, in)
}

// ListUsers mocks base method.
func (m *MockService) ListUsers(ctx context.Context, in *user.ListUsersInput) (entity.Users, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListUsers", ctx, in)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListUsers indicates an expected call of ListUsers.
func (mr *MockServiceMockRecorder) ListUsers(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListUsers", reflect.TypeOf((*MockService)(nil).ListUsers), ctx, in)
}

// MultiGetAdminDevices mocks base method.
func (m *MockService) MultiGetAdminDevices(ctx context.Context, in *user.MultiGetAdminDevicesInput) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetAdminDevices", ctx, in)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetAdminDevices indicates an expected call of MultiGetAdminDevices.
func (mr *MockServiceMockRecorder) MultiGetAdminDevices(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetAdminDevices", reflect.TypeOf((*MockService)(nil).MultiGetAdminDevices), ctx, in)
}

// MultiGetAdministrators mocks base method.
func (m *MockService) MultiGetAdministrators(ctx context.Context, in *user.MultiGetAdministratorsInput) (entity.Administrators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetAdministrators", ctx, in)
	ret0, _ := ret[0].(entity.Administrators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetAdministrators indicates an expected call of MultiGetAdministrators.
func (mr *MockServiceMockRecorder) MultiGetAdministrators(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetAdministrators", reflect.TypeOf((*MockService)(nil).MultiGetAdministrators), ctx, in)
}

// MultiGetAdmins mocks base method.
func (m *MockService) MultiGetAdmins(ctx context.Context, in *user.MultiGetAdminsInput) (entity.Admins, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetAdmins", ctx, in)
	ret0, _ := ret[0].(entity.Admins)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetAdmins indicates an expected call of MultiGetAdmins.
func (mr *MockServiceMockRecorder) MultiGetAdmins(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetAdmins", reflect.TypeOf((*MockService)(nil).MultiGetAdmins), ctx, in)
}

// MultiGetCoordinators mocks base method.
func (m *MockService) MultiGetCoordinators(ctx context.Context, in *user.MultiGetCoordinatorsInput) (entity.Coordinators, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetCoordinators", ctx, in)
	ret0, _ := ret[0].(entity.Coordinators)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetCoordinators indicates an expected call of MultiGetCoordinators.
func (mr *MockServiceMockRecorder) MultiGetCoordinators(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetCoordinators", reflect.TypeOf((*MockService)(nil).MultiGetCoordinators), ctx, in)
}

// MultiGetProducers mocks base method.
func (m *MockService) MultiGetProducers(ctx context.Context, in *user.MultiGetProducersInput) (entity.Producers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetProducers", ctx, in)
	ret0, _ := ret[0].(entity.Producers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetProducers indicates an expected call of MultiGetProducers.
func (mr *MockServiceMockRecorder) MultiGetProducers(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetProducers", reflect.TypeOf((*MockService)(nil).MultiGetProducers), ctx, in)
}

// MultiGetUserDevices mocks base method.
func (m *MockService) MultiGetUserDevices(ctx context.Context, in *user.MultiGetUserDevicesInput) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetUserDevices", ctx, in)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetUserDevices indicates an expected call of MultiGetUserDevices.
func (mr *MockServiceMockRecorder) MultiGetUserDevices(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetUserDevices", reflect.TypeOf((*MockService)(nil).MultiGetUserDevices), ctx, in)
}

// MultiGetUsers mocks base method.
func (m *MockService) MultiGetUsers(ctx context.Context, in *user.MultiGetUsersInput) (entity.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetUsers", ctx, in)
	ret0, _ := ret[0].(entity.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetUsers indicates an expected call of MultiGetUsers.
func (mr *MockServiceMockRecorder) MultiGetUsers(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetUsers", reflect.TypeOf((*MockService)(nil).MultiGetUsers), ctx, in)
}

// RefreshAdminToken mocks base method.
func (m *MockService) RefreshAdminToken(ctx context.Context, in *user.RefreshAdminTokenInput) (*entity.AdminAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshAdminToken", ctx, in)
	ret0, _ := ret[0].(*entity.AdminAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshAdminToken indicates an expected call of RefreshAdminToken.
func (mr *MockServiceMockRecorder) RefreshAdminToken(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshAdminToken", reflect.TypeOf((*MockService)(nil).RefreshAdminToken), ctx, in)
}

// RefreshUserToken mocks base method.
func (m *MockService) RefreshUserToken(ctx context.Context, in *user.RefreshUserTokenInput) (*entity.UserAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshUserToken", ctx, in)
	ret0, _ := ret[0].(*entity.UserAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshUserToken indicates an expected call of RefreshUserToken.
func (mr *MockServiceMockRecorder) RefreshUserToken(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshUserToken", reflect.TypeOf((*MockService)(nil).RefreshUserToken), ctx, in)
}

// RegisterAdminDevice mocks base method.
func (m *MockService) RegisterAdminDevice(ctx context.Context, in *user.RegisterAdminDeviceInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAdminDevice", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterAdminDevice indicates an expected call of RegisterAdminDevice.
func (mr *MockServiceMockRecorder) RegisterAdminDevice(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAdminDevice", reflect.TypeOf((*MockService)(nil).RegisterAdminDevice), ctx, in)
}

// RelateProducers mocks base method.
func (m *MockService) RelateProducers(ctx context.Context, in *user.RelateProducersInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelateProducers", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// RelateProducers indicates an expected call of RelateProducers.
func (mr *MockServiceMockRecorder) RelateProducers(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelateProducers", reflect.TypeOf((*MockService)(nil).RelateProducers), ctx, in)
}

// ResetAdministratorPassword mocks base method.
func (m *MockService) ResetAdministratorPassword(ctx context.Context, in *user.ResetAdministratorPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetAdministratorPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetAdministratorPassword indicates an expected call of ResetAdministratorPassword.
func (mr *MockServiceMockRecorder) ResetAdministratorPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetAdministratorPassword", reflect.TypeOf((*MockService)(nil).ResetAdministratorPassword), ctx, in)
}

// ResetCoordinatorPassword mocks base method.
func (m *MockService) ResetCoordinatorPassword(ctx context.Context, in *user.ResetCoordinatorPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetCoordinatorPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetCoordinatorPassword indicates an expected call of ResetCoordinatorPassword.
func (mr *MockServiceMockRecorder) ResetCoordinatorPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetCoordinatorPassword", reflect.TypeOf((*MockService)(nil).ResetCoordinatorPassword), ctx, in)
}

// ResetProducerPassword mocks base method.
func (m *MockService) ResetProducerPassword(ctx context.Context, in *user.ResetProducerPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResetProducerPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetProducerPassword indicates an expected call of ResetProducerPassword.
func (mr *MockServiceMockRecorder) ResetProducerPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetProducerPassword", reflect.TypeOf((*MockService)(nil).ResetProducerPassword), ctx, in)
}

// SignInAdmin mocks base method.
func (m *MockService) SignInAdmin(ctx context.Context, in *user.SignInAdminInput) (*entity.AdminAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInAdmin", ctx, in)
	ret0, _ := ret[0].(*entity.AdminAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInAdmin indicates an expected call of SignInAdmin.
func (mr *MockServiceMockRecorder) SignInAdmin(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInAdmin", reflect.TypeOf((*MockService)(nil).SignInAdmin), ctx, in)
}

// SignInUser mocks base method.
func (m *MockService) SignInUser(ctx context.Context, in *user.SignInUserInput) (*entity.UserAuth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignInUser", ctx, in)
	ret0, _ := ret[0].(*entity.UserAuth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignInUser indicates an expected call of SignInUser.
func (mr *MockServiceMockRecorder) SignInUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignInUser", reflect.TypeOf((*MockService)(nil).SignInUser), ctx, in)
}

// SignOutAdmin mocks base method.
func (m *MockService) SignOutAdmin(ctx context.Context, in *user.SignOutAdminInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignOutAdmin", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignOutAdmin indicates an expected call of SignOutAdmin.
func (mr *MockServiceMockRecorder) SignOutAdmin(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignOutAdmin", reflect.TypeOf((*MockService)(nil).SignOutAdmin), ctx, in)
}

// SignOutUser mocks base method.
func (m *MockService) SignOutUser(ctx context.Context, in *user.SignOutUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignOutUser", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignOutUser indicates an expected call of SignOutUser.
func (mr *MockServiceMockRecorder) SignOutUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignOutUser", reflect.TypeOf((*MockService)(nil).SignOutUser), ctx, in)
}

// UnrelateProducer mocks base method.
func (m *MockService) UnrelateProducer(ctx context.Context, in *user.UnrelateProducerInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnrelateProducer", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnrelateProducer indicates an expected call of UnrelateProducer.
func (mr *MockServiceMockRecorder) UnrelateProducer(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnrelateProducer", reflect.TypeOf((*MockService)(nil).UnrelateProducer), ctx, in)
}

// UpdateAdminEmail mocks base method.
func (m *MockService) UpdateAdminEmail(ctx context.Context, in *user.UpdateAdminEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminEmail indicates an expected call of UpdateAdminEmail.
func (mr *MockServiceMockRecorder) UpdateAdminEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminEmail", reflect.TypeOf((*MockService)(nil).UpdateAdminEmail), ctx, in)
}

// UpdateAdminPassword mocks base method.
func (m *MockService) UpdateAdminPassword(ctx context.Context, in *user.UpdateAdminPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdminPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdminPassword indicates an expected call of UpdateAdminPassword.
func (mr *MockServiceMockRecorder) UpdateAdminPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdminPassword", reflect.TypeOf((*MockService)(nil).UpdateAdminPassword), ctx, in)
}

// UpdateAdministrator mocks base method.
func (m *MockService) UpdateAdministrator(ctx context.Context, in *user.UpdateAdministratorInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdministrator", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdministrator indicates an expected call of UpdateAdministrator.
func (mr *MockServiceMockRecorder) UpdateAdministrator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdministrator", reflect.TypeOf((*MockService)(nil).UpdateAdministrator), ctx, in)
}

// UpdateAdministratorEmail mocks base method.
func (m *MockService) UpdateAdministratorEmail(ctx context.Context, in *user.UpdateAdministratorEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAdministratorEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAdministratorEmail indicates an expected call of UpdateAdministratorEmail.
func (mr *MockServiceMockRecorder) UpdateAdministratorEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdministratorEmail", reflect.TypeOf((*MockService)(nil).UpdateAdministratorEmail), ctx, in)
}

// UpdateCoordinator mocks base method.
func (m *MockService) UpdateCoordinator(ctx context.Context, in *user.UpdateCoordinatorInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCoordinator", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCoordinator indicates an expected call of UpdateCoordinator.
func (mr *MockServiceMockRecorder) UpdateCoordinator(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCoordinator", reflect.TypeOf((*MockService)(nil).UpdateCoordinator), ctx, in)
}

// UpdateCoordinatorEmail mocks base method.
func (m *MockService) UpdateCoordinatorEmail(ctx context.Context, in *user.UpdateCoordinatorEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCoordinatorEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCoordinatorEmail indicates an expected call of UpdateCoordinatorEmail.
func (mr *MockServiceMockRecorder) UpdateCoordinatorEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCoordinatorEmail", reflect.TypeOf((*MockService)(nil).UpdateCoordinatorEmail), ctx, in)
}

// UpdateCoordinatorHeaders mocks base method.
func (m *MockService) UpdateCoordinatorHeaders(ctx context.Context, in *user.UpdateCoordinatorHeadersInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCoordinatorHeaders", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCoordinatorHeaders indicates an expected call of UpdateCoordinatorHeaders.
func (mr *MockServiceMockRecorder) UpdateCoordinatorHeaders(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCoordinatorHeaders", reflect.TypeOf((*MockService)(nil).UpdateCoordinatorHeaders), ctx, in)
}

// UpdateCoordinatorThumbnails mocks base method.
func (m *MockService) UpdateCoordinatorThumbnails(ctx context.Context, in *user.UpdateCoordinatorThumbnailsInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCoordinatorThumbnails", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCoordinatorThumbnails indicates an expected call of UpdateCoordinatorThumbnails.
func (mr *MockServiceMockRecorder) UpdateCoordinatorThumbnails(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCoordinatorThumbnails", reflect.TypeOf((*MockService)(nil).UpdateCoordinatorThumbnails), ctx, in)
}

// UpdateProducer mocks base method.
func (m *MockService) UpdateProducer(ctx context.Context, in *user.UpdateProducerInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducer", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProducer indicates an expected call of UpdateProducer.
func (mr *MockServiceMockRecorder) UpdateProducer(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducer", reflect.TypeOf((*MockService)(nil).UpdateProducer), ctx, in)
}

// UpdateProducerEmail mocks base method.
func (m *MockService) UpdateProducerEmail(ctx context.Context, in *user.UpdateProducerEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducerEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProducerEmail indicates an expected call of UpdateProducerEmail.
func (mr *MockServiceMockRecorder) UpdateProducerEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducerEmail", reflect.TypeOf((*MockService)(nil).UpdateProducerEmail), ctx, in)
}

// UpdateProducerHeaders mocks base method.
func (m *MockService) UpdateProducerHeaders(ctx context.Context, in *user.UpdateProducerHeadersInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducerHeaders", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProducerHeaders indicates an expected call of UpdateProducerHeaders.
func (mr *MockServiceMockRecorder) UpdateProducerHeaders(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducerHeaders", reflect.TypeOf((*MockService)(nil).UpdateProducerHeaders), ctx, in)
}

// UpdateProducerThumbnails mocks base method.
func (m *MockService) UpdateProducerThumbnails(ctx context.Context, in *user.UpdateProducerThumbnailsInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProducerThumbnails", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProducerThumbnails indicates an expected call of UpdateProducerThumbnails.
func (mr *MockServiceMockRecorder) UpdateProducerThumbnails(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProducerThumbnails", reflect.TypeOf((*MockService)(nil).UpdateProducerThumbnails), ctx, in)
}

// UpdateUserEmail mocks base method.
func (m *MockService) UpdateUserEmail(ctx context.Context, in *user.UpdateUserEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserEmail indicates an expected call of UpdateUserEmail.
func (mr *MockServiceMockRecorder) UpdateUserEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserEmail", reflect.TypeOf((*MockService)(nil).UpdateUserEmail), ctx, in)
}

// UpdateUserPassword mocks base method.
func (m *MockService) UpdateUserPassword(ctx context.Context, in *user.UpdateUserPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockServiceMockRecorder) UpdateUserPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockService)(nil).UpdateUserPassword), ctx, in)
}

// VerifyAdminEmail mocks base method.
func (m *MockService) VerifyAdminEmail(ctx context.Context, in *user.VerifyAdminEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAdminEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyAdminEmail indicates an expected call of VerifyAdminEmail.
func (mr *MockServiceMockRecorder) VerifyAdminEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAdminEmail", reflect.TypeOf((*MockService)(nil).VerifyAdminEmail), ctx, in)
}

// VerifyAdminPassword mocks base method.
func (m *MockService) VerifyAdminPassword(ctx context.Context, in *user.VerifyAdminPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAdminPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyAdminPassword indicates an expected call of VerifyAdminPassword.
func (mr *MockServiceMockRecorder) VerifyAdminPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAdminPassword", reflect.TypeOf((*MockService)(nil).VerifyAdminPassword), ctx, in)
}

// VerifyUser mocks base method.
func (m *MockService) VerifyUser(ctx context.Context, in *user.VerifyUserInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUser", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUser indicates an expected call of VerifyUser.
func (mr *MockServiceMockRecorder) VerifyUser(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUser", reflect.TypeOf((*MockService)(nil).VerifyUser), ctx, in)
}

// VerifyUserEmail mocks base method.
func (m *MockService) VerifyUserEmail(ctx context.Context, in *user.VerifyUserEmailInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUserEmail", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUserEmail indicates an expected call of VerifyUserEmail.
func (mr *MockServiceMockRecorder) VerifyUserEmail(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUserEmail", reflect.TypeOf((*MockService)(nil).VerifyUserEmail), ctx, in)
}

// VerifyUserPassword mocks base method.
func (m *MockService) VerifyUserPassword(ctx context.Context, in *user.VerifyUserPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyUserPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// VerifyUserPassword indicates an expected call of VerifyUserPassword.
func (mr *MockServiceMockRecorder) VerifyUserPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyUserPassword", reflect.TypeOf((*MockService)(nil).VerifyUserPassword), ctx, in)
}
