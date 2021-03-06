// Code generated by MockGen. DO NOT EDIT.
// Source: backend/services/gender.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGenderService is a mock of GenderService interface
type MockGenderService struct {
	ctrl     *gomock.Controller
	recorder *MockGenderServiceMockRecorder
}

// MockGenderServiceMockRecorder is the mock recorder for MockGenderService
type MockGenderServiceMockRecorder struct {
	mock *MockGenderService
}

// NewMockGenderService creates a new mock instance
func NewMockGenderService(ctrl *gomock.Controller) *MockGenderService {
	mock := &MockGenderService{ctrl: ctrl}
	mock.recorder = &MockGenderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGenderService) EXPECT() *MockGenderServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockGenderService) GetAll() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockGenderServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockGenderService)(nil).GetAll))
}
