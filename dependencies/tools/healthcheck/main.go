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
	dbName      = "mysql"
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

	if len(os.Args) < 1 {
		fmt.Println("Too few arguments")
		os.Exit(1)
	}
	cmd := os.Args[1]

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, db)

	conn, err := sql.Open(dbName, dbURL)
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		if err := conn.Ping(); err != nil {
			fmt.Println(err)
			fmt.Println("Mysql is unavailable - sleeping")
			break
		}
		time.Sleep(time.Minute * 1)
	}
	fmt.Println("Mysql is up")

	if _, err := exec.Command(cmd).Output(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
