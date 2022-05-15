// Code generated by MockGen. DO NOT EDIT.
// Source: client.go

// Package mock_mailer is a generated GoMock package.
package mock_mailer

import (
	context "context"
	reflect "reflect"

	mailer "github.com/and-period/marche/api/pkg/mailer"
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

// MultiSend mocks base method.
func (m *MockClient) MultiSend(ctx context.Context, emailID, fromName, fromAddress string, ps []*mailer.Personalization) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiSend", ctx, emailID, fromName, fromAddress, ps)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultiSend indicates an expected call of MultiSend.
func (mr *MockClientMockRecorder) MultiSend(ctx, emailID, fromName, fromAddress, ps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiSend", reflect.TypeOf((*MockClient)(nil).MultiSend), ctx, emailID, fromName, fromAddress, ps)
}

// MultiSendFromInfo mocks base method.
func (m *MockClient) MultiSendFromInfo(ctx context.Context, emailID string, ps []*mailer.Personalization) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiSendFromInfo", ctx, emailID, ps)
	ret0, _ := ret[0].(error)
	return ret0
}

// MultiSendFromInfo indicates an expected call of MultiSendFromInfo.
func (mr *MockClientMockRecorder) MultiSendFromInfo(ctx, emailID, ps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiSendFromInfo", reflect.TypeOf((*MockClient)(nil).MultiSendFromInfo), ctx, emailID, ps)
}

// SendFromInfo mocks base method.
func (m *MockClient) SendFromInfo(ctx context.Context, emailID, toName, toAddress string, substitutions map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendFromInfo", ctx, emailID, toName, toAddress, substitutions)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendFromInfo indicates an expected call of SendFromInfo.
func (mr *MockClientMockRecorder) SendFromInfo(ctx, emailID, toName, toAddress, substitutions interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendFromInfo", reflect.TypeOf((*MockClient)(nil).SendFromInfo), ctx, emailID, toName, toAddress, substitutions)
}
