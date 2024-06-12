package database

import (
	"fmt"

	"github.com/fuatanshori/gorm/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseInit() *gorm.DB {


	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",config.DB_USER,config.DB_PASSWORD,config.DB_HOST,config.DB_PORT,config.DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
