package models

import (
	"time"
)

// Migrations - migrations
type Migrations struct {
	ID    int64     `xorm:"'id' pk autoincr" json:"id" schema:"id"`
	Name  string    `xorm:"name" json:"name" schema:"name"`
	RunOn time.Time `xorm:"run_on" json:"run_on" schema:"run_on"`
}

// TableName - migrations table name
func (t Migrations) TableName() string {
	return "migrations"
}

// SetID - setting id
func (t Migrations) SetID(id int64) {
	t.ID = id
}

// GetID - gets id
func (t Migrations) GetID() int64 {
	return t.ID
}

// Users - generated user struct
type Users struct {
	ID       int64  `xorm:"'id' pk autoincr" json:"id" schema:"id"`
	First    string `xorm:"first" json:"first" schema:"first"`
	Last     string `xorm:"last" json:"last" schema:"last"`
	Email    string `xorm:"email" json:"email" schema:"email"`
	Password string `xorm:"password" json:"password" schema:"password"`
}
