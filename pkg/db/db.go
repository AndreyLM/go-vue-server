package db

import (
	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// Connect - creates connection to db
func Connect() (db *xorm.Engine, err error) {
	db, err = xorm.NewEngine("mysql", "user:password@tcp(localhost:3310)/test?charset=utf8")
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		return
	}
	return
}
