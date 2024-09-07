package ginbook

import (
	"BookHub/common"
	"BookHub/module/book/biz"
	"BookHub/module/book/model"
	"BookHub/module/book/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book model.Book
		if err := c.ShouldBind(&book); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateBookBiz(store)

		if err := business.CreateBook(c.Request.Context(), &book); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(book.BookID))
	}
}
