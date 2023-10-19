package models

type TodoData struct{
	Todo string `json:"todo"`
}

func (t TodoData) TableName() string{
	return "todo_list"
}