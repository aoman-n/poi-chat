package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/laster18/poi/api/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

func NewDb() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}

	conn, err := gorm.Open(mysql.Open(connString()), &gorm.Config{
		Logger: newLogger(),
	})
	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}
	log.Println("success to connect db!")

	dbInstance = conn

	return dbInstance
}

func NewTx() *gorm.DB {
	db := NewDb()
	tx := db.Begin()

	return tx
}

func TransactAndReturnData(txFunc func(*gorm.DB) (interface{}, error)) (data interface{}, err error) {
	tx := NewDb().Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit().Error
		}
	}()

	return txFunc(tx)
}

func connString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Db.User,
		config.Conf.Db.Password,
		config.Conf.Db.Host,
		config.Conf.Db.Port,
		config.Conf.Db.Name,
	)
}

func newLogger() logger.Interface {
	loggerConfig := logger.Config{
		SlowThreshold: time.Second,
		Colorful:      false,
	}
	if config.Conf.GoEnv == "development" {
		loggerConfig.LogLevel = logger.Info
	} else {
		loggerConfig.LogLevel = logger.Silent
	}

	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		loggerConfig,
	)
}
