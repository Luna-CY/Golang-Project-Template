package example

import "github.com/Luna-CY/Golang-Project-Template/internal/interface/service"

func New(example service.Example) *Example {
	return &Example{
		example: example,
	}
}

type Example struct {
	example service.Example
}
