package gin_member

import (
	"BookHub/common"

	biz_member "BookHub/module/member/biz"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		business := biz_member.NewLogoutBiz()
		cookie := business.Logout(c.Request.Context())

		http.SetCookie(c.Writer, cookie)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse("logout successful"))
	}
}
