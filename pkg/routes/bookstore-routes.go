package routes

import (
	"devtry.net/management_book/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(r *gin.Engine) {
	r.POST("/book", controllers.CreateBook)
	r.GET("/book", controllers.GetBooks)
	r.GET("/book/:id", controllers.GetBook)
	r.PUT("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)
}
