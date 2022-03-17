// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_support is a generated GoMock package.
package mock_support

import (
	context "context"
	support "noname-realtime-support-chat/internal/support"
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

// CreateSupport mocks base method.
func (m *MockService) CreateSupport(ctx context.Context, email, name, password string) (*support.DTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSupport", ctx, email, name, password)
	ret0, _ := ret[0].(*support.DTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSupport indicates an expected call of CreateSupport.
func (mr *MockServiceMockRecorder) CreateSupport(ctx, email, name, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSupport", reflect.TypeOf((*MockService)(nil).CreateSupport), ctx, email, name, password)
}

// GetSupportByEmail mocks base method.
func (m *MockService) GetSupportByEmail(ctx context.Context, email string) (*support.DTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportByEmail", ctx, email)
	ret0, _ := ret[0].(*support.DTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportByEmail indicates an expected call of GetSupportByEmail.
func (mr *MockServiceMockRecorder) GetSupportByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportByEmail", reflect.TypeOf((*MockService)(nil).GetSupportByEmail), ctx, email)
}

// GetSupportById mocks base method.
func (m *MockService) GetSupportById(ctx context.Context, id string) (*support.DTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupportById", ctx, id)
	ret0, _ := ret[0].(*support.DTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupportById indicates an expected call of GetSupportById.
func (mr *MockServiceMockRecorder) GetSupportById(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupportById", reflect.TypeOf((*MockService)(nil).GetSupportById), ctx, id)
}
