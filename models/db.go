package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vm/config"
)

var db *gorm.DB

func InitDB() error {
	var err error
	db, err = gorm.Open("mysql", config.GetMysqlURI())
	if err != nil {
		return err
	}
	if db.HasTable("instances") == false {
		db.CreateTable(&Instance{})
	}
	db.LogMode(true)
	return nil
}
