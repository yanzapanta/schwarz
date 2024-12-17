package config

import (
	"log"
	"os"

	"strings"

	"github.com/sirupsen/logrus"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbMap *gorm.DB

func DB() *gorm.DB {
	if dbMap == nil {
		dbMap = connectDatabase()
	}
	return dbMap
}

func connectDatabase() *gorm.DB {
	logLevel := logger.Silent
	connString := getConnectionString()
	var err error
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		logrus.Errorf("Error in connecting to the database: %s", err.Error())
		panic("Failed to connect to the database!")
	}

	log.Println("Database connection established")
	return db
}

func getConnectionString() string {
	var connString strings.Builder
	connString.WriteString(os.Getenv("DB_USER"))
	connString.WriteString(":")
	connString.WriteString(os.Getenv("DB_PASSWORD"))
	connString.WriteString("@tcp(")
	connString.WriteString(os.Getenv("DB_HOST"))
	connString.WriteString(":")
	connString.WriteString(os.Getenv("DB_PORT"))
	connString.WriteString(")/")
	connString.WriteString(os.Getenv("DB_NAME"))
	connString.WriteString("?charset=utf8")
	connString.WriteString("&parseTime=True")
	connString.WriteString("&loc=Local")
	return connString.String()
}
