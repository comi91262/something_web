package main

import (
	"fmt"
	"log"
	"net/http"
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

//type Product struct {
//	gorm.Model
//Code  string
//c	Price uint
//c}

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

func showPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(showPagePath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dsn := "root:password@tcp(db:3306)/world?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	var country Country
	// Read
	db.First(&country)
	t.Execute(w, map[string]Country{"country": country})
	//	term, err := loadTerm(dateTimePath)
	//	if err != nil {
	//		t.Execute(w, map[string]bool{"showBanner": false})
	//		return
	//	}
	//
	//	t.Execute(w, map[string]bool{"showBanner": inTerm(time.Now(), term)})
}

func main() {
	http.HandleFunc(pathTopPage, showTopPage)
	http.HandleFunc(pathShowPage, showPage)

	err := http.ListenAndServe(":"+defaultPort, nil)
	if err != nil {
		log.Fatal(err)
	}

	// gorm
	dsn := "root:password@tcp(db:3306)/world?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create
	// db.Create(&Product{Code: "D42", Price: 100})

	var country Country
	// Read
	db.First(&country)
	fmt.Printf("%v\n", country)
	//db.First(&product, 1)                 // find product with integer primary key
	//db.First(&product, "code = ?", "D42") // find product with code D42

	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//// Delete - delete product
	//db.Delete(&product, 1)

	// engine.Run(":" + defaultPort)

}
