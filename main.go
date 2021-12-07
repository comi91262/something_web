package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/ykonomi/something_web/rest/middleware"

	"log"
	"os"
)

const defaultPort = "3000"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// var wr = bufio.NewWriter(os.Stdout)
	engine := gin.Default()
	// ua := ""
	// ミドルウェア
	engine.Use(func(c *gin.Context) {
		// ua = c.GetHeader("User-Agent")
		// fmt.Fprintf(wr, "%v\n", c.GetHeader(""))
		c.Next()
	})
	engine.Use(middleware.RecordUaAndTime)
	engine.Static("/static", "./static")
	// engine.GET("/", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"message":    "hello world",
	//			"User-Agent": ua,
	//		})
	//	})
	// htmlのディレクトリを指定
	// engine.LoadHTMLGlob("*") //templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			// htmlに渡す変数を定義
			"message": "hello gin",
		})
	})

	//	bookEngine := engine.Group("/book")
	//	{
	//		v1 := bookEngine.Group("/v1")
	//		{
	//			v1.POST("/add", controller.BookAdd)
	//			v1.GET("/list", controller.BookList)
	//			v1.PUT("/update", controller.BookUpdate)
	//			v1.DELETE("/delete", controller.BookDelete)
	//		}
	//	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// TODO to gin
	//http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//http.Handle("/query", srv)

	//log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	//log.Fatal(http.ListenAndServe(":"+port, nil))

	// gorm
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	// Migrate the schema
	// db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

	engine.Run(":" + defaultPort)

}
