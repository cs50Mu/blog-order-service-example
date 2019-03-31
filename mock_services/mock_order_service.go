// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/SebastianCoetzee/blog-order-service-example/services (interfaces: OrderService)

// Package mock_services is a generated GoMock package.
package mock_services

import (
	models "github.com/SebastianCoetzee/blog-order-service-example/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderService is a mock of OrderService interface
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// FindAllOrdersByUserID mocks base method
func (m *MockOrderService) FindAllOrdersByUserID(arg0 int) (models.Orders, error) {
	ret := m.ctrl.Call(m, "FindAllOrdersByUserID", arg0)
	ret0, _ := ret[0].(models.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllOrdersByUserID indicates an expected call of FindAllOrdersByUserID
func (mr *MockOrderServiceMockRecorder) FindAllOrdersByUserID(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllOrdersByUserID", reflect.TypeOf((*MockOrderService)(nil).FindAllOrdersByUserID), arg0)
}
