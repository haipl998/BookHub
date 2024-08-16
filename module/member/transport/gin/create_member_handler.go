package ginmember

import (
	"BookHub/common"
	"BookHub/module/member/biz"
	"BookHub/module/member/model"
	"BookHub/module/member/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMember(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Member
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateMemberBiz(store)

		if err := business.CreateMember(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.MemberID))
	}
}
