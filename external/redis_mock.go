// Code generated by MockGen. DO NOT EDIT.
// Source: external/redis.go

// Package external is a generated GoMock package.
package external

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	data "github.com/khu-dev/khumu-comment/data"
)

// MockRedisAdapter is a mock of RedisAdapter interface.
type MockRedisAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockRedisAdapterMockRecorder
}

// MockRedisAdapterMockRecorder is the mock recorder for MockRedisAdapter.
type MockRedisAdapterMockRecorder struct {
	mock *MockRedisAdapter
}

// NewMockRedisAdapter creates a new mock instance.
func NewMockRedisAdapter(ctrl *gomock.Controller) *MockRedisAdapter {
	mock := &MockRedisAdapter{ctrl: ctrl}
	mock.recorder = &MockRedisAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisAdapter) EXPECT() *MockRedisAdapterMockRecorder {
	return m.recorder
}

// GetAllByArticle mocks base method.
func (m *MockRedisAdapter) GetAllByArticle(articleID int) data.CommentEntities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByArticle", articleID)
	ret0, _ := ret[0].(data.CommentEntities)
	return ret0
}

// GetAllByArticle indicates an expected call of GetAllByArticle.
func (mr *MockRedisAdapterMockRecorder) GetAllByArticle(articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByArticle", reflect.TypeOf((*MockRedisAdapter)(nil).GetAllByArticle), articleID)
}

// InvalidateCommentsOfArticle mocks base method.
func (m *MockRedisAdapter) InvalidateCommentsOfArticle(articleID int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "InvalidateCommentsOfArticle", articleID)
}

// InvalidateCommentsOfArticle indicates an expected call of InvalidateCommentsOfArticle.
func (mr *MockRedisAdapterMockRecorder) InvalidateCommentsOfArticle(articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InvalidateCommentsOfArticle", reflect.TypeOf((*MockRedisAdapter)(nil).InvalidateCommentsOfArticle), articleID)
}

// Refresh mocks base method.
func (m *MockRedisAdapter) Refresh(articleID int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Refresh", articleID)
}

// Refresh indicates an expected call of Refresh.
func (mr *MockRedisAdapterMockRecorder) Refresh(articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockRedisAdapter)(nil).Refresh), articleID)
}
