package ginmember

import (
	"BookHub/common"
	"BookHub/module/member/biz"
	"BookHub/module/member/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListOfMembers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		business := biz.NewListMemberBiz(store)

		result, err := business.ListMember(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
