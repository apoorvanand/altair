// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	entity "github.com/kodefluence/altair/provider/plugin/oauth/entity"
	db "github.com/kodefluence/monorepo/db"
	exception "github.com/kodefluence/monorepo/exception"
	kontext "github.com/kodefluence/monorepo/kontext"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOauthApplicationModel is a mock of OauthApplicationModel interface
type MockOauthApplicationModel struct {
	ctrl     *gomock.Controller
	recorder *MockOauthApplicationModelMockRecorder
}

// MockOauthApplicationModelMockRecorder is the mock recorder for MockOauthApplicationModel
type MockOauthApplicationModelMockRecorder struct {
	mock *MockOauthApplicationModel
}

// NewMockOauthApplicationModel creates a new mock instance
func NewMockOauthApplicationModel(ctrl *gomock.Controller) *MockOauthApplicationModel {
	mock := &MockOauthApplicationModel{ctrl: ctrl}
	mock.recorder = &MockOauthApplicationModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthApplicationModel) EXPECT() *MockOauthApplicationModelMockRecorder {
	return m.recorder
}

// Paginate mocks base method
func (m *MockOauthApplicationModel) Paginate(ktx kontext.Context, offset, limit int, tx db.TX) ([]entity.OauthApplication, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Paginate", ktx, offset, limit, tx)
	ret0, _ := ret[0].([]entity.OauthApplication)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// Paginate indicates an expected call of Paginate
func (mr *MockOauthApplicationModelMockRecorder) Paginate(ktx, offset, limit, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Paginate", reflect.TypeOf((*MockOauthApplicationModel)(nil).Paginate), ktx, offset, limit, tx)
}

// Count mocks base method
func (m *MockOauthApplicationModel) Count(ktx kontext.Context, tx db.TX) (int, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ktx, tx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockOauthApplicationModelMockRecorder) Count(ktx, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockOauthApplicationModel)(nil).Count), ktx, tx)
}

// One mocks base method
func (m *MockOauthApplicationModel) One(ktx kontext.Context, ID int, tx db.TX) (entity.OauthApplication, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ktx, ID, tx)
	ret0, _ := ret[0].(entity.OauthApplication)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockOauthApplicationModelMockRecorder) One(ktx, ID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOauthApplicationModel)(nil).One), ktx, ID, tx)
}

// OneByUIDandSecret mocks base method
func (m *MockOauthApplicationModel) OneByUIDandSecret(ktx kontext.Context, clientUID, clientSecret string, tx db.TX) (entity.OauthApplication, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneByUIDandSecret", ktx, clientUID, clientSecret, tx)
	ret0, _ := ret[0].(entity.OauthApplication)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// OneByUIDandSecret indicates an expected call of OneByUIDandSecret
func (mr *MockOauthApplicationModelMockRecorder) OneByUIDandSecret(ktx, clientUID, clientSecret, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneByUIDandSecret", reflect.TypeOf((*MockOauthApplicationModel)(nil).OneByUIDandSecret), ktx, clientUID, clientSecret, tx)
}

// Create mocks base method
func (m *MockOauthApplicationModel) Create(ktx kontext.Context, data entity.OauthApplicationInsertable, tx db.TX) (int, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ktx, data, tx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOauthApplicationModelMockRecorder) Create(ktx, data, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOauthApplicationModel)(nil).Create), ktx, data, tx)
}

// Update mocks base method
func (m *MockOauthApplicationModel) Update(ktx kontext.Context, ID int, data entity.OauthApplicationUpdateable, tx db.TX) exception.Exception {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ktx, ID, data, tx)
	ret0, _ := ret[0].(exception.Exception)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockOauthApplicationModelMockRecorder) Update(ktx, ID, data, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockOauthApplicationModel)(nil).Update), ktx, ID, data, tx)
}

// MockOauthAccessTokenModel is a mock of OauthAccessTokenModel interface
type MockOauthAccessTokenModel struct {
	ctrl     *gomock.Controller
	recorder *MockOauthAccessTokenModelMockRecorder
}

// MockOauthAccessTokenModelMockRecorder is the mock recorder for MockOauthAccessTokenModel
type MockOauthAccessTokenModelMockRecorder struct {
	mock *MockOauthAccessTokenModel
}

// NewMockOauthAccessTokenModel creates a new mock instance
func NewMockOauthAccessTokenModel(ctrl *gomock.Controller) *MockOauthAccessTokenModel {
	mock := &MockOauthAccessTokenModel{ctrl: ctrl}
	mock.recorder = &MockOauthAccessTokenModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthAccessTokenModel) EXPECT() *MockOauthAccessTokenModelMockRecorder {
	return m.recorder
}

// OneByToken mocks base method
func (m *MockOauthAccessTokenModel) OneByToken(ktx kontext.Context, token string, tx db.TX) (entity.OauthAccessToken, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneByToken", ktx, token, tx)
	ret0, _ := ret[0].(entity.OauthAccessToken)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// OneByToken indicates an expected call of OneByToken
func (mr *MockOauthAccessTokenModelMockRecorder) OneByToken(ktx, token, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneByToken", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).OneByToken), ktx, token, tx)
}

// One mocks base method
func (m *MockOauthAccessTokenModel) One(ktx kontext.Context, ID int, tx db.TX) (entity.OauthAccessToken, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ktx, ID, tx)
	ret0, _ := ret[0].(entity.OauthAccessToken)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockOauthAccessTokenModelMockRecorder) One(ktx, ID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).One), ktx, ID, tx)
}

// Create mocks base method
func (m *MockOauthAccessTokenModel) Create(ktx kontext.Context, data entity.OauthAccessTokenInsertable, tx db.TX) (int, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ktx, data, tx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOauthAccessTokenModelMockRecorder) Create(ktx, data, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).Create), ktx, data, tx)
}

// Revoke mocks base method
func (m *MockOauthAccessTokenModel) Revoke(ktx kontext.Context, token string, tx db.TX) exception.Exception {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", ktx, token, tx)
	ret0, _ := ret[0].(exception.Exception)
	return ret0
}

// Revoke indicates an expected call of Revoke
func (mr *MockOauthAccessTokenModelMockRecorder) Revoke(ktx, token, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockOauthAccessTokenModel)(nil).Revoke), ktx, token, tx)
}

// MockOauthAccessGrantModel is a mock of OauthAccessGrantModel interface
type MockOauthAccessGrantModel struct {
	ctrl     *gomock.Controller
	recorder *MockOauthAccessGrantModelMockRecorder
}

// MockOauthAccessGrantModelMockRecorder is the mock recorder for MockOauthAccessGrantModel
type MockOauthAccessGrantModelMockRecorder struct {
	mock *MockOauthAccessGrantModel
}

// NewMockOauthAccessGrantModel creates a new mock instance
func NewMockOauthAccessGrantModel(ctrl *gomock.Controller) *MockOauthAccessGrantModel {
	mock := &MockOauthAccessGrantModel{ctrl: ctrl}
	mock.recorder = &MockOauthAccessGrantModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthAccessGrantModel) EXPECT() *MockOauthAccessGrantModelMockRecorder {
	return m.recorder
}

// One mocks base method
func (m *MockOauthAccessGrantModel) One(ktx kontext.Context, ID int, tx db.TX) (entity.OauthAccessGrant, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ktx, ID, tx)
	ret0, _ := ret[0].(entity.OauthAccessGrant)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockOauthAccessGrantModelMockRecorder) One(ktx, ID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).One), ktx, ID, tx)
}

// OneByCode mocks base method
func (m *MockOauthAccessGrantModel) OneByCode(ktx kontext.Context, code string, tx db.TX) (entity.OauthAccessGrant, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneByCode", ktx, code, tx)
	ret0, _ := ret[0].(entity.OauthAccessGrant)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// OneByCode indicates an expected call of OneByCode
func (mr *MockOauthAccessGrantModelMockRecorder) OneByCode(ktx, code, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneByCode", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).OneByCode), ktx, code, tx)
}

// Create mocks base method
func (m *MockOauthAccessGrantModel) Create(ktx kontext.Context, data entity.OauthAccessGrantInsertable, tx db.TX) (int, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ktx, data, tx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOauthAccessGrantModelMockRecorder) Create(ktx, data, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).Create), ktx, data, tx)
}

// Revoke mocks base method
func (m *MockOauthAccessGrantModel) Revoke(ktx kontext.Context, code string, tx db.TX) exception.Exception {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", ktx, code, tx)
	ret0, _ := ret[0].(exception.Exception)
	return ret0
}

// Revoke indicates an expected call of Revoke
func (mr *MockOauthAccessGrantModelMockRecorder) Revoke(ktx, code, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockOauthAccessGrantModel)(nil).Revoke), ktx, code, tx)
}

// MockOauthRefreshTokenModel is a mock of OauthRefreshTokenModel interface
type MockOauthRefreshTokenModel struct {
	ctrl     *gomock.Controller
	recorder *MockOauthRefreshTokenModelMockRecorder
}

// MockOauthRefreshTokenModelMockRecorder is the mock recorder for MockOauthRefreshTokenModel
type MockOauthRefreshTokenModelMockRecorder struct {
	mock *MockOauthRefreshTokenModel
}

// NewMockOauthRefreshTokenModel creates a new mock instance
func NewMockOauthRefreshTokenModel(ctrl *gomock.Controller) *MockOauthRefreshTokenModel {
	mock := &MockOauthRefreshTokenModel{ctrl: ctrl}
	mock.recorder = &MockOauthRefreshTokenModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthRefreshTokenModel) EXPECT() *MockOauthRefreshTokenModelMockRecorder {
	return m.recorder
}

// OneByToken mocks base method
func (m *MockOauthRefreshTokenModel) OneByToken(ktx kontext.Context, token string, tx db.TX) (entity.OauthRefreshToken, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OneByToken", ktx, token, tx)
	ret0, _ := ret[0].(entity.OauthRefreshToken)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// OneByToken indicates an expected call of OneByToken
func (mr *MockOauthRefreshTokenModelMockRecorder) OneByToken(ktx, token, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OneByToken", reflect.TypeOf((*MockOauthRefreshTokenModel)(nil).OneByToken), ktx, token, tx)
}

// One mocks base method
func (m *MockOauthRefreshTokenModel) One(ktx kontext.Context, ID int, tx db.TX) (entity.OauthRefreshToken, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ktx, ID, tx)
	ret0, _ := ret[0].(entity.OauthRefreshToken)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockOauthRefreshTokenModelMockRecorder) One(ktx, ID, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockOauthRefreshTokenModel)(nil).One), ktx, ID, tx)
}

// Create mocks base method
func (m *MockOauthRefreshTokenModel) Create(ktx kontext.Context, data entity.OauthRefreshTokenInsertable, tx db.TX) (int, exception.Exception) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ktx, data, tx)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(exception.Exception)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockOauthRefreshTokenModelMockRecorder) Create(ktx, data, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOauthRefreshTokenModel)(nil).Create), ktx, data, tx)
}

// Revoke mocks base method
func (m *MockOauthRefreshTokenModel) Revoke(ktx kontext.Context, token string, tx db.TX) exception.Exception {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", ktx, token, tx)
	ret0, _ := ret[0].(exception.Exception)
	return ret0
}

// Revoke indicates an expected call of Revoke
func (mr *MockOauthRefreshTokenModelMockRecorder) Revoke(ktx, token, tx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockOauthRefreshTokenModel)(nil).Revoke), ktx, token, tx)
}

// MockApplicationManager is a mock of ApplicationManager interface
type MockApplicationManager struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationManagerMockRecorder
}

// MockApplicationManagerMockRecorder is the mock recorder for MockApplicationManager
type MockApplicationManagerMockRecorder struct {
	mock *MockApplicationManager
}

// NewMockApplicationManager creates a new mock instance
func NewMockApplicationManager(ctrl *gomock.Controller) *MockApplicationManager {
	mock := &MockApplicationManager{ctrl: ctrl}
	mock.recorder = &MockApplicationManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockApplicationManager) EXPECT() *MockApplicationManagerMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockApplicationManager) List(ctx context.Context, offset, limit int) ([]entity.OauthApplicationJSON, int, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, offset, limit)
	ret0, _ := ret[0].([]entity.OauthApplicationJSON)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(*entity.Error)
	return ret0, ret1, ret2
}

// List indicates an expected call of List
func (mr *MockApplicationManagerMockRecorder) List(ctx, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockApplicationManager)(nil).List), ctx, offset, limit)
}

// One mocks base method
func (m *MockApplicationManager) One(ctx context.Context, ID int) (entity.OauthApplicationJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "One", ctx, ID)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// One indicates an expected call of One
func (mr *MockApplicationManagerMockRecorder) One(ctx, ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockApplicationManager)(nil).One), ctx, ID)
}

// Create mocks base method
func (m *MockApplicationManager) Create(ctx context.Context, e entity.OauthApplicationJSON) (entity.OauthApplicationJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, e)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockApplicationManagerMockRecorder) Create(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockApplicationManager)(nil).Create), ctx, e)
}

// Update mocks base method
func (m *MockApplicationManager) Update(ctx context.Context, ID int, e entity.OauthApplicationUpdateJSON) (entity.OauthApplicationJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, ID, e)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockApplicationManagerMockRecorder) Update(ctx, ID, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockApplicationManager)(nil).Update), ctx, ID, e)
}

// MockAuthorization is a mock of Authorization interface
type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

// MockAuthorizationMockRecorder is the mock recorder for MockAuthorization
type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

// NewMockAuthorization creates a new mock instance
func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

// Grantor mocks base method
func (m *MockAuthorization) Grantor(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (interface{}, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Grantor", ctx, authorizationReq)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Grantor indicates an expected call of Grantor
func (mr *MockAuthorizationMockRecorder) Grantor(ctx, authorizationReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Grantor", reflect.TypeOf((*MockAuthorization)(nil).Grantor), ctx, authorizationReq)
}

// Grant mocks base method
func (m *MockAuthorization) Grant(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (entity.OauthAccessGrantJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Grant", ctx, authorizationReq)
	ret0, _ := ret[0].(entity.OauthAccessGrantJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Grant indicates an expected call of Grant
func (mr *MockAuthorizationMockRecorder) Grant(ctx, authorizationReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Grant", reflect.TypeOf((*MockAuthorization)(nil).Grant), ctx, authorizationReq)
}

// GrantToken mocks base method
func (m *MockAuthorization) GrantToken(ctx context.Context, authorizationReq entity.AuthorizationRequestJSON) (entity.OauthAccessTokenJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantToken", ctx, authorizationReq)
	ret0, _ := ret[0].(entity.OauthAccessTokenJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// GrantToken indicates an expected call of GrantToken
func (mr *MockAuthorizationMockRecorder) GrantToken(ctx, authorizationReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantToken", reflect.TypeOf((*MockAuthorization)(nil).GrantToken), ctx, authorizationReq)
}

// Token mocks base method
func (m *MockAuthorization) Token(ctx context.Context, accessTokenReq entity.AccessTokenRequestJSON) (entity.OauthAccessTokenJSON, *entity.Error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token", ctx, accessTokenReq)
	ret0, _ := ret[0].(entity.OauthAccessTokenJSON)
	ret1, _ := ret[1].(*entity.Error)
	return ret0, ret1
}

// Token indicates an expected call of Token
func (mr *MockAuthorizationMockRecorder) Token(ctx, accessTokenReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockAuthorization)(nil).Token), ctx, accessTokenReq)
}

// RevokeToken mocks base method
func (m *MockAuthorization) RevokeToken(ctx context.Context, revokeAccessTokenReq entity.RevokeAccessTokenRequestJSON) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeToken", ctx, revokeAccessTokenReq)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// RevokeToken indicates an expected call of RevokeToken
func (mr *MockAuthorizationMockRecorder) RevokeToken(ctx, revokeAccessTokenReq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeToken", reflect.TypeOf((*MockAuthorization)(nil).RevokeToken), ctx, revokeAccessTokenReq)
}

// MockOauthApplicationFormater is a mock of OauthApplicationFormater interface
type MockOauthApplicationFormater struct {
	ctrl     *gomock.Controller
	recorder *MockOauthApplicationFormaterMockRecorder
}

// MockOauthApplicationFormaterMockRecorder is the mock recorder for MockOauthApplicationFormater
type MockOauthApplicationFormaterMockRecorder struct {
	mock *MockOauthApplicationFormater
}

// NewMockOauthApplicationFormater creates a new mock instance
func NewMockOauthApplicationFormater(ctrl *gomock.Controller) *MockOauthApplicationFormater {
	mock := &MockOauthApplicationFormater{ctrl: ctrl}
	mock.recorder = &MockOauthApplicationFormaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthApplicationFormater) EXPECT() *MockOauthApplicationFormaterMockRecorder {
	return m.recorder
}

// ApplicationList mocks base method
func (m *MockOauthApplicationFormater) ApplicationList(ctx context.Context, applications []entity.OauthApplication) []entity.OauthApplicationJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplicationList", ctx, applications)
	ret0, _ := ret[0].([]entity.OauthApplicationJSON)
	return ret0
}

// ApplicationList indicates an expected call of ApplicationList
func (mr *MockOauthApplicationFormaterMockRecorder) ApplicationList(ctx, applications interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplicationList", reflect.TypeOf((*MockOauthApplicationFormater)(nil).ApplicationList), ctx, applications)
}

// Application mocks base method
func (m *MockOauthApplicationFormater) Application(ctx context.Context, application entity.OauthApplication) entity.OauthApplicationJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Application", ctx, application)
	ret0, _ := ret[0].(entity.OauthApplicationJSON)
	return ret0
}

// Application indicates an expected call of Application
func (mr *MockOauthApplicationFormaterMockRecorder) Application(ctx, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockOauthApplicationFormater)(nil).Application), ctx, application)
}

// MockOauthFormatter is a mock of OauthFormatter interface
type MockOauthFormatter struct {
	ctrl     *gomock.Controller
	recorder *MockOauthFormatterMockRecorder
}

// MockOauthFormatterMockRecorder is the mock recorder for MockOauthFormatter
type MockOauthFormatterMockRecorder struct {
	mock *MockOauthFormatter
}

// NewMockOauthFormatter creates a new mock instance
func NewMockOauthFormatter(ctrl *gomock.Controller) *MockOauthFormatter {
	mock := &MockOauthFormatter{ctrl: ctrl}
	mock.recorder = &MockOauthFormatterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthFormatter) EXPECT() *MockOauthFormatterMockRecorder {
	return m.recorder
}

// AccessGrant mocks base method
func (m *MockOauthFormatter) AccessGrant(e entity.OauthAccessGrant) entity.OauthAccessGrantJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessGrant", e)
	ret0, _ := ret[0].(entity.OauthAccessGrantJSON)
	return ret0
}

// AccessGrant indicates an expected call of AccessGrant
func (mr *MockOauthFormatterMockRecorder) AccessGrant(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessGrant", reflect.TypeOf((*MockOauthFormatter)(nil).AccessGrant), e)
}

// AccessToken mocks base method
func (m *MockOauthFormatter) AccessToken(e entity.OauthAccessToken, redirectURI string, refreshTokenJSON *entity.OauthRefreshTokenJSON) entity.OauthAccessTokenJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessToken", e, redirectURI, refreshTokenJSON)
	ret0, _ := ret[0].(entity.OauthAccessTokenJSON)
	return ret0
}

// AccessToken indicates an expected call of AccessToken
func (mr *MockOauthFormatterMockRecorder) AccessToken(e, redirectURI, refreshTokenJSON interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessToken", reflect.TypeOf((*MockOauthFormatter)(nil).AccessToken), e, redirectURI, refreshTokenJSON)
}

// RefreshToken mocks base method
func (m *MockOauthFormatter) RefreshToken(e entity.OauthRefreshToken) entity.OauthRefreshTokenJSON {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", e)
	ret0, _ := ret[0].(entity.OauthRefreshTokenJSON)
	return ret0
}

// RefreshToken indicates an expected call of RefreshToken
func (mr *MockOauthFormatterMockRecorder) RefreshToken(e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockOauthFormatter)(nil).RefreshToken), e)
}

// MockModelFormater is a mock of ModelFormater interface
type MockModelFormater struct {
	ctrl     *gomock.Controller
	recorder *MockModelFormaterMockRecorder
}

// MockModelFormaterMockRecorder is the mock recorder for MockModelFormater
type MockModelFormaterMockRecorder struct {
	mock *MockModelFormater
}

// NewMockModelFormater creates a new mock instance
func NewMockModelFormater(ctrl *gomock.Controller) *MockModelFormater {
	mock := &MockModelFormater{ctrl: ctrl}
	mock.recorder = &MockModelFormaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelFormater) EXPECT() *MockModelFormaterMockRecorder {
	return m.recorder
}

// AccessTokenFromAuthorizationRequest mocks base method
func (m *MockModelFormater) AccessTokenFromAuthorizationRequest(r entity.AuthorizationRequestJSON, application entity.OauthApplication) entity.OauthAccessTokenInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenFromAuthorizationRequest", r, application)
	ret0, _ := ret[0].(entity.OauthAccessTokenInsertable)
	return ret0
}

// AccessTokenFromAuthorizationRequest indicates an expected call of AccessTokenFromAuthorizationRequest
func (mr *MockModelFormaterMockRecorder) AccessTokenFromAuthorizationRequest(r, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenFromAuthorizationRequest", reflect.TypeOf((*MockModelFormater)(nil).AccessTokenFromAuthorizationRequest), r, application)
}

// AccessTokenFromOauthAccessGrant mocks base method
func (m *MockModelFormater) AccessTokenFromOauthAccessGrant(oauthAccessGrant entity.OauthAccessGrant, application entity.OauthApplication) entity.OauthAccessTokenInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenFromOauthAccessGrant", oauthAccessGrant, application)
	ret0, _ := ret[0].(entity.OauthAccessTokenInsertable)
	return ret0
}

// AccessTokenFromOauthAccessGrant indicates an expected call of AccessTokenFromOauthAccessGrant
func (mr *MockModelFormaterMockRecorder) AccessTokenFromOauthAccessGrant(oauthAccessGrant, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenFromOauthAccessGrant", reflect.TypeOf((*MockModelFormater)(nil).AccessTokenFromOauthAccessGrant), oauthAccessGrant, application)
}

// AccessGrantFromAuthorizationRequest mocks base method
func (m *MockModelFormater) AccessGrantFromAuthorizationRequest(r entity.AuthorizationRequestJSON, application entity.OauthApplication) entity.OauthAccessGrantInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessGrantFromAuthorizationRequest", r, application)
	ret0, _ := ret[0].(entity.OauthAccessGrantInsertable)
	return ret0
}

// AccessGrantFromAuthorizationRequest indicates an expected call of AccessGrantFromAuthorizationRequest
func (mr *MockModelFormaterMockRecorder) AccessGrantFromAuthorizationRequest(r, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessGrantFromAuthorizationRequest", reflect.TypeOf((*MockModelFormater)(nil).AccessGrantFromAuthorizationRequest), r, application)
}

// OauthApplication mocks base method
func (m *MockModelFormater) OauthApplication(r entity.OauthApplicationJSON) entity.OauthApplicationInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OauthApplication", r)
	ret0, _ := ret[0].(entity.OauthApplicationInsertable)
	return ret0
}

// OauthApplication indicates an expected call of OauthApplication
func (mr *MockModelFormaterMockRecorder) OauthApplication(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OauthApplication", reflect.TypeOf((*MockModelFormater)(nil).OauthApplication), r)
}

// AccessTokenFromOauthRefreshToken mocks base method
func (m *MockModelFormater) AccessTokenFromOauthRefreshToken(application entity.OauthApplication, accessToken entity.OauthAccessToken) entity.OauthAccessTokenInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AccessTokenFromOauthRefreshToken", application, accessToken)
	ret0, _ := ret[0].(entity.OauthAccessTokenInsertable)
	return ret0
}

// AccessTokenFromOauthRefreshToken indicates an expected call of AccessTokenFromOauthRefreshToken
func (mr *MockModelFormaterMockRecorder) AccessTokenFromOauthRefreshToken(application, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AccessTokenFromOauthRefreshToken", reflect.TypeOf((*MockModelFormater)(nil).AccessTokenFromOauthRefreshToken), application, accessToken)
}

// RefreshToken mocks base method
func (m *MockModelFormater) RefreshToken(application entity.OauthApplication, accessToken entity.OauthAccessToken) entity.OauthRefreshTokenInsertable {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", application, accessToken)
	ret0, _ := ret[0].(entity.OauthRefreshTokenInsertable)
	return ret0
}

// RefreshToken indicates an expected call of RefreshToken
func (mr *MockModelFormaterMockRecorder) RefreshToken(application, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockModelFormater)(nil).RefreshToken), application, accessToken)
}

// MockOauthValidator is a mock of OauthValidator interface
type MockOauthValidator struct {
	ctrl     *gomock.Controller
	recorder *MockOauthValidatorMockRecorder
}

// MockOauthValidatorMockRecorder is the mock recorder for MockOauthValidator
type MockOauthValidatorMockRecorder struct {
	mock *MockOauthValidator
}

// NewMockOauthValidator creates a new mock instance
func NewMockOauthValidator(ctrl *gomock.Controller) *MockOauthValidator {
	mock := &MockOauthValidator{ctrl: ctrl}
	mock.recorder = &MockOauthValidatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOauthValidator) EXPECT() *MockOauthValidatorMockRecorder {
	return m.recorder
}

// ValidateApplication mocks base method
func (m *MockOauthValidator) ValidateApplication(ctx context.Context, data entity.OauthApplicationJSON) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateApplication", ctx, data)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// ValidateApplication indicates an expected call of ValidateApplication
func (mr *MockOauthValidatorMockRecorder) ValidateApplication(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateApplication", reflect.TypeOf((*MockOauthValidator)(nil).ValidateApplication), ctx, data)
}

// ValidateAuthorizationGrant mocks base method
func (m *MockOauthValidator) ValidateAuthorizationGrant(ctx context.Context, r entity.AuthorizationRequestJSON, application entity.OauthApplication) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateAuthorizationGrant", ctx, r, application)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// ValidateAuthorizationGrant indicates an expected call of ValidateAuthorizationGrant
func (mr *MockOauthValidatorMockRecorder) ValidateAuthorizationGrant(ctx, r, application interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateAuthorizationGrant", reflect.TypeOf((*MockOauthValidator)(nil).ValidateAuthorizationGrant), ctx, r, application)
}

// ValidateTokenGrant mocks base method
func (m *MockOauthValidator) ValidateTokenGrant(ctx context.Context, r entity.AccessTokenRequestJSON) *entity.Error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTokenGrant", ctx, r)
	ret0, _ := ret[0].(*entity.Error)
	return ret0
}

// ValidateTokenGrant indicates an expected call of ValidateTokenGrant
func (mr *MockOauthValidatorMockRecorder) ValidateTokenGrant(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTokenGrant", reflect.TypeOf((*MockOauthValidator)(nil).ValidateTokenGrant), ctx, r)
}

// ValidateTokenRefreshToken mocks base method
func (m *MockOauthValidator) ValidateTokenRefreshToken(ctx context.Context, data entity.OauthRefreshToken) exception.Exception {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTokenRefreshToken", ctx, data)
	ret0, _ := ret[0].(exception.Exception)
	return ret0
}

// ValidateTokenRefreshToken indicates an expected call of ValidateTokenRefreshToken
func (mr *MockOauthValidatorMockRecorder) ValidateTokenRefreshToken(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTokenRefreshToken", reflect.TypeOf((*MockOauthValidator)(nil).ValidateTokenRefreshToken), ctx, data)
}

// ValidateTokenAuthorizationCode mocks base method
func (m *MockOauthValidator) ValidateTokenAuthorizationCode(ctx context.Context, r entity.AccessTokenRequestJSON, data entity.OauthAccessGrant) exception.Exception {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateTokenAuthorizationCode", ctx, r, data)
	ret0, _ := ret[0].(exception.Exception)
	return ret0
}

// ValidateTokenAuthorizationCode indicates an expected call of ValidateTokenAuthorizationCode
func (mr *MockOauthValidatorMockRecorder) ValidateTokenAuthorizationCode(ctx, r, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateTokenAuthorizationCode", reflect.TypeOf((*MockOauthValidator)(nil).ValidateTokenAuthorizationCode), ctx, r, data)
}
