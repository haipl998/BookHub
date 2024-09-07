package ginauthor

import (
	"BookHub/common"
	"BookHub/module/author/biz"
	"BookHub/module/author/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListOfAuthors(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		business := biz.NewListAuthorBiz(store)

		result, err := business.ListAuthor(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
