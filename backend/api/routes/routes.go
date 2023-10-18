package routes

import (
	"todo/api/controllers"

	"github.com/gin-gonic/gin"
)

func RouterSetup() *gin.Engine{
	r := gin.Default()
	r.POST("/todos", controllers.AddTodos)

	return r;
}