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

func setupRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func TestHandler_getAuthorByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorService := service.NewMockAuthors(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Authors: mockAuthorService,
		},
	}

	r := setupRouter()
	r.GET("/author/:id", handler.GetAuthorByID)

	expectedAuthor := models.Author{ID: 1, Name: "Author 1"}

	mockAuthorService.EXPECT().GetByID(1).Return(expectedAuthor, nil)

	req, _ := http.NewRequest("GET", "/author/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var author models.Author
	err := json.Unmarshal(w.Body.Bytes(), &author)
	assert.NoError(t, err)
	assert.Equal(t, expectedAuthor, author)
}

func TestHandler_getAuthorByID_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.GET("/author/:id", handler.GetAuthorByID)

	req, _ := http.NewRequest("GET", "/author/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid author ID")
}

func TestHandler_getAuthorByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorService := service.NewMockAuthors(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Authors: mockAuthorService,
		},
	}

	r := setupRouter()
	r.GET("/author/:id", handler.GetAuthorByID)

	mockAuthorService.EXPECT().GetByID(1).Return(models.Author{}, errors.New("author not found"))

	req, _ := http.NewRequest("GET", "/author/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "author not found")
}

func TestHandler_getAllAuthors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorService := service.NewMockAuthors(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Authors: mockAuthorService,
		},
	}

	r := setupRouter()
	r.GET("/author", handler.GetAllAuthors)

	expectedAuthors := []models.Author{
		{ID: 1, Name: "Author 1"},
		{ID: 2, Name: "Author 2"},
	}

	mockAuthorService.EXPECT().GetAll().Return(expectedAuthors, nil)

	req, _ := http.NewRequest("GET", "/author", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var authors []models.Author
	err := json.Unmarshal(w.Body.Bytes(), &authors)
	assert.NoError(t, err)
	assert.Equal(t, expectedAuthors, authors)
}

func TestHandler_createAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorService := service.NewMockAuthors(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Authors: mockAuthorService,
		},
	}

	r := setupRouter()
	r.POST("/author", handler.CreateAuthor)

	newAuthor := models.Author{Name: "New Author"}
	mockAuthorService.EXPECT().Create(newAuthor).Return(nil)

	authorJSON, _ := json.Marshal(newAuthor)
	req, _ := http.NewRequest("POST", "/author", bytes.NewBuffer(authorJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "author created")
}

func TestHandler_createAuthor_InvalidInput(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.POST("/author", handler.CreateAuthor)

	req, _ := http.NewRequest("POST", "/author", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")
}

func TestHandler_updateAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorService := service.NewMockAuthors(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Authors: mockAuthorService,
		},
	}

	r := setupRouter()
	r.PUT("/author/:id", handler.UpdateAuthor)

	updatedAuthor := models.Author{ID: 1, Name: "Updated Author"}
	mockAuthorService.EXPECT().Update(updatedAuthor).Return(nil)

	authorJSON, _ := json.Marshal(updatedAuthor)
	req, _ := http.NewRequest("PUT", "/author/1", bytes.NewBuffer(authorJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "author updated")
}

func TestHandler_updateAuthor_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.PUT("/author/:id", handler.UpdateAuthor)

	req, _ := http.NewRequest("PUT", "/author/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid author ID")
}

func TestHandler_updateAuthor_InvalidInput(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.PUT("/author/:id", handler.UpdateAuthor)

	req, _ := http.NewRequest("PUT", "/author/1", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")
}

func TestHandler_deleteAuthor(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthorService := service.NewMockAuthors(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Authors: mockAuthorService,
		},
	}

	r := setupRouter()
	r.DELETE("/author/:id", handler.DeleteAuthor)

	mockAuthorService.EXPECT().Delete(1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/author/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "author deleted")
}

func TestHandler_deleteAuthor_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.DELETE("/author/:id", handler.DeleteAuthor)

	req, _ := http.NewRequest("DELETE", "/author/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid author ID")
}
