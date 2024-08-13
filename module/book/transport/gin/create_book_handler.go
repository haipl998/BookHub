package ginbook

import (
	"BookHub/module/book/biz"
	"BookHub/module/book/model"
	"BookHub/module/book/storage"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book
		if err := c.ShouldBind(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		book.Title = strings.TrimSpace(book.Title)
		if book.Title == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "title cannot be blank"})
			return
		}

		book.CategoryName = strings.TrimSpace(book.CategoryName)
		if book.CategoryName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "CategoryName cannot be blank"})
			return
		}

		book.FirstName = strings.TrimSpace(book.FirstName)
		if book.FirstName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FirstName cannot be blank"})
			return
		}

		book.LastName = strings.TrimSpace(book.LastName)
		if book.LastName == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "LastName cannot be blank"})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateBookBiz(store)

		if err := business.CreateBook(c.Request.Context(), &book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": book.BookID})
	}
}
