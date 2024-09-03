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

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model_member.LoginForm
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage_member.NewSQLStore(db)
		business := biz_member.NewLoginBiz(store)

		cookie, err := business.Login(c.Request.Context(), &data)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("Login successful"))
	}
}
