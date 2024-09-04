package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"library/models"
)

// GetBookByID @Summary Get Book by ID
// @Tags books
// @Description Get book details by ID
// @ID get-book-by-id
// @Accept  json
// @Produce  json
// @Param   id    path    int     true        "Book ID"
// @Success 200 {object} models.Book
// @Router /book/{id} [get]
func (h *Handler) GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	book, err := h.Services.Books.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// GetAllBooks @Summary Get All Books
// @Tags books
// @Description Get list of all books
// @ID get-all-books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Failure 500 {object} map[string]string "internal server error"
// @Router /book [get]
func (h *Handler) GetAllBooks(c *gin.Context) {
	books, err := h.Services.Books.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

// CreateBook @Summary Create Book
// @Tags books
// @Description Create a new book
// @ID create-book
// @Accept  json
// @Produce  json
// @Param   book  body    models.Book     true        "Book Info"
// @Success 201 {object} map[string]string "status: book created"
// @Router /book [post]
func (h *Handler) CreateBook(c *gin.Context) {
	var input models.Book
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.Services.Books.Create(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "book created"})
}

// UpdateBook @Summary Update Book
// @Tags books
// @Description Update book details by ID
// @ID update-book
// @Accept  json
// @Produce  json
// @Param   id      path    int     true        "Book ID"
// @Param   book    body    models.Book     true        "Book Info"
// @Success 200 {object} map[string]string "status: book updated"
// @Router /book/{id} [put]
func (h *Handler) UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	var input models.Book
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	input.ID = id
	if err := h.Services.Books.Update(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "book updated"})
}

// DeleteBook @Summary Delete Book
// @Tags books
// @Description Delete book by ID
// @ID delete-book
// @Accept  json
// @Produce  json
// @Param   id    path    int     true        "Book ID"
// @Success 200 {object} map[string]string "status: book deleted"
// @Router /book/{id} [delete]
func (h *Handler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	if err := h.Services.Books.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "book deleted"})
}

type Input struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}

// RentBook @Summary Rent Book
// @Tags books
// @Description Rent a book
// @ID rent-book
// @Accept  json
// @Produce  json
// @Param   rent  body    Input    true        "Rent Info"
// @Success 200 {object} map[string]string "status: book rented"
// @Router /rent [post]
func (h *Handler) RentBook(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.Services.Books.RentBook(input.UserID, input.BookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "book rented"})
}

// ReturnBook @Summary Return Book
// @Tags books
// @Description Return a rented book
// @ID return-book
// @Accept  json
// @Produce  json
// @Param   return  body    Input     true        "Return Info"
// @Success 200 {object} map[string]string "status: book returned"
// @Router /rent/return [post]
func (h *Handler) ReturnBook(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.Services.Books.ReturnBook(input.UserID, input.BookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "book returned"})
}
