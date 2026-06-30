package routers

import (
	"github.com/gin-gonic/gin"
	"books-api/controllers"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/book", controllers.CreateBook)
	r.GET("/book", controllers.GetAllBook)
	r.GET("/book/:id", controllers.GetByIdBook)
	r.PUT("/book/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook	)
}