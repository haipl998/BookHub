package gin_member

import (
	"BookHub/common"

	biz_member "BookHub/module/member/biz"
	storage_member "BookHub/module/member/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteMemberById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage_member.NewSQLStore(db)
		business := biz_member.NewDeleteMemberBiz(store)

		if err := business.DeleteMemberByID(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
