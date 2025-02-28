//go:build wireinject
// +build wireinject

package handler

import (
	"github.com/Luna-CY/Golang-Project-Template/server/http/gateway/web/handler/example"
	"github.com/Luna-CY/Golang-Project-Template/server/http/service"
	"github.com/google/wire"
)

func NewExample() *example.Example {
	panic(wire.Build(
		service.ExampleService, example.New,
	))
}
