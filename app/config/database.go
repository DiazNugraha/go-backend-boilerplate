package config

import (
	"backend/app/helpers"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConn() (db gorm.Dialector) {
	dbHost := helpers.GetEnv("DB_HOST", "localhost")
	dbName := helpers.GetEnv("DB_NAME", "secret_chat")
	dbUser := helpers.GetEnv("DB_USER", "root")
	dbPass := helpers.GetEnv("DB_PASS", "root")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db = mysql.Open(dsn)
	return
}
