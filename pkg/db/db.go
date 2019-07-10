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

// Find - finds models
func Find(DB *xorm.Engine, findBy, objects interface{}) (err error) {
	return DB.Find(objects, findBy)
}

// FindBy - finds by conditions
func FindBy(DB *xorm.Engine, model interface{}) (err error) {
	_, err = DB.Get(model)
	return
}

// Exists - check if model exists
func Exists(DB *xorm.Engine, model interface{}) (bool, error) {
	return DB.Get(model)
}

// Update - updates model
func Update(DB *xorm.Engine, id int64, model interface{}) (err error) {
	_, err = DB.ID(id).Update(model)
	return
}

// Store - inserts model
func Store(DB *xorm.Engine, model interface{}) (err error) {
	_, err = DB.Insert(model)
	return
}

// Destroy - deletes model
func Destroy(DB *xorm.Engine, id int64, model interface{}) (err error) {
	_, err = DB.ID(id).Delete(model)
	return
}
