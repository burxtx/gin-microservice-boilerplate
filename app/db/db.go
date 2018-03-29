package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var dbSession *xorm.Engine

func Init(dataSourceName string) (*xorm.Engine, error) {
	if dbSession == nil {
		var err error
		dbSession, err = xorm.NewEngine("mysql", dataSourceName)
		if err != nil {
			return nil, err
		}

		if err = dbSession.Ping(); err != nil {
			return nil, err
		}
	}
	return dbSession, nil
}
