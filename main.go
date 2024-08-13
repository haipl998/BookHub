package main

import (
	ginbook "BookHub/module/book/transport/gin"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:rootpassword@tcp(db:3306)/bookhub?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL:", err)
	}

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/book", ginbook.GetListOfBooks(db)) // get list book
		api.GET("/book/:id", ginbook.GetBook(db))    // get book by id
		api.POST("/book", ginbook.CreateBook(db))    // create new book
	}
	log.Print(api)

	router.Run()
}
