package repository

import (
	"github.com/stretchr/testify/assert"
	"library/models"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAuthorPostgres_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockAuthors(ctrl)

	expectedAuthors := []models.Author{
		{ID: 1, Name: "Author 1"},
		{ID: 2, Name: "Author 2"},
	}

	mockDB.EXPECT().GetAll().Return(expectedAuthors, nil)

	authors, err := mockDB.GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, authors)

}

func TestAuthorPostgres_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockAuthors(ctrl)

	newAuthor := models.Author{Name: "New Author"}

	mockDB.EXPECT().Create(newAuthor).Return(nil)

	err := mockDB.Create(newAuthor)
	assert.Nil(t, err)

}

func TestAuthorPostgres_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockAuthors(ctrl)

	expectedAuthor := models.Author{ID: 1, Name: "Author 1"}

	mockDB.EXPECT().GetByID(1).Return(expectedAuthor, nil)

	author, err := mockDB.GetByID(1)
	assert.Nil(t, err)
	assert.Equal(t, expectedAuthor, author)

}

func TestAuthorPostgres_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockAuthors(ctrl)

	authorID := 1

	mockDB.EXPECT().Delete(authorID).Return(nil)

	err := mockDB.Delete(authorID)
	assert.Nil(t, err)

}

func TestAuthorPostgres_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockAuthors(ctrl)

	updatedAuthor := models.Author{ID: 1, Name: "Updated Author"}

	mockDB.EXPECT().Update(updatedAuthor).Return(nil)

	err := mockDB.Update(updatedAuthor)
	assert.NoError(t, err)

}
