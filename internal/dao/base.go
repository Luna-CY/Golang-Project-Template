package dao

import (
	"fmt"
	"github.com/Luna-CY/Golang-Project-Template/internal/configuration"
	"github.com/Luna-CY/Golang-Project-Template/internal/context"
	"github.com/Luna-CY/Golang-Project-Template/internal/context/contextutil"
	"github.com/Luna-CY/Golang-Project-Template/internal/errors"
	"github.com/Luna-CY/Golang-Project-Template/internal/interface/transactional"
	"github.com/Luna-CY/Golang-Project-Template/internal/logger"
	transactional2 "github.com/Luna-CY/Golang-Project-Template/internal/transactional"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	mo sync.Once
	mi *gorm.DB
)

func New() *BaseDao {
	return &BaseDao{}
}

type BaseDao struct{}

func (cls *BaseDao) GetDb(ctx context.Context) *gorm.DB {
	if transaction, ok := contextutil.GetTransactional(ctx); ok {
		return transaction.Session()
	}

	return cls.mysql(ctx).WithContext(ctx)
}

func (cls *BaseDao) BeginTransaction(ctx context.Context) (transactional.Transactional, errors.Error) {
	trans, ok := contextutil.GetTransactional(ctx)
	if ok {
		return trans, nil
	}

	var db = cls.GetDb(ctx).Begin()
	if nil != db.Error {
		logger.SugarLogger(ctx).Errorf("I.D.BaseDao.BeginTransaction start transaction error: %v", db.Error)

		return nil, errors.New(errors.ErrorTypeServerInternalError, "ID_AO.BD_AO.BT_ON.47", "start transaction error: %v", db.Error)
	}

	return transactional2.New(db), nil
}

func (cls *BaseDao) mysql(ctx context.Context) *gorm.DB {
	mo.Do(func() {
		var config gorm.Config
		config.DisableAutomaticPing = false

		var err error
		mi, err = gorm.Open(mysql.Open(configuration.Configuration.Database.Mysql.Dsn), &config)
		if err != nil {
			panic(fmt.Sprintf("I.D.BaseDao.mysql connection database error: %s", err))
		}

		if configuration.Configuration.Database.Mysql.ConnPool.Enable && 0 != configuration.Configuration.Database.Mysql.ConnPool.MaxIdleConn && 0 != configuration.Configuration.Database.Mysql.ConnPool.MaxOpenConn {
			logger.SugarLogger(ctx).Infof("enable mysql connections pool，Max Idle Conn: %d, Max Open Conn: %d, Max Idle Conn Life Time: %d min", configuration.Configuration.Database.Mysql.ConnPool.MaxIdleConn, configuration.Configuration.Database.Mysql.ConnPool.MaxOpenConn, configuration.Configuration.Database.Mysql.ConnPool.MaxIdleLifeTime)

			driver, err := mi.DB()
			if nil != err {
				panic(fmt.Sprintf("I.D.BaseDao.mysql connection database error: %s", err))
			}

			driver.SetMaxIdleConns(configuration.Configuration.Database.Mysql.ConnPool.MaxIdleConn)
			driver.SetMaxOpenConns(configuration.Configuration.Database.Mysql.ConnPool.MaxOpenConn)
			driver.SetConnMaxIdleTime(time.Duration(configuration.Configuration.Database.Mysql.ConnPool.MaxIdleLifeTime) * time.Minute)
		}
	})

	return mi
}
