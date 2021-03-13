package db

import (
	"fmt"
	"log"

	"github.com/laster18/poi/api/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

var dbInstance *Db

func NewDb() *Db {
	if dbInstance != nil {
		return dbInstance
	}

	conn, err := gorm.Open(mysql.Open(connString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}
	log.Println("success to connect db!")

	dbInstance = &Db{conn}

	return dbInstance
}

func NewTx() *Db {
	db := NewDb()
	tx := db.Begin()

	return &Db{tx}
}

func TransactAndReturnData(txFunc func(*Db) (interface{}, error)) (data interface{}, err error) {
	tx := &Db{NewDb().Begin()}
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
