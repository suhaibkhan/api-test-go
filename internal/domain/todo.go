package domain

import (
	"fmt"
)

type Todo struct {
	Id   uint64 `json:"id" gorm:"primary_key"`
	Item string `json:"item"`
	Done bool   `json:"done"`
}

func (todo *Todo) ToString() string {
	return fmt.Sprintf("%+v\n", todo)
}
