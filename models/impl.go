package models

import (
	"fmt"
	"guardTai/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	logging "github.com/op/go-logging"
)

var db *gorm.DB
var logger = logging.MustGetLogger("storage")

func NewDbBackend(cfg *setting.Database) (DBBackend, error) {
	var (
		err                                  error
		dbType, dbName, user, password, host string
	)
	dbType = cfg.Type
	dbName = cfg.DBname
	user = cfg.User
	password = cfg.Password
	host = cfg.Host
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		logger.Error("error")
		return nil, err
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db, nil
}
