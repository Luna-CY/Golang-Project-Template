package example_service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/service"
)

func New(example dao.Example) *Example {
	return &Example{
		BaseService: service.New(example),
		example:     example,
	}
}

type Example struct {
	*service.BaseService
	example dao.Example
}
