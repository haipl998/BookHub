package ginloan

import (
	"BookHub/common"
	"BookHub/module/Loan/biz"
	"BookHub/module/Loan/model"
	"BookHub/module/Loan/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateLoan(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		//TODO: chua nghi ra phuong phap su ly
		var data model.LoanUpdate
		// if err := c.ShouldBind(&data); err != nil {
		// 	c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		// 	return
		// }

		store := storage.NewSQLStore(db)
		business := biz.NewUpdateLoanBiz(store)

		if err := business.UpdateLoanById(c.Request.Context(), &data, id); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("true"))
	}
}
