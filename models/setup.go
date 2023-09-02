package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func KonekDB() {
	db, err := gorm.Open(mysql.Open("root:admin@tcp(localhost:3306)/rest_gin"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	DB = db
}
