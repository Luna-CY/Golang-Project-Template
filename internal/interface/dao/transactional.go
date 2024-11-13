package dao

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/transactional"
)

type Transactional interface {
	// BeginTransaction begin transaction on manual
	// provider a method to start transaction for services, do not call this method in DAO layer
	BeginTransaction(ctx icontext.Context) (transactional.Transactional, error)
}
