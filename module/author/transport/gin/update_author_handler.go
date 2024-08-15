package ginauthor

import (
	"BookHub/common"
	"BookHub/module/author/biz"
	"BookHub/module/author/model"
	"BookHub/module/author/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdatAuthorByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		var data model.AuthorUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		log.Print(data)

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateAuthorBiz(store)

		if err := business.UpdateAuthor(c.Request.Context(), &data, id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
