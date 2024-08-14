package ginbook

import (
	"BookHub/common"
	"BookHub/module/book/biz"
	"BookHub/module/book/model"
	"BookHub/module/book/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateBookById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var book model.BookUpdate
		if err := c.ShouldBind(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateBookByIdBiz(store)

		if err := business.UpdateBookById(c.Request.Context(), id, &book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
