package main

import (
	"database/sql"
	"fmt"
	"os/exec"

	// without the underscore _, you will get imported but not
	// used error message
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName = "mysql"
	dbURL  = "root:password@tcp(db:3306)/world"
)

func main() {

	conn, err := sql.Open(dbName, dbURL)
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := conn.Ping(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// if no error. Ping is successful
	fmt.Println("Ping to database successful, connection is still alive")

	out, err := exec.Command(os.Args[1]).Output()
	fmt.Println(out)
}
