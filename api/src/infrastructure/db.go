package infrastructure

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

func NewDb() *Db {
	conn, err := gorm.Open(mysql.Open(connString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection error: %v", err)
	}
	log.Println("success to connect db!")

	return &Db{conn}
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
