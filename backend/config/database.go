package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DEFAULT_MAX_IDLE_CONNS = 10
	DEFAULT_MAX_OPEN_CONNS = 100
)

func ConnectToDB() *gorm.DB {
	var err error
	dsn := os.Getenv("DB_DSN")
	env := os.Getenv("ENV")
	var db *gorm.DB
	customLogger := NewCustomLogger()
	if env == "TEST" {
		db, err = gorm.Open(sqlite.Open("database_test.db"), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: customLogger, PrepareStmt: true})
	}

	if err != nil {
		log.Fatal("Error connecting to database. Error: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	maxOpenConns, err := strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	if err != nil {
		maxOpenConns = DEFAULT_MAX_OPEN_CONNS
	}
	maxIdleConns, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	if err != nil {
		maxOpenConns = DEFAULT_MAX_IDLE_CONNS
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
