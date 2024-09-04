// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package service is a generated GoMock package.
package service

import (
	models "library/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAuthors is a mock of Authors interface.
type MockAuthors struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorsMockRecorder
}

// MockAuthorsMockRecorder is the mock recorder for MockAuthors.
type MockAuthorsMockRecorder struct {
	mock *MockAuthors
}

// NewMockAuthors creates a new mock instance.
func NewMockAuthors(ctrl *gomock.Controller) *MockAuthors {
	mock := &MockAuthors{ctrl: ctrl}
	mock.recorder = &MockAuthorsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthors) EXPECT() *MockAuthorsMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockAuthors) Create(author models.Author) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", author)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockAuthorsMockRecorder) Create(author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthors)(nil).Create), author)
}

// Delete mocks base method.
func (m *MockAuthors) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthorsMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthors)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockAuthors) GetAll() ([]models.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockAuthorsMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockAuthors)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockAuthors) GetByID(id int) (models.Author, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.Author)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAuthorsMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAuthors)(nil).GetByID), id)
}

// Update mocks base method.
func (m *MockAuthors) Update(author models.Author) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", author)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockAuthorsMockRecorder) Update(author interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAuthors)(nil).Update), author)
}

// MockBooks is a mock of Books interface.
type MockBooks struct {
	ctrl     *gomock.Controller
	recorder *MockBooksMockRecorder
}

// MockBooksMockRecorder is the mock recorder for MockBooks.
type MockBooksMockRecorder struct {
	mock *MockBooks
}

// NewMockBooks creates a new mock instance.
func NewMockBooks(ctrl *gomock.Controller) *MockBooks {
	mock := &MockBooks{ctrl: ctrl}
	mock.recorder = &MockBooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBooks) EXPECT() *MockBooksMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockBooks) Create(book models.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", book)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockBooksMockRecorder) Create(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBooks)(nil).Create), book)
}

// Delete mocks base method.
func (m *MockBooks) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBooksMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBooks)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockBooks) GetAll() ([]models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockBooksMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBooks)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockBooks) GetByID(id int) (models.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockBooksMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockBooks)(nil).GetByID), id)
}

// RentBook mocks base method.
func (m *MockBooks) RentBook(userID, bookID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RentBook", userID, bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RentBook indicates an expected call of RentBook.
func (mr *MockBooksMockRecorder) RentBook(userID, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RentBook", reflect.TypeOf((*MockBooks)(nil).RentBook), userID, bookID)
}

// ReturnBook mocks base method.
func (m *MockBooks) ReturnBook(userID, bookID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReturnBook", userID, bookID)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReturnBook indicates an expected call of ReturnBook.
func (mr *MockBooksMockRecorder) ReturnBook(userID, bookID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReturnBook", reflect.TypeOf((*MockBooks)(nil).ReturnBook), userID, bookID)
}

// Update mocks base method.
func (m *MockBooks) Update(book models.Book) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", book)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockBooksMockRecorder) Update(book interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBooks)(nil).Update), book)
}

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUsers) Create(user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockUsersMockRecorder) Create(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUsers)(nil).Create), user)
}

// Delete mocks base method.
func (m *MockUsers) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUsersMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUsers)(nil).Delete), id)
}

// GetAll mocks base method.
func (m *MockUsers) GetAll() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockUsersMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockUsers)(nil).GetAll))
}

// GetByID mocks base method.
func (m *MockUsers) GetByID(id int) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", id)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockUsersMockRecorder) GetByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockUsers)(nil).GetByID), id)
}

// Update mocks base method.
func (m *MockUsers) Update(user models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUsersMockRecorder) Update(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUsers)(nil).Update), user)
}
