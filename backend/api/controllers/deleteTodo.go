package controllers

import (
	"net/http"
	"todo/api/services"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context){
	var todo models.TodoData

	id := c.Param("id")

	err := services.DeleteTodo(&todo,id)

	if(err != nil){
		c.JSON(http.StatusBadGateway,gin.H{
			"error": "could not delete data",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"message": "Data deleted successfully!",
		})
	}
}