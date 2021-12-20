package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	defaultPort  = "3000"
	pathTopPage  = "/"
	pathShowPage = "/show"

	indexPagePath = "static/index.html"
	showPagePath  = "static/list.html"

	driverName  = "mysql"
	hostEnv     = "MYSQL_HOST"
	databaseEnv = "MYSQL_DATABASE"
	userEnv     = "MYSQL_USER"
	passwordEnv = "MYSQL_PASSWORD"
)

type Country struct {
	Code           string  `gorm:"primaryKey; char(3); not null; default ''"`
	Name           string  `gorm:"char(52); not null; default ''"`
	Continent      string  `gorm:"enum('Asia','Europe','North America','Africa','Oceania','Antarctica','South America'); not null; default 'Asia'"`
	Region         string  `gorm:"char(26); not null; default ''"`
	SurfaceArea    float64 `gorm:"decimal(10,2); not null; default '0.00'"`
	IndepYear      int     `gorm:"smallint; default null"`
	Population     int     `gorm:"not null; default '0'"`
	LifeExpectancy float64 `gorm:"decimal(3,1); default null"`
	GNP            float64 `gorm:"decimal(10,2); default null"`
	GNPOld         float64 `gorm:"decimal(10,2); default null"`
	LocalName      string  `gorm:"char(45); not null; default ''"`
	GovernmentForm string  `gorm:"char(45); not null; default ''"`
	HeadOfState    string  `gorm:"char(60); default null"`
	Capital        int     `gorm:"int; default null"`
	Code2          string  `gorm:"char(2); not null; default ''"`
}

func showTopPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(indexPagePath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	t.Execute(w, map[int]int{})

	//	term, err := loadTerm(dateTimePath)
	//	if err != nil {
	//		t.Execute(w, map[string]bool{"showBanner": false})
	//		return
	//	}
	//
	//	t.Execute(w, map[string]bool{"showBanner": inTerm(time.Now(), term)})
}

func connectDB() (*gorm.DB, error) {
	user := os.Getenv(userEnv)
	password := os.Getenv(passwordEnv)
	host := os.Getenv(hostEnv)
	db := os.Getenv(databaseEnv)

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, db)
	conn, err := gorm.Open(mysql.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func showPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(showPagePath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	conn, err := connectDB()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var country Country
	conn.First(&country)
	t.Execute(w, map[string]Country{"country": country})
}

func main() {
	http.HandleFunc(pathTopPage, showTopPage)
	http.HandleFunc(pathShowPage, showPage)

	err := http.ListenAndServe(":"+defaultPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}
