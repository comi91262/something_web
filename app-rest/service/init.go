package service

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbConn *gorm.DB

const dsn = "root:password@tcp(db:3306)/world?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	err := errors.New("")
	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("init data base ok")
}
