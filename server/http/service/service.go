package service

import (
	example2 "github.com/Luna-CY/Golang-Project-Template/internal/dao/example"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/service"
	example3 "github.com/Luna-CY/Golang-Project-Template/internal/service/example"
	"github.com/google/wire"
)

var ExampleService = wire.NewSet(
	example2.New, wire.Bind(new(dao.Example), new(*example2.Example)),
	example3.New, wire.Bind(new(service.Example), new(*example3.Example)),
)
