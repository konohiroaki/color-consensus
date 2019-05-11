// Code generated by MockGen. DO NOT EDIT.
// Source: backend/repositories/nationality.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNationalityRepository is a mock of NationalityRepository interface
type MockNationalityRepository struct {
	ctrl     *gomock.Controller
	recorder *MockNationalityRepositoryMockRecorder
}

// MockNationalityRepositoryMockRecorder is the mock recorder for MockNationalityRepository
type MockNationalityRepositoryMockRecorder struct {
	mock *MockNationalityRepository
}

// NewMockNationalityRepository creates a new mock instance
func NewMockNationalityRepository(ctrl *gomock.Controller) *MockNationalityRepository {
	mock := &MockNationalityRepository{ctrl: ctrl}
	mock.recorder = &MockNationalityRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNationalityRepository) EXPECT() *MockNationalityRepositoryMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockNationalityRepository) GetAll() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockNationalityRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNationalityRepository)(nil).GetAll))
}

// IsCodePresent mocks base method
func (m *MockNationalityRepository) IsCodePresent(code string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsCodePresent", code)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsCodePresent indicates an expected call of IsCodePresent
func (mr *MockNationalityRepositoryMockRecorder) IsCodePresent(code interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsCodePresent", reflect.TypeOf((*MockNationalityRepository)(nil).IsCodePresent), code)
}