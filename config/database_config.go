package config

import (
	"belajar-golang/exception"
	"belajar-golang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb(env ConfigEnv) *gorm.DB {
	databaseUrl := env.Get("MYSQL_CONNECTION")

	db, err := gorm.Open(mysql.Open(databaseUrl), &gorm.Config{})
	exception.PanicIfNeeded(err)

	errDb := db.AutoMigrate(&model.Employee{})
	exception.PanicIfNeeded(errDb)
	return db
}
