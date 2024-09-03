package gin_member

import (
	"BookHub/common"
	biz_member "BookHub/module/member/biz"
	storage_member "BookHub/module/member/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListOfMembers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := storage_member.NewSQLStore(db)
		business := biz_member.NewListMemberBiz(store)

		result, err := business.ListMember(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
