package models

type TodoData struct{
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Todo string `json:"todo"`
}

func (t TodoData) TableName() string{
	return "todo_list"
}