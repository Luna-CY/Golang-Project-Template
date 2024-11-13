package example

import "github.com/Luna-CY/Golang-Project-Template/internal/dao"

func New() *Example {
	return &Example{
		BaseDao: dao.New(),
	}
}

type Example struct {
	*dao.BaseDao
}
