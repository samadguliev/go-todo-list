package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Task   string `json:"task" gorm:"text; not null;"`
	IsDone bool   `json:"is_done" gorm:"bool; not null; default=false"`
}

type TodoAPI struct {
	ID     uint   `json:"ID"`
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}
