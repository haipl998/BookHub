package ginloan

import (
	"BookHub/common"
	"BookHub/module/Loan/biz"
	"BookHub/module/Loan/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetListOfLoans(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		business := biz.NewListLoanBiz(store)

		result, err := business.ListLoan(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
