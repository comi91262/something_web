package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// var wr = bufio.NewWriter(os.Stdout)
	engine := gin.Default()
	// ua := ""
	// ミドルウェアを使用
	engine.Use(func(c *gin.Context) {
		// ua = c.GetHeader("User-Agent")
		// fmt.Fprintf(wr, "%v\n", c.GetHeader(""))
		c.Next()
	})

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

	engine.Run(":3000")
}
