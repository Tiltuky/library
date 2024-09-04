package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"library/models"
)

// GetUserByID @Summary Get User by ID
// @Tags users
// @Description Get user details by ID
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param   id    path    int     true        "User ID"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
func (h *Handler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, err := h.Services.Users.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers @Summary Get All Users
// @Tags users
// @Description Get list of all users
// @ID get-all-users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /user [get]
func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.Services.Users.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// CreateUser @Summary Create User
// @Tags users
// @Description Create a new user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param   user  body    models.User     true        "User Info"
// @Success 201 {object} map[string]string "status: user created"
// @Router /user [post]
func (h *Handler) CreateUser(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.Services.Users.Create(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "user created"})
}

// UpdateUser @Summary Update User
// @Tags users
// @Description Update user details by ID
// @ID update-user
// @Accept  json
// @Produce  json
// @Param   id      path    int     true        "User ID"
// @Param   user    body    models.User     true        "User Info"
// @Success 200 {object} map[string]string "status: user updated"
// @Router /user/{id} [put]
func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	var input models.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	input.ID = id
	if err := h.Services.Users.Update(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user updated"})
}

// DeleteUser @Summary Delete User
// @Tags users
// @Description Delete user by ID
// @ID delete-user
// @Accept  json
// @Produce  json
// @Param   id    path    int     true        "User ID"
// @Success 200 {object} map[string]string "status: user deleted"
// @Router /user/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	if err := h.Services.Users.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "user deleted"})
}
