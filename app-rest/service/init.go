package service

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	driverName  = "mysql"
	hostEnv     = "MYSQL_HOST"
	portEnv     = "MYSQL_PORT"
	databaseEnv = "MYSQL_DATABASE"
	userEnv     = "MYSQL_USER"
	passwordEnv = "MYSQL_PASSWORD"
)

var dbConn *gorm.DB

func connectDB() (*gorm.DB, error) {
	user := os.Getenv(userEnv)
	password := os.Getenv(passwordEnv)
	host := os.Getenv(hostEnv)
	port := os.Getenv(portEnv)
	db := os.Getenv(databaseEnv)

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, db)
	conn, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func init() {
	err := errors.New("")
	dbConn, err = connectDB()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("init database ok")
}
