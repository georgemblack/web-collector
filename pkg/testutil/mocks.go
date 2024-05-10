// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/api.go
//
// Generated by this command:
//
//	mockgen -source=pkg/api/api.go -destination=pkg/testutil/mocks.go -package=testutil
//

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	repo "github.com/georgemblack/web-api/pkg/repo"
	types "github.com/georgemblack/web-api/pkg/types"
	gomock "go.uber.org/mock/gomock"
)

// MockFirestoreService is a mock of FirestoreService interface.
type MockFirestoreService struct {
	ctrl     *gomock.Controller
	recorder *MockFirestoreServiceMockRecorder
}

// MockFirestoreServiceMockRecorder is the mock recorder for MockFirestoreService.
type MockFirestoreServiceMockRecorder struct {
	mock *MockFirestoreService
}

// NewMockFirestoreService creates a new mock instance.
func NewMockFirestoreService(ctrl *gomock.Controller) *MockFirestoreService {
	mock := &MockFirestoreService{ctrl: ctrl}
	mock.recorder = &MockFirestoreServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFirestoreService) EXPECT() *MockFirestoreServiceMockRecorder {
	return m.recorder
}

// AddLike mocks base method.
func (m *MockFirestoreService) AddLike(like types.Like) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLike", like)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddLike indicates an expected call of AddLike.
func (mr *MockFirestoreServiceMockRecorder) AddLike(like any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLike", reflect.TypeOf((*MockFirestoreService)(nil).AddLike), like)
}

// AddPost mocks base method.
func (m *MockFirestoreService) AddPost(post types.Post) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPost", post)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPost indicates an expected call of AddPost.
func (mr *MockFirestoreServiceMockRecorder) AddPost(post any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPost", reflect.TypeOf((*MockFirestoreService)(nil).AddPost), post)
}

// Close mocks base method.
func (m *MockFirestoreService) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockFirestoreServiceMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFirestoreService)(nil).Close))
}

// DeleteLike mocks base method.
func (m *MockFirestoreService) DeleteLike(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLike", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLike indicates an expected call of DeleteLike.
func (mr *MockFirestoreServiceMockRecorder) DeleteLike(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLike", reflect.TypeOf((*MockFirestoreService)(nil).DeleteLike), id)
}

// DeletePost mocks base method.
func (m *MockFirestoreService) DeletePost(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockFirestoreServiceMockRecorder) DeletePost(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockFirestoreService)(nil).DeletePost), id)
}

// GetHashList mocks base method.
func (m *MockFirestoreService) GetHashList() (types.HashList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHashList")
	ret0, _ := ret[0].(types.HashList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHashList indicates an expected call of GetHashList.
func (mr *MockFirestoreServiceMockRecorder) GetHashList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHashList", reflect.TypeOf((*MockFirestoreService)(nil).GetHashList))
}

// GetLike mocks base method.
func (m *MockFirestoreService) GetLike(id string) (types.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLike", id)
	ret0, _ := ret[0].(types.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLike indicates an expected call of GetLike.
func (mr *MockFirestoreServiceMockRecorder) GetLike(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLike", reflect.TypeOf((*MockFirestoreService)(nil).GetLike), id)
}

// GetLikes mocks base method.
func (m *MockFirestoreService) GetLikes() ([]types.Like, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLikes")
	ret0, _ := ret[0].([]types.Like)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLikes indicates an expected call of GetLikes.
func (mr *MockFirestoreServiceMockRecorder) GetLikes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLikes", reflect.TypeOf((*MockFirestoreService)(nil).GetLikes))
}

// GetPost mocks base method.
func (m *MockFirestoreService) GetPost(id string) (types.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPost", id)
	ret0, _ := ret[0].(types.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPost indicates an expected call of GetPost.
func (mr *MockFirestoreServiceMockRecorder) GetPost(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPost", reflect.TypeOf((*MockFirestoreService)(nil).GetPost), id)
}

// GetPosts mocks base method.
func (m *MockFirestoreService) GetPosts(filters repo.PostFilters) ([]types.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPosts", filters)
	ret0, _ := ret[0].([]types.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPosts indicates an expected call of GetPosts.
func (mr *MockFirestoreServiceMockRecorder) GetPosts(filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPosts", reflect.TypeOf((*MockFirestoreService)(nil).GetPosts), filters)
}

// UpdateHashList mocks base method.
func (m *MockFirestoreService) UpdateHashList(hashList types.HashList) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateHashList", hashList)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateHashList indicates an expected call of UpdateHashList.
func (mr *MockFirestoreServiceMockRecorder) UpdateHashList(hashList any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateHashList", reflect.TypeOf((*MockFirestoreService)(nil).UpdateHashList), hashList)
}

// UpdatePost mocks base method.
func (m *MockFirestoreService) UpdatePost(post types.Post) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", post)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockFirestoreServiceMockRecorder) UpdatePost(post any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockFirestoreService)(nil).UpdatePost), post)
}
