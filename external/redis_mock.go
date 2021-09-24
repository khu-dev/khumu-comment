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

// FindAllByCommentID mocks base method.
func (m *MockRedisAdapter) FindAllByCommentID(commentID int) (data.LikeCommentEntities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllByCommentID", commentID)
	ret0, _ := ret[0].(data.LikeCommentEntities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllByCommentID indicates an expected call of FindAllByCommentID.
func (mr *MockRedisAdapterMockRecorder) FindAllByCommentID(commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllByCommentID", reflect.TypeOf((*MockRedisAdapter)(nil).FindAllByCommentID), commentID)
}

// FindAllParentCommentsByArticleID mocks base method.
func (m *MockRedisAdapter) FindAllParentCommentsByArticleID(articleID int) (data.CommentEntities, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllParentCommentsByArticleID", articleID)
	ret0, _ := ret[0].(data.CommentEntities)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllParentCommentsByArticleID indicates an expected call of FindAllParentCommentsByArticleID.
func (mr *MockRedisAdapterMockRecorder) FindAllParentCommentsByArticleID(articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllParentCommentsByArticleID", reflect.TypeOf((*MockRedisAdapter)(nil).FindAllParentCommentsByArticleID), articleID)
}

// SetCommentsByArticleID mocks base method.
func (m *MockRedisAdapter) SetCommentsByArticleID(articleID int, coms data.CommentEntities) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetCommentsByArticleID", articleID, coms)
}

// SetCommentsByArticleID indicates an expected call of SetCommentsByArticleID.
func (mr *MockRedisAdapterMockRecorder) SetCommentsByArticleID(articleID, coms interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCommentsByArticleID", reflect.TypeOf((*MockRedisAdapter)(nil).SetCommentsByArticleID), articleID, coms)
}

// SetLikesByCommentID mocks base method.
func (m *MockRedisAdapter) SetLikesByCommentID(commentID int, likes data.LikeCommentEntities) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLikesByCommentID", commentID, likes)
}

// SetLikesByCommentID indicates an expected call of SetLikesByCommentID.
func (mr *MockRedisAdapterMockRecorder) SetLikesByCommentID(commentID, likes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLikesByCommentID", reflect.TypeOf((*MockRedisAdapter)(nil).SetLikesByCommentID), commentID, likes)
}
