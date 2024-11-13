package transactional

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"gorm.io/gorm"
	"sync"
)

func New(db *gorm.DB) *Transactional {
	return &Transactional{
		db: db,
	}
}

type Transactional struct {
	db *gorm.DB

	mutex sync.Mutex
	flag  bool

	rollbackHooks []func(ctx icontext.Context)
	commitHooks   []func(ctx icontext.Context)
}
