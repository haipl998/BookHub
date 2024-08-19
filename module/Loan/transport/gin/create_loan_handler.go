package ginloan

import (
	"BookHub/common"
	"BookHub/module/Loan/biz"
	"BookHub/module/Loan/model"
	"BookHub/module/Loan/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatetLoan(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.LoanCreation
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := storage.NewSQLStore(db)
		business := biz.NewCreateLoanBiz(store)

		if err := business.CreateLoan(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.LoanID))
	}
}
