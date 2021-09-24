// Code generated by MockGen. DO NOT EDIT.
// Source: repository/comment.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	data "github.com/khu-dev/khumu-comment/data"
	ent "github.com/khu-dev/khumu-comment/ent"
)

// MockCommentRepository is a mock of CommentRepository interface.
type MockCommentRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCommentRepositoryMockRecorder
}

// MockCommentRepositoryMockRecorder is the mock recorder for MockCommentRepository.
type MockCommentRepositoryMockRecorder struct {
	mock *MockCommentRepository
}

// NewMockCommentRepository creates a new mock instance.
func NewMockCommentRepository(ctrl *gomock.Controller) *MockCommentRepository {
	mock := &MockCommentRepository{ctrl: ctrl}
	mock.recorder = &MockCommentRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentRepository) EXPECT() *MockCommentRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCommentRepository) Create(createInput *data.CommentInput, isWrittenByArticleAuthor bool) (*ent.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", createInput, isWrittenByArticleAuthor)
	ret0, _ := ret[0].(*ent.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCommentRepositoryMockRecorder) Create(createInput, isWrittenByArticleAuthor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommentRepository)(nil).Create), createInput, isWrittenByArticleAuthor)
}

// Delete mocks base method.
func (m *MockCommentRepository) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentRepositoryMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommentRepository)(nil).Delete), id)
}

// FindAllParentCommentsByArticleID mocks base method.
func (m *MockCommentRepository) FindAllParentCommentsByArticleID(articleID int) ([]*ent.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllParentCommentsByArticleID", articleID)
	ret0, _ := ret[0].([]*ent.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllParentCommentsByArticleID indicates an expected call of FindAllParentCommentsByArticleID.
func (mr *MockCommentRepositoryMockRecorder) FindAllParentCommentsByArticleID(articleID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllParentCommentsByArticleID", reflect.TypeOf((*MockCommentRepository)(nil).FindAllParentCommentsByArticleID), articleID)
}

// FindAllParentCommentsByAuthorID mocks base method.
func (m *MockCommentRepository) FindAllParentCommentsByAuthorID(authorID string) ([]*ent.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllParentCommentsByAuthorID", authorID)
	ret0, _ := ret[0].([]*ent.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllParentCommentsByAuthorID indicates an expected call of FindAllParentCommentsByAuthorID.
func (mr *MockCommentRepositoryMockRecorder) FindAllParentCommentsByAuthorID(authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllParentCommentsByAuthorID", reflect.TypeOf((*MockCommentRepository)(nil).FindAllParentCommentsByAuthorID), authorID)
}

// Get mocks base method.
func (m *MockCommentRepository) Get(id int) (*ent.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*ent.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCommentRepositoryMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCommentRepository)(nil).Get), id)
}

// Update mocks base method.
func (m *MockCommentRepository) Update(id int, updateInput map[string]interface{}) (*ent.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, updateInput)
	ret0, _ := ret[0].(*ent.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCommentRepositoryMockRecorder) Update(id, updateInput interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommentRepository)(nil).Update), id, updateInput)
}
