package service

import (
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
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

func (cls *BaseService) WithTransaction(ctx context.Context, call func(ctx context.Context) errors.Error) (err errors.Error) {
	var transaction transactional.Transactional

	if !contextutil.CheckOnTransactional(ctx) {
		transaction, err = cls.transactional.BeginTransaction(ctx)
		if nil != err {
			return err.Relation(errors.ErrorServerInternalError("IS_CE.BS_CE.WT_ON.283257"))
		}

		defer func() {
			if err := transaction.Rollback(ctx); nil != err {
				logger.SugarLogger(logger.WithRequestId(ctx), logger.WithStack()).Errorf("I.S.Transaction.WithTransaction rollback transaction faillback. err: %v", err)
			}
		}()

		ctx = contextutil.SetTransactional(ctx, transaction)
	}

	if err := call(ctx); nil != err {
		return err.Relation(errors.ErrorServerInternalError("IS_CE.BS_CE.WT_ON.413302"))
	}

	if nil != transaction {
		if err := transaction.Commit(ctx); nil != err {
			return err.Relation(errors.ErrorServerInternalError("IS_CE.BS_CE.WT_ON.463305"))
		}
	}

	return nil
}
