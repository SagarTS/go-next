package models

import "gorm.io/gorm"

type TodoData struct{
	gorm.Model
	Todo string `json:"todo"`
}