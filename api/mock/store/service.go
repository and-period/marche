// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	context "context"
	reflect "reflect"

	store "github.com/and-period/furumaru/api/internal/store"
	entity "github.com/and-period/furumaru/api/internal/store/entity"
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

// CreateCategory mocks base method.
func (m *MockService) CreateCategory(ctx context.Context, in *store.CreateCategoryInput) (*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", ctx, in)
	ret0, _ := ret[0].(*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockServiceMockRecorder) CreateCategory(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockService)(nil).CreateCategory), ctx, in)
}

// CreateProduct mocks base method.
func (m *MockService) CreateProduct(ctx context.Context, in *store.CreateProductInput) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", ctx, in)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockServiceMockRecorder) CreateProduct(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockService)(nil).CreateProduct), ctx, in)
}

// CreateProductType mocks base method.
func (m *MockService) CreateProductType(ctx context.Context, in *store.CreateProductTypeInput) (*entity.ProductType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProductType", ctx, in)
	ret0, _ := ret[0].(*entity.ProductType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProductType indicates an expected call of CreateProductType.
func (mr *MockServiceMockRecorder) CreateProductType(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProductType", reflect.TypeOf((*MockService)(nil).CreateProductType), ctx, in)
}

// CreatePromotion mocks base method.
func (m *MockService) CreatePromotion(ctx context.Context, in *store.CreatePromotionInput) (*entity.Promotion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePromotion", ctx, in)
	ret0, _ := ret[0].(*entity.Promotion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePromotion indicates an expected call of CreatePromotion.
func (mr *MockServiceMockRecorder) CreatePromotion(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePromotion", reflect.TypeOf((*MockService)(nil).CreatePromotion), ctx, in)
}

// CreateSchedule mocks base method.
func (m *MockService) CreateSchedule(ctx context.Context, in *store.CreateScheduleInput) (*entity.Schedule, entity.Lives, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSchedule", ctx, in)
	ret0, _ := ret[0].(*entity.Schedule)
	ret1, _ := ret[1].(entity.Lives)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateSchedule indicates an expected call of CreateSchedule.
func (mr *MockServiceMockRecorder) CreateSchedule(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSchedule", reflect.TypeOf((*MockService)(nil).CreateSchedule), ctx, in)
}

// CreateShipping mocks base method.
func (m *MockService) CreateShipping(ctx context.Context, in *store.CreateShippingInput) (*entity.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShipping", ctx, in)
	ret0, _ := ret[0].(*entity.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateShipping indicates an expected call of CreateShipping.
func (mr *MockServiceMockRecorder) CreateShipping(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShipping", reflect.TypeOf((*MockService)(nil).CreateShipping), ctx, in)
}

// DeleteCategory mocks base method.
func (m *MockService) DeleteCategory(ctx context.Context, in *store.DeleteCategoryInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockServiceMockRecorder) DeleteCategory(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockService)(nil).DeleteCategory), ctx, in)
}

// DeleteProduct mocks base method.
func (m *MockService) DeleteProduct(ctx context.Context, in *store.DeleteProductInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockServiceMockRecorder) DeleteProduct(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockService)(nil).DeleteProduct), ctx, in)
}

// DeleteProductType mocks base method.
func (m *MockService) DeleteProductType(ctx context.Context, in *store.DeleteProductTypeInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProductType", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProductType indicates an expected call of DeleteProductType.
func (mr *MockServiceMockRecorder) DeleteProductType(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProductType", reflect.TypeOf((*MockService)(nil).DeleteProductType), ctx, in)
}

// DeletePromotion mocks base method.
func (m *MockService) DeletePromotion(ctx context.Context, in *store.DeletePromotionInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePromotion", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePromotion indicates an expected call of DeletePromotion.
func (mr *MockServiceMockRecorder) DeletePromotion(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePromotion", reflect.TypeOf((*MockService)(nil).DeletePromotion), ctx, in)
}

// DeleteShipping mocks base method.
func (m *MockService) DeleteShipping(ctx context.Context, in *store.DeleteShippingInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteShipping", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteShipping indicates an expected call of DeleteShipping.
func (mr *MockServiceMockRecorder) DeleteShipping(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteShipping", reflect.TypeOf((*MockService)(nil).DeleteShipping), ctx, in)
}

// GetCategory mocks base method.
func (m *MockService) GetCategory(ctx context.Context, in *store.GetCategoryInput) (*entity.Category, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategory", ctx, in)
	ret0, _ := ret[0].(*entity.Category)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCategory indicates an expected call of GetCategory.
func (mr *MockServiceMockRecorder) GetCategory(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategory", reflect.TypeOf((*MockService)(nil).GetCategory), ctx, in)
}

// GetProduct mocks base method.
func (m *MockService) GetProduct(ctx context.Context, in *store.GetProductInput) (*entity.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", ctx, in)
	ret0, _ := ret[0].(*entity.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockServiceMockRecorder) GetProduct(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockService)(nil).GetProduct), ctx, in)
}

// GetProductType mocks base method.
func (m *MockService) GetProductType(ctx context.Context, in *store.GetProductTypeInput) (*entity.ProductType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductType", ctx, in)
	ret0, _ := ret[0].(*entity.ProductType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductType indicates an expected call of GetProductType.
func (mr *MockServiceMockRecorder) GetProductType(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductType", reflect.TypeOf((*MockService)(nil).GetProductType), ctx, in)
}

// GetPromotion mocks base method.
func (m *MockService) GetPromotion(ctx context.Context, in *store.GetPromotionInput) (*entity.Promotion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPromotion", ctx, in)
	ret0, _ := ret[0].(*entity.Promotion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPromotion indicates an expected call of GetPromotion.
func (mr *MockServiceMockRecorder) GetPromotion(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPromotion", reflect.TypeOf((*MockService)(nil).GetPromotion), ctx, in)
}

// GetShipping mocks base method.
func (m *MockService) GetShipping(ctx context.Context, in *store.GetShippingInput) (*entity.Shipping, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShipping", ctx, in)
	ret0, _ := ret[0].(*entity.Shipping)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShipping indicates an expected call of GetShipping.
func (mr *MockServiceMockRecorder) GetShipping(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShipping", reflect.TypeOf((*MockService)(nil).GetShipping), ctx, in)
}

// ListCategories mocks base method.
func (m *MockService) ListCategories(ctx context.Context, in *store.ListCategoriesInput) (entity.Categories, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCategories", ctx, in)
	ret0, _ := ret[0].(entity.Categories)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListCategories indicates an expected call of ListCategories.
func (mr *MockServiceMockRecorder) ListCategories(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCategories", reflect.TypeOf((*MockService)(nil).ListCategories), ctx, in)
}

// ListProductTypes mocks base method.
func (m *MockService) ListProductTypes(ctx context.Context, in *store.ListProductTypesInput) (entity.ProductTypes, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProductTypes", ctx, in)
	ret0, _ := ret[0].(entity.ProductTypes)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProductTypes indicates an expected call of ListProductTypes.
func (mr *MockServiceMockRecorder) ListProductTypes(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProductTypes", reflect.TypeOf((*MockService)(nil).ListProductTypes), ctx, in)
}

// ListProducts mocks base method.
func (m *MockService) ListProducts(ctx context.Context, in *store.ListProductsInput) (entity.Products, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", ctx, in)
	ret0, _ := ret[0].(entity.Products)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockServiceMockRecorder) ListProducts(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockService)(nil).ListProducts), ctx, in)
}

// ListPromotions mocks base method.
func (m *MockService) ListPromotions(ctx context.Context, in *store.ListPromotionsInput) (entity.Promotions, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPromotions", ctx, in)
	ret0, _ := ret[0].(entity.Promotions)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListPromotions indicates an expected call of ListPromotions.
func (mr *MockServiceMockRecorder) ListPromotions(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPromotions", reflect.TypeOf((*MockService)(nil).ListPromotions), ctx, in)
}

// ListShippings mocks base method.
func (m *MockService) ListShippings(ctx context.Context, in *store.ListShippingsInput) (entity.Shippings, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListShippings", ctx, in)
	ret0, _ := ret[0].(entity.Shippings)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListShippings indicates an expected call of ListShippings.
func (mr *MockServiceMockRecorder) ListShippings(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListShippings", reflect.TypeOf((*MockService)(nil).ListShippings), ctx, in)
}

// MultiGetCategories mocks base method.
func (m *MockService) MultiGetCategories(ctx context.Context, in *store.MultiGetCategoriesInput) (entity.Categories, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetCategories", ctx, in)
	ret0, _ := ret[0].(entity.Categories)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetCategories indicates an expected call of MultiGetCategories.
func (mr *MockServiceMockRecorder) MultiGetCategories(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetCategories", reflect.TypeOf((*MockService)(nil).MultiGetCategories), ctx, in)
}

// MultiGetProductTypes mocks base method.
func (m *MockService) MultiGetProductTypes(ctx context.Context, in *store.MultiGetProductTypesInput) (entity.ProductTypes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiGetProductTypes", ctx, in)
	ret0, _ := ret[0].(entity.ProductTypes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MultiGetProductTypes indicates an expected call of MultiGetProductTypes.
func (mr *MockServiceMockRecorder) MultiGetProductTypes(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiGetProductTypes", reflect.TypeOf((*MockService)(nil).MultiGetProductTypes), ctx, in)
}

// UpdateCategory mocks base method.
func (m *MockService) UpdateCategory(ctx context.Context, in *store.UpdateCategoryInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategory", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCategory indicates an expected call of UpdateCategory.
func (mr *MockServiceMockRecorder) UpdateCategory(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategory", reflect.TypeOf((*MockService)(nil).UpdateCategory), ctx, in)
}

// UpdateProduct mocks base method.
func (m *MockService) UpdateProduct(ctx context.Context, in *store.UpdateProductInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockServiceMockRecorder) UpdateProduct(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockService)(nil).UpdateProduct), ctx, in)
}

// UpdateProductType mocks base method.
func (m *MockService) UpdateProductType(ctx context.Context, in *store.UpdateProductTypeInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductType", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductType indicates an expected call of UpdateProductType.
func (mr *MockServiceMockRecorder) UpdateProductType(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductType", reflect.TypeOf((*MockService)(nil).UpdateProductType), ctx, in)
}

// UpdatePromotion mocks base method.
func (m *MockService) UpdatePromotion(ctx context.Context, in *store.UpdatePromotionInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePromotion", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePromotion indicates an expected call of UpdatePromotion.
func (mr *MockServiceMockRecorder) UpdatePromotion(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePromotion", reflect.TypeOf((*MockService)(nil).UpdatePromotion), ctx, in)
}

// UpdateShipping mocks base method.
func (m *MockService) UpdateShipping(ctx context.Context, in *store.UpdateShippingInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateShipping", ctx, in)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateShipping indicates an expected call of UpdateShipping.
func (mr *MockServiceMockRecorder) UpdateShipping(ctx, in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateShipping", reflect.TypeOf((*MockService)(nil).UpdateShipping), ctx, in)
}
