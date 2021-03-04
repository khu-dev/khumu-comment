// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/comment.go

// Package usecase is a generated GoMock package.
package usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/khu-dev/khumu-comment/model"
	repository "github.com/khu-dev/khumu-comment/repository"
)

// MockCommentUseCaseInterface is a mock of CommentUseCaseInterface interface.
type MockCommentUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCommentUseCaseInterfaceMockRecorder
}

// MockCommentUseCaseInterfaceMockRecorder is the mock recorder for MockCommentUseCaseInterface.
type MockCommentUseCaseInterfaceMockRecorder struct {
	mock *MockCommentUseCaseInterface
}

// NewMockCommentUseCaseInterface creates a new mock instance.
func NewMockCommentUseCaseInterface(ctrl *gomock.Controller) *MockCommentUseCaseInterface {
	mock := &MockCommentUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockCommentUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommentUseCaseInterface) EXPECT() *MockCommentUseCaseInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCommentUseCaseInterface) Create(comment *model.Comment) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", comment)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCommentUseCaseInterfaceMockRecorder) Create(comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCommentUseCaseInterface)(nil).Create), comment)
}

// Delete mocks base method.
func (m *MockCommentUseCaseInterface) Delete(id int) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockCommentUseCaseInterfaceMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCommentUseCaseInterface)(nil).Delete), id)
}

// Get mocks base method.
func (m *MockCommentUseCaseInterface) Get(username string, id int) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", username, id)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCommentUseCaseInterfaceMockRecorder) Get(username, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCommentUseCaseInterface)(nil).Get), username, id)
}

// List mocks base method.
func (m *MockCommentUseCaseInterface) List(username string, opt *repository.CommentQueryOption) ([]*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", username, opt)
	ret0, _ := ret[0].([]*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockCommentUseCaseInterfaceMockRecorder) List(username, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockCommentUseCaseInterface)(nil).List), username, opt)
}

// Update mocks base method.
func (m *MockCommentUseCaseInterface) Update(username string, id int, opt map[string]interface{}) (*model.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", username, id, opt)
	ret0, _ := ret[0].(*model.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCommentUseCaseInterfaceMockRecorder) Update(username, id, opt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCommentUseCaseInterface)(nil).Update), username, id, opt)
}

// MockLikeCommentUseCaseInterface is a mock of LikeCommentUseCaseInterface interface.
type MockLikeCommentUseCaseInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLikeCommentUseCaseInterfaceMockRecorder
}

// MockLikeCommentUseCaseInterfaceMockRecorder is the mock recorder for MockLikeCommentUseCaseInterface.
type MockLikeCommentUseCaseInterfaceMockRecorder struct {
	mock *MockLikeCommentUseCaseInterface
}

// NewMockLikeCommentUseCaseInterface creates a new mock instance.
func NewMockLikeCommentUseCaseInterface(ctrl *gomock.Controller) *MockLikeCommentUseCaseInterface {
	mock := &MockLikeCommentUseCaseInterface{ctrl: ctrl}
	mock.recorder = &MockLikeCommentUseCaseInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLikeCommentUseCaseInterface) EXPECT() *MockLikeCommentUseCaseInterfaceMockRecorder {
	return m.recorder
}

// Toggle mocks base method.
func (m *MockLikeCommentUseCaseInterface) Toggle(like *model.LikeComment) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Toggle", like)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Toggle indicates an expected call of Toggle.
func (mr *MockLikeCommentUseCaseInterfaceMockRecorder) Toggle(like interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Toggle", reflect.TypeOf((*MockLikeCommentUseCaseInterface)(nil).Toggle), like)
}