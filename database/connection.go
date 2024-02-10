package database

import (
	"fmt"
	"github.com/yudisaputra/assignment-bookandlink/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Mysql() {
	cDb := config.Database()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true", cDb.User, cDb.Password, cDb.Host, cDb.Port, cDb.Name)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic(err)
	}

	Instance = db

}
