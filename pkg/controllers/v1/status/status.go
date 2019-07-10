package status

import (
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

// Init - init status controller
func Init(DB *xorm.Engine) {
	db = DB
}
