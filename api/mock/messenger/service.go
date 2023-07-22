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

// CreateContact mocks base method.
func (m *MockService) CreateContact(ctx context.Context, in *messenger.CreateContactInput) (*entity.Contact, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContact", ctx, in)
	ret0, _ := ret[0].(*entity.Contact)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContact indicates an expected call of CreateContact.
func (mr *MockServiceMockRecorder) CreateContact(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContact", reflect.TypeOf((*MockService)(nil).CreateContact), ctx, in)
}

// CreateContactRead mocks base method.
func (m *MockService) CreateContactRead(ctx context.Context, in *messenger.CreateContactReadInput) (*entity.ContactRead, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateContactRead", ctx, in)
	ret0, _ := ret[0].(*entity.ContactRead)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateContactRead indicates an expected call of CreateContactRead.
func (mr *MockServiceMockRecorder) CreateContactRead(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContactRead", reflect.TypeOf((*MockService)(nil).CreateContactRead), ctx, in)
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

// CreateThread mocks base method.
func (m *MockService) CreateThread(ctx context.Context, in *messenger.CreateThreadInput) (*entity.Thread, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateThread", ctx, in)
	ret0, _ := ret[0].(*entity.Thread)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateThread indicates an expected call of CreateThread.
func (mr *MockServiceMockRecorder) CreateThread(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateThread", reflect.TypeOf((*MockService)(nil).CreateThread), ctx, in)
}

// DeleteContact mocks base method.
func (m *MockService) DeleteContact(ctx context.Context, in *messenger.DeleteContactInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContact", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContact indicates an expected call of DeleteContact.
func (mr *MockServiceMockRecorder) DeleteContact(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContact", reflect.TypeOf((*MockService)(nil).DeleteContact), ctx, in)
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

// DeleteThread mocks base method.
func (m *MockService) DeleteThread(ctx context.Context, in *messenger.DeleteThreadInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteThread", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteThread indicates an expected call of DeleteThread.
func (mr *MockServiceMockRecorder) DeleteThread(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteThread", reflect.TypeOf((*MockService)(nil).DeleteThread), ctx, in)
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

// GetContactRead mocks base method.
func (m *MockService) GetContactRead(ctx context.Context, in *messenger.GetContactReadInput) (*entity.ContactRead, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContactRead", ctx, in)
	ret0, _ := ret[0].(*entity.ContactRead)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContactRead indicates an expected call of GetContactRead.
func (mr *MockServiceMockRecorder) GetContactRead(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContactRead", reflect.TypeOf((*MockService)(nil).GetContactRead), ctx, in)
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

// ListContactCategories mocks base method.
func (m *MockService) ListContactCategories(ctx context.Context, in *messenger.ListContactCategoriesInput) (entity.ContactCategories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContactCategories", ctx, in)
	ret0, _ := ret[0].(entity.ContactCategories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListContactCategories indicates an expected call of ListContactCategories.
func (mr *MockServiceMockRecorder) ListContactCategories(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContactCategories", reflect.TypeOf((*MockService)(nil).ListContactCategories), ctx, in)
}

// ListContacts mocks base method.
func (m *MockService) ListContacts(ctx context.Context, in *messenger.ListContactsInput) (entity.Contacts, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListContacts", ctx, in)
	ret0, _ := ret[0].(entity.Contacts)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListContacts indicates an expected call of ListContacts.
func (mr *MockServiceMockRecorder) ListContacts(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContacts", reflect.TypeOf((*MockService)(nil).ListContacts), ctx, in)
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

// ListThreadsByContactID mocks base method.
func (m *MockService) ListThreadsByContactID(ctx context.Context, in *messenger.ListThreadsByContactIDInput) (entity.Threads, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListThreadsByContactID", ctx, in)
	ret0, _ := ret[0].(entity.Threads)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListThreadsByContactID indicates an expected call of ListThreadsByContactID.
func (mr *MockServiceMockRecorder) ListThreadsByContactID(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListThreadsByContactID", reflect.TypeOf((*MockService)(nil).ListThreadsByContactID), ctx, in)
}

// MultiGetContactCategories mocks base method.
func (m *MockService) MultiGetContactCategories(ctx context.Context, in *messenger.MultiGetContactCategoriesInput) (entity.ContactCategories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetContactCategories", ctx, in)
	ret0, _ := ret[0].(entity.ContactCategories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetContactCategories indicates an expected call of MultiGetContactCategories.
func (mr *MockServiceMockRecorder) MultiGetContactCategories(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetContactCategories", reflect.TypeOf((*MockService)(nil).MultiGetContactCategories), ctx, in)
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

// UpdateContact mocks base method.
func (m *MockService) UpdateContact(ctx context.Context, in *messenger.UpdateContactInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContact", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContact indicates an expected call of UpdateContact.
func (mr *MockServiceMockRecorder) UpdateContact(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContact", reflect.TypeOf((*MockService)(nil).UpdateContact), ctx, in)
}

// UpdateContactReadFlag mocks base method.
func (m *MockService) UpdateContactReadFlag(ctx context.Context, in *messenger.UpdateContactReadFlagInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContactReadFlag", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContactReadFlag indicates an expected call of UpdateContactReadFlag.
func (mr *MockServiceMockRecorder) UpdateContactReadFlag(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContactReadFlag", reflect.TypeOf((*MockService)(nil).UpdateContactReadFlag), ctx, in)
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

// UpdateThread mocks base method.
func (m *MockService) UpdateThread(ctx context.Context, in *messenger.UpdateThreadInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateThread", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateThread indicates an expected call of UpdateThread.
func (mr *MockServiceMockRecorder) UpdateThread(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateThread", reflect.TypeOf((*MockService)(nil).UpdateThread), ctx, in)
}
