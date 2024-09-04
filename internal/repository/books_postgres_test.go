package repository

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"library/models"
	"testing"
)

func TestBookPostgres_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockBooks(ctrl)

	expectedBooks := []models.Book{
		{ID: 1, Title: "Book 1", AuthorID: 1},
		{ID: 2, Title: "Book 2", AuthorID: 2},
	}

	mockDB.EXPECT().GetAll().Return(expectedBooks, nil)

	books, err := mockDB.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expectedBooks, books)
}

func TestBookPostgres_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockBooks(ctrl)

	newBook := models.Book{Title: "New Book", AuthorID: 1}

	mockDB.EXPECT().Create(newBook).Return(nil)

	err := mockDB.Create(newBook)
	assert.NoError(t, err)
}

func TestBookPostgres_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockBooks(ctrl)

	expectedBook := models.Book{ID: 1, Title: "Book 1", AuthorID: 1}

	mockDB.EXPECT().GetByID(1).Return(expectedBook, nil)

	book, err := mockDB.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedBook, book)
}

func TestBookPostgres_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockBooks(ctrl)

	bookID := 1

	mockDB.EXPECT().Delete(bookID).Return(nil)

	err := mockDB.Delete(bookID)
	assert.NoError(t, err)
}

func TestBookPostgres_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockBooks(ctrl)

	updatedBook := models.Book{ID: 1, Title: "Updated Book", AuthorID: 1}

	mockDB.EXPECT().Update(updatedBook).Return(nil)

	err := mockDB.Update(updatedBook)
	assert.NoError(t, err)
}

func TestBookPostgres_RentBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockBooks(ctrl)
	userID := 1
	bookID := 1

	// Test case where the book is already rented
	mockDB.EXPECT().RentBook(userID, bookID).Return(nil).Times(1)

	err := mockDB.RentBook(userID, bookID)
	assert.NoError(t, err)
}

func TestBookPostgres_ReturnBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockBooks(ctrl)

	userID := 1
	bookID := 1

	mockDB.EXPECT().ReturnBook(userID, bookID).Return(nil).Times(1)

	err := mockDB.ReturnBook(userID, bookID)
	assert.NoError(t, err)
}
