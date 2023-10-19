package controllers

import (
	"net/http"
	"todo/api/services"
	"todo/models"

	"github.com/gin-gonic/gin"
)

func AddTodos(c *gin.Context){
	var body models.TodoData

	err:= c.BindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not get data",
		})
	}

	error := services.AddTodosService(&body)

	if(error != nil){
		c.JSON(http.StatusBadRequest,gin.H{
			"error": "could not create data",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"result": body,
		})
	}

}