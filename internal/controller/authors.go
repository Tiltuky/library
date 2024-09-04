package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"library/models"
)

// GetAuthorByID @Summary Get Author by ID
// @Tags author
// @Description Get author details by ID
// @ID get-author-by-id
// @Accept  json
// @Produce  json
// @Param   id    path    int     true        "Author ID"
// @Success 200 {object} models.Author
// @Router /author/{id} [get]
func (h *Handler) GetAuthorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	author, err := h.Services.Authors.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, author)
}

// GetAllAuthors @Summary Get All Authors
// @Tags author
// @Description Get list of all authors
// @ID get-all-authors
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Author
// @Router /author [get]
func (h *Handler) GetAllAuthors(c *gin.Context) {
	authors, err := h.Services.Authors.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, authors)
}

// CreateAuthor @Summary Create Author
// @Tags author
// @Description Create a new author
// @ID create-author
// @Accept  json
// @Produce  json
// @Param   author  body    models.Author     true        "Author Info"
// @Success 201 {object} map[string]string "status: author created"
// @Router /author [post]
func (h *Handler) CreateAuthor(c *gin.Context) {
	var input models.Author
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.Services.Authors.Create(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "author created"})
}

// UpdateAuthor @Summary Update Author
// @Tags author
// @Description Update author details by ID
// @ID update-author
// @Accept  json
// @Produce  json
// @Param   id      path    int     true        "Author ID"
// @Param   author  body    models.Author     true        "Author Info"
// @Success 200 {object} map[string]string "status: author updated"
// @Router /author/{id} [put]
func (h *Handler) UpdateAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	var input models.Author
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	input.ID = id
	if err := h.Services.Authors.Update(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "author updated"})
}

// DeleteAuthor @Summary Delete Author
// @Tags author
// @Description Delete author by ID
// @ID delete-author
// @Accept  json
// @Produce  json
// @Param   id    path    int     true        "Author ID"
// @Success 200 {object} map[string]string "status: author deleted"
// @Router /author/{id} [delete]
func (h *Handler) DeleteAuthor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid author ID"})
		return
	}

	if err := h.Services.Authors.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "author deleted"})
}
