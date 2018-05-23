// Code generated by MockGen. DO NOT EDIT.
// Source: generator.go

// Package mock is a generated GoMock package.
package mock

import (
	generator "github.com/deadcheat/awsset/generator"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// LoadFiles mocks base method
func (m *MockUseCase) LoadFiles(paths, ignorePatterns []string) (*generator.Entity, error) {
	ret := m.ctrl.Call(m, "LoadFiles", paths, ignorePatterns)
	ret0, _ := ret[0].(*generator.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadFiles indicates an expected call of LoadFiles
func (mr *MockUseCaseMockRecorder) LoadFiles(paths, ignorePatterns interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadFiles", reflect.TypeOf((*MockUseCase)(nil).LoadFiles), paths, ignorePatterns)
}

// MockRegexpRepository is a mock of RegexpRepository interface
type MockRegexpRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRegexpRepositoryMockRecorder
}

// MockRegexpRepositoryMockRecorder is the mock recorder for MockRegexpRepository
type MockRegexpRepositoryMockRecorder struct {
	mock *MockRegexpRepository
}

// NewMockRegexpRepository creates a new mock instance
func NewMockRegexpRepository(ctrl *gomock.Controller) *MockRegexpRepository {
	mock := &MockRegexpRepository{ctrl: ctrl}
	mock.recorder = &MockRegexpRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRegexpRepository) EXPECT() *MockRegexpRepositoryMockRecorder {
	return m.recorder
}

// CompilePatterns mocks base method
func (m *MockRegexpRepository) CompilePatterns(patterns []string) error {
	ret := m.ctrl.Call(m, "CompilePatterns", patterns)
	ret0, _ := ret[0].(error)
	return ret0
}

// CompilePatterns indicates an expected call of CompilePatterns
func (mr *MockRegexpRepositoryMockRecorder) CompilePatterns(patterns interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompilePatterns", reflect.TypeOf((*MockRegexpRepository)(nil).CompilePatterns), patterns)
}

// MatchAny mocks base method
func (m *MockRegexpRepository) MatchAny(path string) bool {
	ret := m.ctrl.Call(m, "MatchAny", path)
	ret0, _ := ret[0].(bool)
	return ret0
}

// MatchAny indicates an expected call of MatchAny
func (mr *MockRegexpRepositoryMockRecorder) MatchAny(path interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MatchAny", reflect.TypeOf((*MockRegexpRepository)(nil).MatchAny), path)
}
