package repository

import (
	"library/models"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserPostgres_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockUsers(ctrl)

	expectedUsers := []models.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
	}

	mockDB.EXPECT().GetAll().Return(expectedUsers, nil)

	users, err := mockDB.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestUserPostgres_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockUsers(ctrl)

	newUser := models.User{Name: "New User"}

	mockDB.EXPECT().Create(newUser).Return(nil)

	err := mockDB.Create(newUser)
	assert.Nil(t, err)
}

func TestUserPostgres_GetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockUsers(ctrl)

	expectedUser := models.User{ID: 1, Name: "User 1"}

	mockDB.EXPECT().GetByID(1).Return(expectedUser, nil)

	user, err := mockDB.GetByID(1)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestUserPostgres_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockUsers(ctrl)

	userID := 1

	mockDB.EXPECT().Delete(userID).Return(nil)

	err := mockDB.Delete(userID)
	assert.NoError(t, err)
}

func TestUserPostgres_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDB := NewMockUsers(ctrl)

	updatedUser := models.User{ID: 1, Name: "Updated User"}

	mockDB.EXPECT().Update(updatedUser).Return(nil)

	err := mockDB.Update(updatedUser)
	assert.NoError(t, err)
}
