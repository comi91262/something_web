package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/ykonomi/something_web/rest/middleware"

	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ykonomi/something_web/graph"
	"github.com/ykonomi/something_web/graph/generated"
)

const defaultPort = "3000"

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

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// TODO to gin
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	engine.Run(":" + defaultPort)
}
