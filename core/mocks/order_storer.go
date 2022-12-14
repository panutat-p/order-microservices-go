// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/panutat-p/order-microservices-go/core/port (interfaces: OrderStorer)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/panutat-p/order-microservices-go/core/domain"
)

// MockOrderStorer is a mock of OrderStorer interface.
type MockOrderStorer struct {
	ctrl     *gomock.Controller
	recorder *MockOrderStorerMockRecorder
}

// MockOrderStorerMockRecorder is the mock recorder for MockOrderStorer.
type MockOrderStorerMockRecorder struct {
	mock *MockOrderStorer
}

// NewMockOrderStorer creates a new mock instance.
func NewMockOrderStorer(ctrl *gomock.Controller) *MockOrderStorer {
	mock := &MockOrderStorer{ctrl: ctrl}
	mock.recorder = &MockOrderStorerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderStorer) EXPECT() *MockOrderStorerMockRecorder {
	return m.recorder
}

// Find mocks base method.
func (m *MockOrderStorer) Find(arg0 string) ([]domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", arg0)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockOrderStorerMockRecorder) Find(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockOrderStorer)(nil).Find), arg0)
}

// Save mocks base method.
func (m *MockOrderStorer) Save(arg0 *domain.Order) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOrderStorerMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOrderStorer)(nil).Save), arg0)
}
