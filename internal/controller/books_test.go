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

func setupRout() *gin.Engine {
	r := gin.Default()
	return r
}

func TestHandler_getBookByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRout()
	r.GET("/book/:id", handler.GetBookByID)

	expectedBook := models.Book{ID: 1, Title: "Book 1", AuthorID: 1}

	mockBookService.EXPECT().GetByID(1).Return(expectedBook, nil)

	req, _ := http.NewRequest("GET", "/book/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var book models.Book
	err := json.Unmarshal(w.Body.Bytes(), &book)
	assert.NoError(t, err)
	assert.Equal(t, expectedBook, book)
}

func TestHandler_getBookByID_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRout()
	r.GET("/book/:id", handler.GetBookByID)

	req, _ := http.NewRequest("GET", "/book/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid book ID")
}

func TestHandler_getBookByID_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRout()
	r.GET("/book/:id", handler.GetBookByID)

	mockBookService.EXPECT().GetByID(1).Return(models.Book{}, errors.New("book not found"))

	req, _ := http.NewRequest("GET", "/book/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "book not found")
}

func TestHandler_getAllBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRout()
	r.GET("/book", handler.GetAllBooks)

	expectedBooks := []models.Book{
		{ID: 1, Title: "Book 1", AuthorID: 1},
		{ID: 2, Title: "Book 2", AuthorID: 2},
	}

	mockBookService.EXPECT().GetAll().Return(expectedBooks, nil)

	req, _ := http.NewRequest("GET", "/book", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var books []models.Book
	err := json.Unmarshal(w.Body.Bytes(), &books)
	assert.NoError(t, err)
	assert.Equal(t, expectedBooks, books)
}

func TestHandler_createBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRout()
	r.POST("/book", handler.CreateBook)

	newBook := models.Book{Title: "New Book", AuthorID: 1}
	mockBookService.EXPECT().Create(newBook).Return(nil)

	bookJSON, _ := json.Marshal(newBook)
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "book created")
}

func TestHandler_createBook_InvalidInput(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRout()
	r.POST("/book", handler.CreateBook)

	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")
}

func TestHandler_updateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRout()
	r.PUT("/book/:id", handler.UpdateBook)

	updatedBook := models.Book{ID: 1, Title: "Updated Book", AuthorID: 1}
	mockBookService.EXPECT().Update(updatedBook).Return(nil)

	bookJSON, _ := json.Marshal(updatedBook)
	req, _ := http.NewRequest("PUT", "/book/1", bytes.NewBuffer(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "book updated")
}

func TestHandler_updateBook_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.PUT("/book/:id", handler.UpdateBook)

	updatedBook := models.Book{Title: "Updated Book", AuthorID: 1}
	bookJSON, _ := json.Marshal(updatedBook)
	req, _ := http.NewRequest("PUT", "/book/invalid", bytes.NewBuffer(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid book ID")

}

func TestHandler_updateBook_InvalidInput(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.PUT("/book/:id", handler.UpdateBook)

	req, _ := http.NewRequest("PUT", "/book/1", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")

}

func TestHandler_deleteBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRouter()
	r.DELETE("/book/:id", handler.DeleteBook)

	mockBookService.EXPECT().Delete(1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/book/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "book deleted")

}

func TestHandler_deleteBook_InvalidID(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.DELETE("/book/:id", handler.DeleteBook)

	req, _ := http.NewRequest("DELETE", "/book/invalid", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid book ID")

}

func TestHandler_deleteBook_NotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRouter()
	r.DELETE("/book/:id", handler.DeleteBook)

	mockBookService.EXPECT().Delete(1).Return(errors.New("book not found"))

	req, _ := http.NewRequest("DELETE", "/book/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "book not found")

}

func TestHandler_rentBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRouter()
	r.POST("/rent", handler.RentBook)

	rentInfo := controller.Input{UserID: 1, BookID: 1}
	mockBookService.EXPECT().RentBook(rentInfo.UserID, rentInfo.BookID).Return(nil)

	rentJSON, _ := json.Marshal(rentInfo)
	req, _ := http.NewRequest("POST", "/rent", bytes.NewBuffer(rentJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "book rented")

}

func TestHandler_rentBook_InvalidInput(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.POST("/rent", handler.RentBook)

	req, _ := http.NewRequest("POST", "/rent", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")

}

func TestHandler_returnBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockBookService := service.NewMockBooks(ctrl)
	handler := &controller.Handler{
		Services: &service.Service{
			Books: mockBookService,
		},
	}

	r := setupRouter()
	r.POST("/rent/return", handler.ReturnBook)

	returnInfo := controller.Input{UserID: 1, BookID: 1}
	mockBookService.EXPECT().ReturnBook(returnInfo.UserID, returnInfo.BookID).Return(nil)

	returnJSON, _ := json.Marshal(returnInfo)
	req, _ := http.NewRequest("POST", "/rent/return", bytes.NewBuffer(returnJSON))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "book returned")

}

func TestHandler_returnBook_InvalidInput(t *testing.T) {
	handler := &controller.Handler{}

	r := setupRouter()
	r.POST("/rent/return", handler.ReturnBook)

	req, _ := http.NewRequest("POST", "/rent/return", bytes.NewBuffer([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "invalid input")
}
