package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type DBConnection struct {
	Host         string
	User         string
	Password     string
	DatabaseName string
}

func (D DBConnection) SQL() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=5432 sslmode=disable TimeZone=America/Los_Angeles",
		D.Host,
		D.User,
		D.Password,
		D.DatabaseName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("error : %v", err.Error())
	}
	return db
}

type MySQLDatabase struct {
	dbconnection *DBConnection
}

func NewSQLDB(conn DBConnection) *MySQLDatabase {
	return &MySQLDatabase{dbconnection: &conn}
}
