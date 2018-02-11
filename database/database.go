package database

import "github.com/jinzhu/gorm"

var DB *gorm.DB
func Init(){
	var err error
	DB, err = gorm.Open("mysql", "root:3333@/stacker?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect database")
	}
	DB.SingularTable(true)
	DB.LogMode(true)
}
