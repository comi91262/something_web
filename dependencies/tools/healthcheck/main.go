package main

import (
	"database/sql"
	"fmt"
	"os/exec"
	"time"

	// without the underscore _, you will get imported but not
	// used error message
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName  = "mysql"
	hostEnv     = "MYSQL_HOST"
	databaseEnv = "MYSQL_DATABASE"
	userEnv     = "MYSQL_USER"
	passwordEnv = "MYSQL_PASSWORD"
)

func main() {
	user := os.Getenv(userEnv)
	password := os.Getenv(passwordEnv)
	host := os.Getenv(hostEnv)
	db := os.Getenv(databaseEnv)

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, db)

	var conn *sql.DB
	var err error
	for {
		if conn, err = sql.Open(driverName, dbURL); err == nil {
			break
		}

		fmt.Printf("Mysql is unavailable: %v\n", err)
		time.Sleep(time.Second)
	}
	defer conn.Close()

	for {
		if err = conn.Ping(); err == nil {
			break
		}

		fmt.Printf("Mysql is unavailable: %v\n", err)
		time.Sleep(time.Second)
	}

	fmt.Println("Mysql is up")

	if len(os.Args) < 2 {
		os.Exit(0)
	}

	cmd := os.Args[1]
	if _, err := exec.Command(cmd).Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
