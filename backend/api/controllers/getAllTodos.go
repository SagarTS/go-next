package controllers

import (
	"net/http"
	"todo/api/services"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func GetAllTodos(c *gin.Context){
	var todos []models.TodoData

	error := services.GetAllTodos(&todos)

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not fetch data",
		})
	} else {
		c.JSON(http.StatusOK, todos)
	}
}