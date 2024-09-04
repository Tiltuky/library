package controller

import (
	"github.com/gin-gonic/gin"
	"library/internal/service"
	"library/swagger"
	"net/http"

	_ "library/docs"
)

type Handler struct {
	Services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	// Swagger documentation route
	router.Static("/docs", "./docs")
	router.GET("/swagger/*any", gin.WrapH(http.HandlerFunc(swagger.SwaggerUI)))

	api := router.Group("/api")
	{
		authors := api.Group("/author")
		{
			authors.GET("/:id", h.GetAuthorByID)
			authors.GET("/", h.GetAllAuthors)
			authors.POST("/", h.CreateAuthor)
			authors.PUT("/:id", h.UpdateAuthor)
			authors.DELETE("/:id", h.DeleteAuthor)
		}

		books := api.Group("/book")
		{
			books.GET("/:id", h.GetBookByID)
			books.GET("/", h.GetAllBooks)
			books.POST("/", h.CreateBook)
			books.PUT("/:id", h.UpdateBook)
			books.DELETE("/:id", h.DeleteBook)
		}

		users := api.Group("/user")
		{
			users.GET("/:id", h.GetUserByID)
			users.GET("/", h.GetAllUsers)
			users.POST("/", h.CreateUser)
			users.PUT("/:id", h.UpdateUser)
			users.DELETE("/:id", h.DeleteUser)
		}

		rent := api.Group("/rent")
		{
			rent.POST("/", h.RentBook)
			rent.POST("/return", h.ReturnBook)
		}
	}

	return router
}
