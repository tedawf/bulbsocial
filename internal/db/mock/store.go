// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/tedawf/bulbsocial/internal/db (interfaces: Store,Querier)
//
// Generated by this command:
//
//	mockgen -package=mockdb -destination internal/db/mock/store.go github.com/tedawf/bulbsocial/internal/db Store,Querier
//

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	db "github.com/tedawf/bulbsocial/internal/db"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
	isgomock struct{}
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockStore) CreateComment(ctx context.Context, arg db.CreateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, arg)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockStoreMockRecorder) CreateComment(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockStore)(nil).CreateComment), ctx, arg)
}

// CreatePost mocks base method.
func (m *MockStore) CreatePost(ctx context.Context, arg db.CreatePostParams) (db.CreatePostRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, arg)
	ret0, _ := ret[0].(db.CreatePostRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockStoreMockRecorder) CreatePost(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockStore)(nil).CreatePost), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), ctx, arg)
}

// DeleteComment mocks base method.
func (m *MockStore) DeleteComment(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockStoreMockRecorder) DeleteComment(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockStore)(nil).DeleteComment), ctx, id)
}

// DeletePost mocks base method.
func (m *MockStore) DeletePost(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockStoreMockRecorder) DeletePost(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockStore)(nil).DeletePost), ctx, id)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), ctx, id)
}

// ExecTx mocks base method.
func (m *MockStore) ExecTx(ctx context.Context, fn func(db.Querier) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecTx", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecTx indicates an expected call of ExecTx.
func (mr *MockStoreMockRecorder) ExecTx(ctx, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecTx", reflect.TypeOf((*MockStore)(nil).ExecTx), ctx, fn)
}

// FollowUser mocks base method.
func (m *MockStore) FollowUser(ctx context.Context, arg db.FollowUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FollowUser", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// FollowUser indicates an expected call of FollowUser.
func (mr *MockStoreMockRecorder) FollowUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FollowUser", reflect.TypeOf((*MockStore)(nil).FollowUser), ctx, arg)
}

// GetAllPosts mocks base method.
func (m *MockStore) GetAllPosts(ctx context.Context, arg db.GetAllPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPosts", ctx, arg)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPosts indicates an expected call of GetAllPosts.
func (mr *MockStoreMockRecorder) GetAllPosts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPosts", reflect.TypeOf((*MockStore)(nil).GetAllPosts), ctx, arg)
}

// GetCommentsByPost mocks base method.
func (m *MockStore) GetCommentsByPost(ctx context.Context, arg db.GetCommentsByPostParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsByPost", ctx, arg)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsByPost indicates an expected call of GetCommentsByPost.
func (mr *MockStoreMockRecorder) GetCommentsByPost(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsByPost", reflect.TypeOf((*MockStore)(nil).GetCommentsByPost), ctx, arg)
}

// GetFollowees mocks base method.
func (m *MockStore) GetFollowees(ctx context.Context, arg db.GetFolloweesParams) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowees", ctx, arg)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowees indicates an expected call of GetFollowees.
func (mr *MockStoreMockRecorder) GetFollowees(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowees", reflect.TypeOf((*MockStore)(nil).GetFollowees), ctx, arg)
}

// GetFollowers mocks base method.
func (m *MockStore) GetFollowers(ctx context.Context, arg db.GetFollowersParams) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowers", ctx, arg)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowers indicates an expected call of GetFollowers.
func (mr *MockStoreMockRecorder) GetFollowers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowers", reflect.TypeOf((*MockStore)(nil).GetFollowers), ctx, arg)
}

// GetPostByID mocks base method.
func (m *MockStore) GetPostByID(ctx context.Context, id int64) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", ctx, id)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockStoreMockRecorder) GetPostByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockStore)(nil).GetPostByID), ctx, id)
}

// GetPostsByUser mocks base method.
func (m *MockStore) GetPostsByUser(ctx context.Context, arg db.GetPostsByUserParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByUser", ctx, arg)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByUser indicates an expected call of GetPostsByUser.
func (mr *MockStoreMockRecorder) GetPostsByUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByUser", reflect.TypeOf((*MockStore)(nil).GetPostsByUser), ctx, arg)
}

// GetUserByID mocks base method.
func (m *MockStore) GetUserByID(ctx context.Context, id int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockStoreMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockStore)(nil).GetUserByID), ctx, id)
}

// SearchComments mocks base method.
func (m *MockStore) SearchComments(ctx context.Context, arg db.SearchCommentsParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchComments", ctx, arg)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchComments indicates an expected call of SearchComments.
func (mr *MockStoreMockRecorder) SearchComments(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchComments", reflect.TypeOf((*MockStore)(nil).SearchComments), ctx, arg)
}

// SearchPosts mocks base method.
func (m *MockStore) SearchPosts(ctx context.Context, arg db.SearchPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPosts", ctx, arg)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPosts indicates an expected call of SearchPosts.
func (mr *MockStoreMockRecorder) SearchPosts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPosts", reflect.TypeOf((*MockStore)(nil).SearchPosts), ctx, arg)
}

// SearchUsers mocks base method.
func (m *MockStore) SearchUsers(ctx context.Context, arg db.SearchUsersParams) ([]db.SearchUsersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsers", ctx, arg)
	ret0, _ := ret[0].([]db.SearchUsersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUsers indicates an expected call of SearchUsers.
func (mr *MockStoreMockRecorder) SearchUsers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsers", reflect.TypeOf((*MockStore)(nil).SearchUsers), ctx, arg)
}

// UnfollowUser mocks base method.
func (m *MockStore) UnfollowUser(ctx context.Context, arg db.UnfollowUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnfollowUser", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnfollowUser indicates an expected call of UnfollowUser.
func (mr *MockStoreMockRecorder) UnfollowUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnfollowUser", reflect.TypeOf((*MockStore)(nil).UnfollowUser), ctx, arg)
}

// UpdatePost mocks base method.
func (m *MockStore) UpdatePost(ctx context.Context, arg db.UpdatePostParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockStoreMockRecorder) UpdatePost(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockStore)(nil).UpdatePost), ctx, arg)
}

// UpdateUserPassword mocks base method.
func (m *MockStore) UpdateUserPassword(ctx context.Context, arg db.UpdateUserPasswordParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockStoreMockRecorder) UpdateUserPassword(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockStore)(nil).UpdateUserPassword), ctx, arg)
}

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
	isgomock struct{}
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreateComment mocks base method.
func (m *MockQuerier) CreateComment(ctx context.Context, arg db.CreateCommentParams) (db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", ctx, arg)
	ret0, _ := ret[0].(db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockQuerierMockRecorder) CreateComment(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockQuerier)(nil).CreateComment), ctx, arg)
}

// CreatePost mocks base method.
func (m *MockQuerier) CreatePost(ctx context.Context, arg db.CreatePostParams) (db.CreatePostRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, arg)
	ret0, _ := ret[0].(db.CreatePostRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockQuerierMockRecorder) CreatePost(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockQuerier)(nil).CreatePost), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), ctx, arg)
}

// DeleteComment mocks base method.
func (m *MockQuerier) DeleteComment(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockQuerierMockRecorder) DeleteComment(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockQuerier)(nil).DeleteComment), ctx, id)
}

// DeletePost mocks base method.
func (m *MockQuerier) DeletePost(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePost", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePost indicates an expected call of DeletePost.
func (mr *MockQuerierMockRecorder) DeletePost(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePost", reflect.TypeOf((*MockQuerier)(nil).DeletePost), ctx, id)
}

// DeleteUser mocks base method.
func (m *MockQuerier) DeleteUser(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockQuerierMockRecorder) DeleteUser(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockQuerier)(nil).DeleteUser), ctx, id)
}

// FollowUser mocks base method.
func (m *MockQuerier) FollowUser(ctx context.Context, arg db.FollowUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FollowUser", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// FollowUser indicates an expected call of FollowUser.
func (mr *MockQuerierMockRecorder) FollowUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FollowUser", reflect.TypeOf((*MockQuerier)(nil).FollowUser), ctx, arg)
}

// GetAllPosts mocks base method.
func (m *MockQuerier) GetAllPosts(ctx context.Context, arg db.GetAllPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPosts", ctx, arg)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPosts indicates an expected call of GetAllPosts.
func (mr *MockQuerierMockRecorder) GetAllPosts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPosts", reflect.TypeOf((*MockQuerier)(nil).GetAllPosts), ctx, arg)
}

// GetCommentsByPost mocks base method.
func (m *MockQuerier) GetCommentsByPost(ctx context.Context, arg db.GetCommentsByPostParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsByPost", ctx, arg)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsByPost indicates an expected call of GetCommentsByPost.
func (mr *MockQuerierMockRecorder) GetCommentsByPost(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsByPost", reflect.TypeOf((*MockQuerier)(nil).GetCommentsByPost), ctx, arg)
}

// GetFollowees mocks base method.
func (m *MockQuerier) GetFollowees(ctx context.Context, arg db.GetFolloweesParams) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowees", ctx, arg)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowees indicates an expected call of GetFollowees.
func (mr *MockQuerierMockRecorder) GetFollowees(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowees", reflect.TypeOf((*MockQuerier)(nil).GetFollowees), ctx, arg)
}

// GetFollowers mocks base method.
func (m *MockQuerier) GetFollowers(ctx context.Context, arg db.GetFollowersParams) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFollowers", ctx, arg)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFollowers indicates an expected call of GetFollowers.
func (mr *MockQuerierMockRecorder) GetFollowers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFollowers", reflect.TypeOf((*MockQuerier)(nil).GetFollowers), ctx, arg)
}

// GetPostByID mocks base method.
func (m *MockQuerier) GetPostByID(ctx context.Context, id int64) (db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostByID", ctx, id)
	ret0, _ := ret[0].(db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostByID indicates an expected call of GetPostByID.
func (mr *MockQuerierMockRecorder) GetPostByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostByID", reflect.TypeOf((*MockQuerier)(nil).GetPostByID), ctx, id)
}

// GetPostsByUser mocks base method.
func (m *MockQuerier) GetPostsByUser(ctx context.Context, arg db.GetPostsByUserParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPostsByUser", ctx, arg)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPostsByUser indicates an expected call of GetPostsByUser.
func (mr *MockQuerierMockRecorder) GetPostsByUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPostsByUser", reflect.TypeOf((*MockQuerier)(nil).GetPostsByUser), ctx, arg)
}

// GetUserByID mocks base method.
func (m *MockQuerier) GetUserByID(ctx context.Context, id int64) (db.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", ctx, id)
	ret0, _ := ret[0].(db.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockQuerierMockRecorder) GetUserByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockQuerier)(nil).GetUserByID), ctx, id)
}

// SearchComments mocks base method.
func (m *MockQuerier) SearchComments(ctx context.Context, arg db.SearchCommentsParams) ([]db.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchComments", ctx, arg)
	ret0, _ := ret[0].([]db.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchComments indicates an expected call of SearchComments.
func (mr *MockQuerierMockRecorder) SearchComments(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchComments", reflect.TypeOf((*MockQuerier)(nil).SearchComments), ctx, arg)
}

// SearchPosts mocks base method.
func (m *MockQuerier) SearchPosts(ctx context.Context, arg db.SearchPostsParams) ([]db.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPosts", ctx, arg)
	ret0, _ := ret[0].([]db.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPosts indicates an expected call of SearchPosts.
func (mr *MockQuerierMockRecorder) SearchPosts(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPosts", reflect.TypeOf((*MockQuerier)(nil).SearchPosts), ctx, arg)
}

// SearchUsers mocks base method.
func (m *MockQuerier) SearchUsers(ctx context.Context, arg db.SearchUsersParams) ([]db.SearchUsersRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchUsers", ctx, arg)
	ret0, _ := ret[0].([]db.SearchUsersRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchUsers indicates an expected call of SearchUsers.
func (mr *MockQuerierMockRecorder) SearchUsers(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchUsers", reflect.TypeOf((*MockQuerier)(nil).SearchUsers), ctx, arg)
}

// UnfollowUser mocks base method.
func (m *MockQuerier) UnfollowUser(ctx context.Context, arg db.UnfollowUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnfollowUser", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnfollowUser indicates an expected call of UnfollowUser.
func (mr *MockQuerierMockRecorder) UnfollowUser(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnfollowUser", reflect.TypeOf((*MockQuerier)(nil).UnfollowUser), ctx, arg)
}

// UpdatePost mocks base method.
func (m *MockQuerier) UpdatePost(ctx context.Context, arg db.UpdatePostParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockQuerierMockRecorder) UpdatePost(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockQuerier)(nil).UpdatePost), ctx, arg)
}

// UpdateUserPassword mocks base method.
func (m *MockQuerier) UpdateUserPassword(ctx context.Context, arg db.UpdateUserPasswordParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserPassword", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserPassword indicates an expected call of UpdateUserPassword.
func (mr *MockQuerierMockRecorder) UpdateUserPassword(ctx, arg any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserPassword", reflect.TypeOf((*MockQuerier)(nil).UpdateUserPassword), ctx, arg)
}
