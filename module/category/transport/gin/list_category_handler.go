package gincategory

import (
	"BookHub/common"
	"BookHub/module/category/biz"
	"BookHub/module/category/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCategoryOfCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := storage.NewSQLStore(db)
		business := biz.NewListCategoryBiz(store)

		result, err := business.ListCategory(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
