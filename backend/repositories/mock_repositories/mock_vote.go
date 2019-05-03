// Code generated by MockGen. DO NOT EDIT.
// Source: backend/repositories/vote.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockVoteRepository is a mock of VoteRepository interface
type MockVoteRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVoteRepositoryMockRecorder
}

// MockVoteRepositoryMockRecorder is the mock recorder for MockVoteRepository
type MockVoteRepositoryMockRecorder struct {
	mock *MockVoteRepository
}

// NewMockVoteRepository creates a new mock instance
func NewMockVoteRepository(ctrl *gomock.Controller) *MockVoteRepository {
	mock := &MockVoteRepository{ctrl: ctrl}
	mock.recorder = &MockVoteRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockVoteRepository) EXPECT() *MockVoteRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockVoteRepository) Add(user, lang, name string, newColors []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", user, lang, name, newColors)
}

// Add indicates an expected call of Add
func (mr *MockVoteRepositoryMockRecorder) Add(user, lang, name, newColors interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockVoteRepository)(nil).Add), user, lang, name, newColors)
}

// Get mocks base method
func (m *MockVoteRepository) Get(lang, name string, fields []string) []map[string]interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", lang, name, fields)
	ret0, _ := ret[0].([]map[string]interface{})
	return ret0
}

// Get indicates an expected call of Get
func (mr *MockVoteRepositoryMockRecorder) Get(lang, name, fields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVoteRepository)(nil).Get), lang, name, fields)
}

// RemoveByUser mocks base method
func (m *MockVoteRepository) RemoveByUser(userID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RemoveByUser", userID)
}

// RemoveByUser indicates an expected call of RemoveByUser
func (mr *MockVoteRepositoryMockRecorder) RemoveByUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveByUser", reflect.TypeOf((*MockVoteRepository)(nil).RemoveByUser), userID)
}