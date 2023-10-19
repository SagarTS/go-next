package routes

import (
	"todo/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine{
	r := gin.Default()
	r.Use(cors.Default())

	r.POST("/todos", controllers.AddTodos)

	return r;
}