package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext"
	"github.com/Luna-CY/Golang-Project-Template/internal/icontext/icontextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/dao"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/transactional"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
)

func New(transactional dao.Transactional) *BaseService {
	return &BaseService{
		transactional: transactional,
	}
}

type BaseService struct {
	transactional dao.Transactional
}

func (cls *BaseService) WithTransaction(ctx icontext.Context, call func(ctx icontext.Context) error) (err error) {
	var transaction transactional.Transactional

	if !icontextutil.CheckOnTransactional(ctx) {
		transaction, err = cls.transactional.BeginTransaction(ctx)
		if nil != err {
			return err
		}

		defer func() {
			if err := transaction.Rollback(ctx); nil != err {
				logger.SugarLogger(ctx).Errorf("I.S.Transaction.WithTransaction rollback transaction faillback. err: %v", err)
			}
		}()

		ctx = icontextutil.SetTransactional(ctx, transaction)
	}

	if err := call(ctx); nil != err {
		return err
	}

	if nil != transaction {
		if err := transaction.Commit(ctx); nil != err {
			return err
		}
	}

	return nil
}
