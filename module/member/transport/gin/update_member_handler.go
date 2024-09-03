package gin_member

import (
	"BookHub/common"
	biz_member "BookHub/module/member/biz"
	model_member "BookHub/module/member/model"
	storage_member "BookHub/module/member/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateMemberById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		var data model_member.MemberUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage_member.NewSQLStore(db)
		business := biz_member.NewUpdateMemberBiz(store)

		if err := business.UpdateMember(c.Request.Context(), &data, id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("updata member successful"))
	}
}
