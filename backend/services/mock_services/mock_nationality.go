// Code generated by MockGen. DO NOT EDIT.
// Source: backend/services/nationality.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNationalityService is a mock of NationalityService interface
type MockNationalityService struct {
	ctrl     *gomock.Controller
	recorder *MockNationalityServiceMockRecorder
}

// MockNationalityServiceMockRecorder is the mock recorder for MockNationalityService
type MockNationalityServiceMockRecorder struct {
	mock *MockNationalityService
}

// NewMockNationalityService creates a new mock instance
func NewMockNationalityService(ctrl *gomock.Controller) *MockNationalityService {
	mock := &MockNationalityService{ctrl: ctrl}
	mock.recorder = &MockNationalityServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNationalityService) EXPECT() *MockNationalityServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockNationalityService) GetAll() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockNationalityServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNationalityService)(nil).GetAll))
}
