package home

import (
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

// Init - init home contoller
func Init(DB *xorm.Engine) {
	db = DB
}
