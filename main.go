package main

import (
	"BookHub/middleware"
	ginbook "BookHub/module/book/transport/gin"
	gin_member "BookHub/module/member/transport/gin"
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

	router.POST("/login", gin_member.Login(db))
	router.GET("/logout", middleware.AuthenticateJWT(), gin_member.Logout())

	api := router.Group("/api")
	api.Use(middleware.AuthenticateJWT())
	{
		api.GET("/book", ginbook.GetListOfBooks(db))                                // get list book
		api.GET("/book/:id", ginbook.GetBookById(db))                               // get book by id
		api.POST("/book", middleware.OnlyAdmin(), ginbook.CreateBook(db))           // create new book
		api.PUT("/book/:id", middleware.OnlyAdmin(), ginbook.UpdateBookById(db))    // update book
		api.DELETE("/book/:id", middleware.OnlyAdmin(), ginbook.DeleteBookById(db)) // delete book

		//memmber
		api.POST("/member/register", middleware.OnlyAdmin(), gin_member.Register(db))
		api.GET("/member", middleware.OnlyAdmin(), gin_member.GetListOfMembers(db))
		api.GET("/member/:id", middleware.AuthorizeSelf(), gin_member.GetMemberById(db)) // cần xem set lại
		//api.POST("/member", middleware.OnlyAdmin(), gin_member.CreateMember(db))            // cần xem set lại
		api.PUT("/member/:id", middleware.AuthorizeSelf(), gin_member.UpdateMemberById(db)) // cần xem set lại
		api.DELETE("member/:id", middleware.OnlyAdmin(), gin_member.DeleteMemberById(db))
	}
	router.Run()
}
