// Code generated by MockGen. DO NOT EDIT.
// Source: app/domain/repository/db/transaction.go
//
// Generated by this command:
//
//	mockgen -source=app/domain/repository/db/transaction.go -destination=app/domain/repository/db/mock/transaction.go
//

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// WithInTx mocks base method.
func (m *MockTransactionRepository) WithInTx(ctx context.Context, f func(context.Context) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithInTx", ctx, f)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithInTx indicates an expected call of WithInTx.
func (mr *MockTransactionRepositoryMockRecorder) WithInTx(ctx, f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithInTx", reflect.TypeOf((*MockTransactionRepository)(nil).WithInTx), ctx, f)
}
