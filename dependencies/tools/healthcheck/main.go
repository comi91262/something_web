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
	dbName = "mysql"
	dbURL  = "root:password@tcp(db:3306)/"
)

func main() {
	tableName := os.Args[2]

	conn, err := sql.Open(dbName, dbURL+tableName)
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

	out, err := exec.Command(os.Args[1]).Output()
	fmt.Println(out)
}
