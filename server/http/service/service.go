package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/dao/example_dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/service"
	"github.com/Luna-CY/Golang-Project-Template/internal/service/example_service"
	"github.com/google/wire"
)

var ExampleService = wire.NewSet(
	example_dao.New, wire.Bind(new(dao.Example), new(*example_dao.Example)),
	example_service.New, wire.Bind(new(service.Example), new(*example_service.Example)),
)
