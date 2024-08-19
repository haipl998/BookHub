package ginloan

import (
	"BookHub/common"
	"BookHub/module/Loan/biz"
	"BookHub/module/Loan/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetLoanById(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewGetLoanBiz(store)

		result, err := business.GetLoan(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
