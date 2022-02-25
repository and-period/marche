// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_cognito is a generated GoMock package.
package mock_cognito

import (
	context "context"
	reflect "reflect"

	cognito "github.com/and-period/marche/api/pkg/cognito"
	gomock "github.com/golang/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ChangeEmail mocks base method.
func (m *MockClient) ChangeEmail(ctx context.Context, params *cognito.ChangeEmailParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeEmail", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeEmail indicates an expected call of ChangeEmail.
func (mr *MockClientMockRecorder) ChangeEmail(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEmail", reflect.TypeOf((*MockClient)(nil).ChangeEmail), ctx, params)
}

// ChangePassword mocks base method.
func (m *MockClient) ChangePassword(ctx context.Context, params *cognito.ChangePasswordParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockClientMockRecorder) ChangePassword(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockClient)(nil).ChangePassword), ctx, params)
}

// ConfirmChangeEmail mocks base method.
func (m *MockClient) ConfirmChangeEmail(ctx context.Context, params *cognito.ConfirmChangeEmailParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmChangeEmail", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfirmChangeEmail indicates an expected call of ConfirmChangeEmail.
func (mr *MockClientMockRecorder) ConfirmChangeEmail(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmChangeEmail", reflect.TypeOf((*MockClient)(nil).ConfirmChangeEmail), ctx, params)
}

// ConfirmForgotPassword mocks base method.
func (m *MockClient) ConfirmForgotPassword(ctx context.Context, params *cognito.ConfirmForgotPasswordParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmForgotPassword", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfirmForgotPassword indicates an expected call of ConfirmForgotPassword.
func (mr *MockClientMockRecorder) ConfirmForgotPassword(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmForgotPassword", reflect.TypeOf((*MockClient)(nil).ConfirmForgotPassword), ctx, params)
}

// ConfirmSignUp mocks base method.
func (m *MockClient) ConfirmSignUp(ctx context.Context, username, verifyCode string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmSignUp", ctx, username, verifyCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConfirmSignUp indicates an expected call of ConfirmSignUp.
func (mr *MockClientMockRecorder) ConfirmSignUp(ctx, username, verifyCode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmSignUp", reflect.TypeOf((*MockClient)(nil).ConfirmSignUp), ctx, username, verifyCode)
}

// DeleteUser mocks base method.
func (m *MockClient) DeleteUser(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockClientMockRecorder) DeleteUser(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockClient)(nil).DeleteUser), ctx, username)
}

// ForgotPassword mocks base method.
func (m *MockClient) ForgotPassword(ctx context.Context, username string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForgotPassword", ctx, username)
	ret0, _ := ret[0].(error)
	return ret0
}

// ForgotPassword indicates an expected call of ForgotPassword.
func (mr *MockClientMockRecorder) ForgotPassword(ctx, username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForgotPassword", reflect.TypeOf((*MockClient)(nil).ForgotPassword), ctx, username)
}

// GetUsername mocks base method.
func (m *MockClient) GetUsername(ctx context.Context, accessToken string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsername", ctx, accessToken)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsername indicates an expected call of GetUsername.
func (mr *MockClientMockRecorder) GetUsername(ctx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsername", reflect.TypeOf((*MockClient)(nil).GetUsername), ctx, accessToken)
}

// RefreshToken mocks base method.
func (m *MockClient) RefreshToken(ctx context.Context, refreshToken string) (*cognito.AuthResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", ctx, refreshToken)
	ret0, _ := ret[0].(*cognito.AuthResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken.
func (mr *MockClientMockRecorder) RefreshToken(ctx, refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockClient)(nil).RefreshToken), ctx, refreshToken)
}

// SignIn mocks base method.
func (m *MockClient) SignIn(ctx context.Context, username, password string) (*cognito.AuthResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, username, password)
	ret0, _ := ret[0].(*cognito.AuthResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockClientMockRecorder) SignIn(ctx, username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockClient)(nil).SignIn), ctx, username, password)
}

// SignOut mocks base method.
func (m *MockClient) SignOut(ctx context.Context, accessToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignOut", ctx, accessToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignOut indicates an expected call of SignOut.
func (mr *MockClientMockRecorder) SignOut(ctx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignOut", reflect.TypeOf((*MockClient)(nil).SignOut), ctx, accessToken)
}

// SignUp mocks base method.
func (m *MockClient) SignUp(ctx context.Context, params *cognito.SignUpParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockClientMockRecorder) SignUp(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockClient)(nil).SignUp), ctx, params)
}
