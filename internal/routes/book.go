package routes

import (
	"github.com/f1rezy/book-store/internal/controllers"
	"github.com/f1rezy/book-store/internal/jwt"
	"github.com/gin-gonic/gin"
)

func BookRoutes(r *gin.Engine) {
	r.Use(jwt.IsAuthorized())
	r.POST("/books", controllers.CreateBook)
	r.GET("/books", controllers.GetAllBooks)
	r.GET("/books/:pk", controllers.GetBook)
}
