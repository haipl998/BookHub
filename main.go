package main

import (
	"BookHub/middleware"
	ginloan "BookHub/module/Loan/transport/gin"
	ginauthor "BookHub/module/author/transport/gin"
	ginbook "BookHub/module/book/transport/gin"
	gincategory "BookHub/module/category/transport/gin"
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
		//book
		api.GET("/book", ginbook.GetListOfBooks(db))
		api.GET("/book/:id", ginbook.GetBookById(db))
		api.POST("/book", ginbook.CreateBook(db))
		api.PUT("/book/:id", ginbook.UpdateBookById(db))
		api.DELETE("/book/:id", ginbook.DeleteBookById(db))

		// category
		category := api.Group("/category")
		category.Use(middleware.OnlyAdmin())
		{
			category.GET("/", gincategory.GetCategoryOfCategories(db))
			category.GET("/:id", gincategory.GetCategoryById(db))
			category.POST("/", gincategory.CreateCategory(db))
			category.PUT("/:id", gincategory.UpdatCategoryByID(db))
			category.DELETE("/:id", gincategory.DeleteCategoryById(db))
		}

		//author
		author := api.Group("/author")
		author.Use(middleware.OnlyAdmin())
		{
			author.GET("/", ginauthor.GetListOfAuthors(db))
			author.GET("/:id", ginauthor.GetAuthorById(db))
			author.POST("/", ginauthor.CreateAuthor(db))
			author.PUT("/:id", ginauthor.UpdatAuthorByID(db))
			author.DELETE("/:id", ginauthor.DeleteAuthorById(db))
		}

		//memmber
		api.POST("/member/register", middleware.OnlyAdmin(), middleware.ValidateEmailAndPhone(), gin_member.Register(db))
		api.GET("/member", middleware.OnlyAdmin(), gin_member.GetListOfMembers(db))
		api.GET("/member/:id", middleware.AuthorizeSelf(), gin_member.GetMemberById(db))
		api.PUT("/member/:id", middleware.AuthorizeSelf(), gin_member.UpdateMemberById(db))
		api.DELETE("member/:id", middleware.OnlyAdmin(), gin_member.DeleteMemberById(db))

		//Loan
		loan := api.Group("/loan")
		loan.Use(middleware.AuthorizeSelf())
		{
			loan.GET("/", middleware.OnlyAdmin(), ginloan.GetListOfLoans(db))
			loan.GET("/:id", ginloan.GetLoanById(db))
			loan.POST("/", middleware.OnlyAdmin(), ginloan.CreatetLoan(db))
			loan.PUT("/:id", ginloan.UpdateLoan(db))
			loan.DELETE("/:id", middleware.OnlyAdmin(), ginloan.DeleteLoanById(db))
		}
	}
	router.Run()
}
