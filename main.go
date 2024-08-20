package main

import (
	ginbook "BookHub/module/book/transport/gin"
	ginmember "BookHub/module/member/transport/gin"
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
		api.GET("/book", ginbook.GetListOfBooks(db))        // get list book
		api.GET("/book/:id", ginbook.GetBookById(db))       // get book by id
		api.POST("/book", ginbook.CreateBook(db))           // create new book
		api.PUT("/book/:id", ginbook.UpdateBookById(db))    // update book
		api.DELETE("/book/:id", ginbook.DeleteBookById(db)) // delete book

		//memmber
		api.GET("/member", ginmember.GetListOfMembers(db))
		api.GET("/member/:id", ginmember.GetMemberById(db))
		api.POST("/member", ginmember.CreateMember(db))
		api.PUT("/member/:id", ginmember.UpdateMemberById(db))
		api.DELETE("member/:id", ginmember.DeleteMemberById(db))
	}
	router.Run()
}
