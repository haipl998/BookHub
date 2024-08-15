package ginauthor

import (
	"BookHub/common"
	"BookHub/module/author/biz"
	"BookHub/module/author/model"
	"BookHub/module/author/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAuthor(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Authors
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateAuthorBiz(store)

		if err := business.CreateAuthor(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.AuthorID))
	}
}
