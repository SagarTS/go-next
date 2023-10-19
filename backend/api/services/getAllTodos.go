package services

import (
	"todo/infrastructure"
	"todo/models"
)

func GetAllTodos(todos *[]models.TodoData)error{
	result := infrastructure.DB.Find(todos)

	return result.Error
}