package services

import (
	"todo/infrastructure"
	"todo/models"
)

func AddTodosService(body *models.TodoData) error{
	result := infrastructure.DB.Create(&body)
	
	return result.Error
}