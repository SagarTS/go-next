package services

import (
	"todo/infrastructure"
	"todo/models"
)

func DeleteTodo(todo *models.TodoData,id string) error{
	err := infrastructure.DB.Delete(todo,id).Error

	return err
}