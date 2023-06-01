// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_messenger is a generated GoMock package.
package mock_messenger

import (
	context "context"
	reflect "reflect"

	messenger "github.com/and-period/furumaru/api/internal/messenger"
	entity "github.com/and-period/furumaru/api/internal/messenger/entity"
	gomock "github.com/golang/mock/gomock"
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

// CreateNotification mocks base method.
func (m *MockService) CreateNotification(ctx context.Context, in *messenger.CreateNotificationInput) (*entity.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNotification", ctx, in)
	ret0, _ := ret[0].(*entity.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNotification indicates an expected call of CreateNotification.
func (mr *MockServiceMockRecorder) CreateNotification(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNotification", reflect.TypeOf((*MockService)(nil).CreateNotification), ctx, in)
}

// DeleteNotification mocks base method.
func (m *MockService) DeleteNotification(ctx context.Context, in *messenger.DeleteNotificationInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNotification", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNotification indicates an expected call of DeleteNotification.
func (mr *MockServiceMockRecorder) DeleteNotification(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNotification", reflect.TypeOf((*MockService)(nil).DeleteNotification), ctx, in)
}

// GetContact mocks base method.
func (m *MockService) GetContact(ctx context.Context, in *messenger.GetContactInput) (*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContact", ctx, in)
	ret0, _ := ret[0].(*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContact indicates an expected call of GetContact.
func (mr *MockServiceMockRecorder) GetContact(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContact", reflect.TypeOf((*MockService)(nil).GetContact), ctx, in)
}

// GetContactCategory mocks base method.
func (m *MockService) GetContactCategory(ctx context.Context, in *messenger.GetContactCategoryInput) (*entity.ContactCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactCategory", ctx, in)
	ret0, _ := ret[0].(*entity.ContactCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactCategory indicates an expected call of GetContactCategory.
func (mr *MockServiceMockRecorder) GetContactCategory(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactCategory", reflect.TypeOf((*MockService)(nil).GetContactCategory), ctx, in)
}

// GetMessage mocks base method.
func (m *MockService) GetMessage(ctx context.Context, in *messenger.GetMessageInput) (*entity.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage", ctx, in)
	ret0, _ := ret[0].(*entity.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockServiceMockRecorder) GetMessage(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockService)(nil).GetMessage), ctx, in)
}

// GetNotification mocks base method.
func (m *MockService) GetNotification(ctx context.Context, in *messenger.GetNotificationInput) (*entity.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotification", ctx, in)
	ret0, _ := ret[0].(*entity.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotification indicates an expected call of GetNotification.
func (mr *MockServiceMockRecorder) GetNotification(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotification", reflect.TypeOf((*MockService)(nil).GetNotification), ctx, in)
}

// GetThread mocks base method.
func (m *MockService) GetThread(ctx context.Context, in *messenger.GetThreadInput) (*entity.Thread, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThread", ctx, in)
	ret0, _ := ret[0].(*entity.Thread)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetThread indicates an expected call of GetThread.
func (mr *MockServiceMockRecorder) GetThread(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThread", reflect.TypeOf((*MockService)(nil).GetThread), ctx, in)
}

// ListMessages mocks base method.
func (m *MockService) ListMessages(ctx context.Context, in *messenger.ListMessagesInput) (entity.Messages, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMessages", ctx, in)
	ret0, _ := ret[0].(entity.Messages)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMessages indicates an expected call of ListMessages.
func (mr *MockServiceMockRecorder) ListMessages(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMessages", reflect.TypeOf((*MockService)(nil).ListMessages), ctx, in)
}

// ListNotifications mocks base method.
func (m *MockService) ListNotifications(ctx context.Context, in *messenger.ListNotificationsInput) (entity.Notifications, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNotifications", ctx, in)
	ret0, _ := ret[0].(entity.Notifications)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListNotifications indicates an expected call of ListNotifications.
func (mr *MockServiceMockRecorder) ListNotifications(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNotifications", reflect.TypeOf((*MockService)(nil).ListNotifications), ctx, in)
}

// NotifyNotification mocks base method.
func (m *MockService) NotifyNotification(ctx context.Context, in *messenger.NotifyNotificationInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyNotification", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyNotification indicates an expected call of NotifyNotification.
func (mr *MockServiceMockRecorder) NotifyNotification(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyNotification", reflect.TypeOf((*MockService)(nil).NotifyNotification), ctx, in)
}

// NotifyRegisterAdmin mocks base method.
func (m *MockService) NotifyRegisterAdmin(ctx context.Context, in *messenger.NotifyRegisterAdminInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyRegisterAdmin", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyRegisterAdmin indicates an expected call of NotifyRegisterAdmin.
func (mr *MockServiceMockRecorder) NotifyRegisterAdmin(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyRegisterAdmin", reflect.TypeOf((*MockService)(nil).NotifyRegisterAdmin), ctx, in)
}

// NotifyResetAdminPassword mocks base method.
func (m *MockService) NotifyResetAdminPassword(ctx context.Context, in *messenger.NotifyResetAdminPasswordInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyResetAdminPassword", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyResetAdminPassword indicates an expected call of NotifyResetAdminPassword.
func (mr *MockServiceMockRecorder) NotifyResetAdminPassword(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyResetAdminPassword", reflect.TypeOf((*MockService)(nil).NotifyResetAdminPassword), ctx, in)
}

// UpdateNotification mocks base method.
func (m *MockService) UpdateNotification(ctx context.Context, in *messenger.UpdateNotificationInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNotification", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNotification indicates an expected call of UpdateNotification.
func (mr *MockServiceMockRecorder) UpdateNotification(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotification", reflect.TypeOf((*MockService)(nil).UpdateNotification), ctx, in)
}
