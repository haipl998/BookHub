package gin_member

import (
	"BookHub/common"

	biz_member "BookHub/module/member/biz"
	model_member "BookHub/module/member/model"
	storage_member "BookHub/module/member/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model_member.MemberCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage_member.NewSQLStore(db)
		business := biz_member.NewRegisterBiz(store)

		if err := business.Register(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.MemberID))
	}
}
