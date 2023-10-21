package routes

import (
	"todo/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine{
	r := gin.Default()

	// allow all origins
	r.Use(cors.Default())

	r.POST("/todos", controllers.AddTodos)
	r.GET("/todos", controllers.GetAllTodos)
	r.DELETE("/todo/:id", controllers.DeleteTodo)

	return r;
}