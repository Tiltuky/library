package controller_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"library/internal/controller"
	"library/internal/service"
	"library/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func setupRoute() *gin.Engine {
	r := gin.Default()
	return r
}

func TestHandler_getUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := service.NewMockUsers(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Users: mockUserService,
		},
	}

	r := setupRouter()
	r.GET("/user/:id", handler.GetUserByID)

	expectedUser := models.User{ID: 1, Name: "User 1"}

	mockUserService.EXPECT().GetByID(1).Return(expectedUser, nil)

	req, _ := http.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var user models.User
	err := json.Unmarshal(w.Body.Bytes(), &user)
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
}

func TestHandler_getUserByID_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.GET("/user/:id", handler.GetUserByID)

	req, _ := http.NewRequest("GET", "/user/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid user ID")
}

func TestHandler_getUserByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := service.NewMockUsers(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Users: mockUserService,
		},
	}

	r := setupRouter()
	r.GET("/user/:id", handler.GetUserByID)

	mockUserService.EXPECT().GetByID(1).Return(models.User{}, errors.New("user not found"))

	req, _ := http.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")
}

func TestHandler_getAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := service.NewMockUsers(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Users: mockUserService,
		},
	}

	r := setupRouter()
	r.GET("/user", handler.GetAllUsers)

	expectedUsers := []models.User{
		{ID: 1, Name: "User 1"},
		{ID: 2, Name: "User 2"},
	}

	mockUserService.EXPECT().GetAll().Return(expectedUsers, nil)

	req, _ := http.NewRequest("GET", "/user", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var users []models.User
	err := json.Unmarshal(w.Body.Bytes(), &users)
	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestHandler_createUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := service.NewMockUsers(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Users: mockUserService,
		},
	}

	r := setupRouter()
	r.POST("/user", handler.CreateUser)

	newUser := models.User{Name: "New User"}
	mockUserService.EXPECT().Create(newUser).Return(nil)

	userJSON, _ := json.Marshal(newUser)
	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "user created")
}

func TestHandler_createUser_InvalidInput(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.POST("/user", handler.CreateUser)

	req, _ := http.NewRequest("POST", "/user", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")
}

func TestHandler_updateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := service.NewMockUsers(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Users: mockUserService,
		},
	}

	r := setupRouter()
	r.PUT("/user/:id", handler.UpdateUser)

	updatedUser := models.User{ID: 1, Name: "Updated User"}
	mockUserService.EXPECT().Update(updatedUser).Return(nil)

	userJSON, _ := json.Marshal(updatedUser)
	req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user updated")
}

func TestHandler_updateUser_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.PUT("/user/:id", handler.UpdateUser)

	updatedUser := models.User{Name: "Updated User"}
	userJSON, _ := json.Marshal(updatedUser)
	req, _ := http.NewRequest("PUT", "/user/invalid", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid user ID")
}

func TestHandler_updateUser_InvalidInput(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.PUT("/user/:id", handler.UpdateUser)

	req, _ := http.NewRequest("PUT", "/user/1", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")
}

func TestHandler_deleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := service.NewMockUsers(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Users: mockUserService,
		},
	}

	r := setupRouter()
	r.DELETE("/user/:id", handler.DeleteUser)

	mockUserService.EXPECT().Delete(1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "user deleted")
}

func TestHandler_deleteUser_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.DELETE("/user/:id", handler.DeleteUser)

	req, _ := http.NewRequest("DELETE", "/user/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid user ID")
}

func TestHandler_deleteUser_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := service.NewMockUsers(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Users: mockUserService,
		},
	}

	r := setupRouter()
	r.DELETE("/user/:id", handler.DeleteUser)

	mockUserService.EXPECT().Delete(1).Return(errors.New("user not found"))

	req, _ := http.NewRequest("DELETE", "/user/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "user not found")
}
