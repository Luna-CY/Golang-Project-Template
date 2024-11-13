//go:build wireinject
// +build wireinject

package handler

import (
	example2 "github.com/Luna-CY/Golang-Project-Template/internal/dao/example"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/service"
	"github.com/Luna-CY/Golang-Project-Template/internal/server/http/web/handler/example"
	example3 "github.com/Luna-CY/Golang-Project-Template/internal/service/example"
	"github.com/google/wire"
)

func NewExample() *example.Example {
	panic(wire.Build(
		example2.New, wire.Bind(new(dao.Example), new(*example2.Example)),
		example3.New, wire.Bind(new(service.Example), new(*example3.Example)),
		example.New,
	))
}
