package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ykonomi/something_web/controller"
	"github.com/ykonomi/something_web/middleware"
	_ "gorm.io/driver/mysql"
)

const defaultPort = "8080"

func main() {
	engine := gin.Default()
	// ミドルウェア
	engine.Use(middleware.RecordUaAndTime)
	// CRUD 書籍
	bookEngine := engine.Group("/book")
	{
		v1 := bookEngine.Group("/v1")
		{
			v1.POST("/add", controller.BookAdd)
			v1.GET("/list", controller.BookList)
			v1.PUT("/update", controller.BookUpdate)
			v1.DELETE("/delete", controller.BookDelete)
		}
	}
	engine.Run(":" + defaultPort)
}
