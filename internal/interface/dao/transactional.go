package dao

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/transactional"
)

type Transactional interface {
	// BeginTransaction 手动开始事务
	// 提供一个方法给服务层开始事务，不要在 DAO 层调用此方法
	BeginTransaction(ctx context.Context) (transactional.Transactional, errors.Error)
}
