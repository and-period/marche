// Code generated by MockGen. DO NOT EDIT.
// Source: database.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	common "github.com/and-period/furumaru/api/internal/common"
	database "github.com/and-period/furumaru/api/internal/store/database"
	entity "github.com/and-period/furumaru/api/internal/store/entity"
	gomock "github.com/golang/mock/gomock"
)

// MockAddress is a mock of Address interface.
type MockAddress struct {
	ctrl     *gomock.Controller
	recorder *MockAddressMockRecorder
}

// MockAddressMockRecorder is the mock recorder for MockAddress.
type MockAddressMockRecorder struct {
	mock *MockAddress
}

// NewMockAddress creates a new mock instance.
func NewMockAddress(ctrl *gomock.Controller) *MockAddress {
	mock := &MockAddress{ctrl: ctrl}
	mock.recorder = &MockAddressMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAddress) EXPECT() *MockAddressMockRecorder {
	return m.recorder
}

// MultiGet mocks base method.
func (m *MockAddress) MultiGet(ctx context.Context, addressIDs []string, fields ...string) (entity.Addresses, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, addressIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.Addresses)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockAddressMockRecorder) MultiGet(ctx, addressIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, addressIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockAddress)(nil).MultiGet), varargs...)
}

// MockCategory is a mock of Category interface.
type MockCategory struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryMockRecorder
}

// MockCategoryMockRecorder is the mock recorder for MockCategory.
type MockCategoryMockRecorder struct {
	mock *MockCategory
}

// NewMockCategory creates a new mock instance.
func NewMockCategory(ctrl *gomock.Controller) *MockCategory {
	mock := &MockCategory{ctrl: ctrl}
	mock.recorder = &MockCategoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategory) EXPECT() *MockCategoryMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockCategory) Count(ctx context.Context, params *database.ListCategoriesParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockCategoryMockRecorder) Count(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockCategory)(nil).Count), ctx, params)
}

// Create mocks base method.
func (m *MockCategory) Create(ctx context.Context, category *entity.Category) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, category)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockCategoryMockRecorder) Create(ctx, category interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCategory)(nil).Create), ctx, category)
}

// Delete mocks base method.
func (m *MockCategory) Delete(ctx context.Context, categoryID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, categoryID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCategoryMockRecorder) Delete(ctx, categoryID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCategory)(nil).Delete), ctx, categoryID)
}

// Get mocks base method.
func (m *MockCategory) Get(ctx context.Context, categoryID string, fields ...string) (*entity.Category, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, categoryID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCategoryMockRecorder) Get(ctx, categoryID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, categoryID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCategory)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockCategory) List(ctx context.Context, params *database.ListCategoriesParams, fields ...string) (entity.Categories, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCategoryMockRecorder) List(ctx, params interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCategory)(nil).List), varargs...)
}

// MultiGet mocks base method.
func (m *MockCategory) MultiGet(ctx context.Context, categoryIDs []string, fields ...string) (entity.Categories, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, categoryIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockCategoryMockRecorder) MultiGet(ctx, categoryIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, categoryIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockCategory)(nil).MultiGet), varargs...)
}

// Update mocks base method.
func (m *MockCategory) Update(ctx context.Context, categoryID, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, categoryID, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCategoryMockRecorder) Update(ctx, categoryID, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCategory)(nil).Update), ctx, categoryID, name)
}

// MockOrder is a mock of Order interface.
type MockOrder struct {
	ctrl     *gomock.Controller
	recorder *MockOrderMockRecorder
}

// MockOrderMockRecorder is the mock recorder for MockOrder.
type MockOrderMockRecorder struct {
	mock *MockOrder
}

// NewMockOrder creates a new mock instance.
func NewMockOrder(ctrl *gomock.Controller) *MockOrder {
	mock := &MockOrder{ctrl: ctrl}
	mock.recorder = &MockOrderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrder) EXPECT() *MockOrderMockRecorder {
	return m.recorder
}

// Aggregate mocks base method.
func (m *MockOrder) Aggregate(ctx context.Context, userIDs []string) (entity.AggregatedOrders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Aggregate", ctx, userIDs)
	ret0, _ := ret[0].(entity.AggregatedOrders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Aggregate indicates an expected call of Aggregate.
func (mr *MockOrderMockRecorder) Aggregate(ctx, userIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Aggregate", reflect.TypeOf((*MockOrder)(nil).Aggregate), ctx, userIDs)
}

// Count mocks base method.
func (m *MockOrder) Count(ctx context.Context, params *database.ListOrdersParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockOrderMockRecorder) Count(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOrder)(nil).Count), ctx, params)
}

// Get mocks base method.
func (m *MockOrder) Get(ctx context.Context, orderID string, fields ...string) (*entity.Order, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, orderID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOrderMockRecorder) Get(ctx, orderID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, orderID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOrder)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockOrder) List(ctx context.Context, params *database.ListOrdersParams, fields ...string) (entity.Orders, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockOrderMockRecorder) List(ctx, params interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockOrder)(nil).List), varargs...)
}

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockProduct) Count(ctx context.Context, params *database.ListProductsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockProductMockRecorder) Count(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockProduct)(nil).Count), ctx, params)
}

// Create mocks base method.
func (m *MockProduct) Create(ctx context.Context, product *entity.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProductMockRecorder) Create(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProduct)(nil).Create), ctx, product)
}

// Delete mocks base method.
func (m *MockProduct) Delete(ctx context.Context, productID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductMockRecorder) Delete(ctx, productID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProduct)(nil).Delete), ctx, productID)
}

// Get mocks base method.
func (m *MockProduct) Get(ctx context.Context, productID string, fields ...string) (*entity.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, productID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductMockRecorder) Get(ctx, productID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, productID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProduct)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockProduct) List(ctx context.Context, params *database.ListProductsParams, fields ...string) (entity.Products, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductMockRecorder) List(ctx, params interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProduct)(nil).List), varargs...)
}

// MultiGet mocks base method.
func (m *MockProduct) MultiGet(ctx context.Context, productIDs []string, fields ...string) (entity.Products, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, productIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.Products)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockProductMockRecorder) MultiGet(ctx, productIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, productIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockProduct)(nil).MultiGet), varargs...)
}

// Update mocks base method.
func (m *MockProduct) Update(ctx context.Context, productID string, params *database.UpdateProductParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, productID, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProductMockRecorder) Update(ctx, productID, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProduct)(nil).Update), ctx, productID, params)
}

// UpdateMedia mocks base method.
func (m *MockProduct) UpdateMedia(ctx context.Context, productID string, set func(entity.MultiProductMedia) bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMedia", ctx, productID, set)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMedia indicates an expected call of UpdateMedia.
func (mr *MockProductMockRecorder) UpdateMedia(ctx, productID, set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMedia", reflect.TypeOf((*MockProduct)(nil).UpdateMedia), ctx, productID, set)
}

// MockProductType is a mock of ProductType interface.
type MockProductType struct {
	ctrl     *gomock.Controller
	recorder *MockProductTypeMockRecorder
}

// MockProductTypeMockRecorder is the mock recorder for MockProductType.
type MockProductTypeMockRecorder struct {
	mock *MockProductType
}

// NewMockProductType creates a new mock instance.
func NewMockProductType(ctrl *gomock.Controller) *MockProductType {
	mock := &MockProductType{ctrl: ctrl}
	mock.recorder = &MockProductTypeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductType) EXPECT() *MockProductTypeMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockProductType) Count(ctx context.Context, params *database.ListProductTypesParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockProductTypeMockRecorder) Count(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockProductType)(nil).Count), ctx, params)
}

// Create mocks base method.
func (m *MockProductType) Create(ctx context.Context, productType *entity.ProductType) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, productType)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProductTypeMockRecorder) Create(ctx, productType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProductType)(nil).Create), ctx, productType)
}

// Delete mocks base method.
func (m *MockProductType) Delete(ctx context.Context, productTypeID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, productTypeID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockProductTypeMockRecorder) Delete(ctx, productTypeID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProductType)(nil).Delete), ctx, productTypeID)
}

// Get mocks base method.
func (m *MockProductType) Get(ctx context.Context, productTypeID string, fields ...string) (*entity.ProductType, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, productTypeID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.ProductType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProductTypeMockRecorder) Get(ctx, productTypeID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, productTypeID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductType)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockProductType) List(ctx context.Context, params *database.ListProductTypesParams, fields ...string) (entity.ProductTypes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.ProductTypes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockProductTypeMockRecorder) List(ctx, params interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockProductType)(nil).List), varargs...)
}

// MultiGet mocks base method.
func (m *MockProductType) MultiGet(ctx context.Context, productTypeIDs []string, fields ...string) (entity.ProductTypes, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, productTypeIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.ProductTypes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockProductTypeMockRecorder) MultiGet(ctx, productTypeIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, productTypeIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockProductType)(nil).MultiGet), varargs...)
}

// Update mocks base method.
func (m *MockProductType) Update(ctx context.Context, productTypeID, name, iconURL string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, productTypeID, name, iconURL)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockProductTypeMockRecorder) Update(ctx, productTypeID, name, iconURL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProductType)(nil).Update), ctx, productTypeID, name, iconURL)
}

// UpdateIcons mocks base method.
func (m *MockProductType) UpdateIcons(ctx context.Context, productTypeID string, icons common.Images) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateIcons", ctx, productTypeID, icons)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateIcons indicates an expected call of UpdateIcons.
func (mr *MockProductTypeMockRecorder) UpdateIcons(ctx, productTypeID, icons interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIcons", reflect.TypeOf((*MockProductType)(nil).UpdateIcons), ctx, productTypeID, icons)
}

// MockPromotion is a mock of Promotion interface.
type MockPromotion struct {
	ctrl     *gomock.Controller
	recorder *MockPromotionMockRecorder
}

// MockPromotionMockRecorder is the mock recorder for MockPromotion.
type MockPromotionMockRecorder struct {
	mock *MockPromotion
}

// NewMockPromotion creates a new mock instance.
func NewMockPromotion(ctrl *gomock.Controller) *MockPromotion {
	mock := &MockPromotion{ctrl: ctrl}
	mock.recorder = &MockPromotionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPromotion) EXPECT() *MockPromotionMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockPromotion) Count(ctx context.Context, params *database.ListPromotionsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockPromotionMockRecorder) Count(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockPromotion)(nil).Count), ctx, params)
}

// Create mocks base method.
func (m *MockPromotion) Create(ctx context.Context, promotion *entity.Promotion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, promotion)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPromotionMockRecorder) Create(ctx, promotion interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPromotion)(nil).Create), ctx, promotion)
}

// Delete mocks base method.
func (m *MockPromotion) Delete(ctx context.Context, promotionID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, promotionID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPromotionMockRecorder) Delete(ctx, promotionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPromotion)(nil).Delete), ctx, promotionID)
}

// Get mocks base method.
func (m *MockPromotion) Get(ctx context.Context, promotionID string, fields ...string) (*entity.Promotion, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, promotionID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Promotion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPromotionMockRecorder) Get(ctx, promotionID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, promotionID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPromotion)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockPromotion) List(ctx context.Context, params *database.ListPromotionsParams, fields ...string) (entity.Promotions, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Promotions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockPromotionMockRecorder) List(ctx, params interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockPromotion)(nil).List), varargs...)
}

// Update mocks base method.
func (m *MockPromotion) Update(ctx context.Context, promotionID string, params *database.UpdatePromotionParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, promotionID, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPromotionMockRecorder) Update(ctx, promotionID, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPromotion)(nil).Update), ctx, promotionID, params)
}

// MockRehearsal is a mock of Rehearsal interface.
type MockRehearsal struct {
	ctrl     *gomock.Controller
	recorder *MockRehearsalMockRecorder
}

// MockRehearsalMockRecorder is the mock recorder for MockRehearsal.
type MockRehearsalMockRecorder struct {
	mock *MockRehearsal
}

// NewMockRehearsal creates a new mock instance.
func NewMockRehearsal(ctrl *gomock.Controller) *MockRehearsal {
	mock := &MockRehearsal{ctrl: ctrl}
	mock.recorder = &MockRehearsalMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRehearsal) EXPECT() *MockRehearsalMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockRehearsal) Create(ctx context.Context, rehearsal *entity.Rehearsal) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, rehearsal)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockRehearsalMockRecorder) Create(ctx, rehearsal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRehearsal)(nil).Create), ctx, rehearsal)
}

// Get mocks base method.
func (m *MockRehearsal) Get(ctx context.Context, liveID string) (*entity.Rehearsal, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, liveID)
	ret0, _ := ret[0].(*entity.Rehearsal)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockRehearsalMockRecorder) Get(ctx, liveID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockRehearsal)(nil).Get), ctx, liveID)
}

// MockShipping is a mock of Shipping interface.
type MockShipping struct {
	ctrl     *gomock.Controller
	recorder *MockShippingMockRecorder
}

// MockShippingMockRecorder is the mock recorder for MockShipping.
type MockShippingMockRecorder struct {
	mock *MockShipping
}

// NewMockShipping creates a new mock instance.
func NewMockShipping(ctrl *gomock.Controller) *MockShipping {
	mock := &MockShipping{ctrl: ctrl}
	mock.recorder = &MockShippingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockShipping) EXPECT() *MockShippingMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockShipping) Count(ctx context.Context, params *database.ListShippingsParams) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, params)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockShippingMockRecorder) Count(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockShipping)(nil).Count), ctx, params)
}

// Create mocks base method.
func (m *MockShipping) Create(ctx context.Context, shipping *entity.Shipping) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, shipping)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockShippingMockRecorder) Create(ctx, shipping interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockShipping)(nil).Create), ctx, shipping)
}

// Delete mocks base method.
func (m *MockShipping) Delete(ctx context.Context, shippingID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, shippingID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockShippingMockRecorder) Delete(ctx, shippingID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockShipping)(nil).Delete), ctx, shippingID)
}

// Get mocks base method.
func (m *MockShipping) Get(ctx context.Context, shoppingID string, fields ...string) (*entity.Shipping, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, shoppingID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockShippingMockRecorder) Get(ctx, shoppingID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, shoppingID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockShipping)(nil).Get), varargs...)
}

// List mocks base method.
func (m *MockShipping) List(ctx context.Context, params *database.ListShippingsParams, fields ...string) (entity.Shippings, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(entity.Shippings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockShippingMockRecorder) List(ctx, params interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockShipping)(nil).List), varargs...)
}

// MultiGet mocks base method.
func (m *MockShipping) MultiGet(ctx context.Context, shippingIDs []string, fields ...string) (entity.Shippings, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, shippingIDs}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MultiGet", varargs...)
	ret0, _ := ret[0].(entity.Shippings)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGet indicates an expected call of MultiGet.
func (mr *MockShippingMockRecorder) MultiGet(ctx, shippingIDs interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, shippingIDs}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGet", reflect.TypeOf((*MockShipping)(nil).MultiGet), varargs...)
}

// Update mocks base method.
func (m *MockShipping) Update(ctx context.Context, shippingID string, params *database.UpdateShippingParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, shippingID, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockShippingMockRecorder) Update(ctx, shippingID, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockShipping)(nil).Update), ctx, shippingID, params)
}

// MockSchedule is a mock of Schedule interface.
type MockSchedule struct {
	ctrl     *gomock.Controller
	recorder *MockScheduleMockRecorder
}

// MockScheduleMockRecorder is the mock recorder for MockSchedule.
type MockScheduleMockRecorder struct {
	mock *MockSchedule
}

// NewMockSchedule creates a new mock instance.
func NewMockSchedule(ctrl *gomock.Controller) *MockSchedule {
	mock := &MockSchedule{ctrl: ctrl}
	mock.recorder = &MockScheduleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSchedule) EXPECT() *MockScheduleMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSchedule) Create(ctx context.Context, schedule *entity.Schedule, lives entity.Lives, products entity.LiveProducts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, schedule, lives, products)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockScheduleMockRecorder) Create(ctx, schedule, lives, products interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSchedule)(nil).Create), ctx, schedule, lives, products)
}

// MockLive is a mock of Live interface.
type MockLive struct {
	ctrl     *gomock.Controller
	recorder *MockLiveMockRecorder
}

// MockLiveMockRecorder is the mock recorder for MockLive.
type MockLiveMockRecorder struct {
	mock *MockLive
}

// NewMockLive creates a new mock instance.
func NewMockLive(ctrl *gomock.Controller) *MockLive {
	mock := &MockLive{ctrl: ctrl}
	mock.recorder = &MockLiveMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLive) EXPECT() *MockLiveMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockLive) Get(ctx context.Context, liveID string, fields ...string) (*entity.Live, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, liveID}
	for _, a := range fields {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*entity.Live)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockLiveMockRecorder) Get(ctx, liveID interface{}, fields ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, liveID}, fields...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLive)(nil).Get), varargs...)
}

// Update mocks base method.
func (m *MockLive) Update(ctx context.Context, liveID string, params *database.UpdateLiveParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, liveID, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockLiveMockRecorder) Update(ctx, liveID, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockLive)(nil).Update), ctx, liveID, params)
}

// UpdatePublic mocks base method.
func (m *MockLive) UpdatePublic(ctx context.Context, liveID string, params *database.UpdateLivePublicParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePublic", ctx, liveID, params)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePublic indicates an expected call of UpdatePublic.
func (mr *MockLiveMockRecorder) UpdatePublic(ctx, liveID, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePublic", reflect.TypeOf((*MockLive)(nil).UpdatePublic), ctx, liveID, params)
}
