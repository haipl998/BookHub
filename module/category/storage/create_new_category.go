package storage

import (
	"BookHub/common"
	"BookHub/module/category/model"
	"context"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *model.Category) (err error) {
	db := s.db.Table(model.Category{}.TableName())

	var existingCategory model.Category
	result := db.Where("CategoryName = ?", data.CategoryName).Find(&existingCategory)

	if result.RowsAffected > 0 {
		return common.EntityExisted
	}

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
