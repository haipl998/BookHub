package ginbook

import (
	"BookHub/common"
	"BookHub/module/book/biz"
	"BookHub/module/book/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListOfBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		business := biz.NewListBookBiz(store)

		result, err := business.ListBook(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
