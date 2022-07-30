// Code generated by MockGen. DO NOT EDIT.
// Source: eshop-mock-api/internal/services/product_service (interfaces: Service)

// Package mock_product_service is a generated GoMock package.
package mock_product_service

import (
	models "eshop-mock-api/internal/models"
	product_service "eshop-mock-api/internal/services/product_service"
	struct_templates "eshop-mock-api/internal/struct_templates"
	reflect "reflect"

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

// Create mocks base method.
func (m *MockService) Create(arg0 models.Product) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockServiceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockService)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockService) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockServiceMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockService)(nil).Delete), arg0)
}

// Detail mocks base method.
func (m *MockService) Detail(arg0 product_service.SearchDetailData, arg1 ...product_service.Preload) (product_service.Product, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Detail", varargs...)
	ret0, _ := ret[0].(product_service.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Detail indicates an expected call of Detail.
func (mr *MockServiceMockRecorder) Detail(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Detail", reflect.TypeOf((*MockService)(nil).Detail), varargs...)
}

// GetFirstOrCreateUpdate mocks base method.
func (m *MockService) GetFirstOrCreateUpdate(arg0 product_service.WhereGetFirstOrCreateData, arg1 models.Product) (models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFirstOrCreateUpdate", arg0, arg1)
	ret0, _ := ret[0].(models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFirstOrCreateUpdate indicates an expected call of GetFirstOrCreateUpdate.
func (mr *MockServiceMockRecorder) GetFirstOrCreateUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFirstOrCreateUpdate", reflect.TypeOf((*MockService)(nil).GetFirstOrCreateUpdate), arg0, arg1)
}

// Paginate mocks base method.
func (m *MockService) Paginate(arg0 product_service.SearchPageData, arg1 ...product_service.Preload) ([]product_service.Product, struct_templates.Pagination, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Paginate", varargs...)
	ret0, _ := ret[0].([]product_service.Product)
	ret1, _ := ret[1].(struct_templates.Pagination)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Paginate indicates an expected call of Paginate.
func (mr *MockServiceMockRecorder) Paginate(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockService)(nil).Paginate), varargs...)
}

// Update mocks base method.
func (m *MockService) Update(arg0 []string, arg1 product_service.WhereUpdateData, arg2 models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockServiceMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockService)(nil).Update), arg0, arg1, arg2)
}
