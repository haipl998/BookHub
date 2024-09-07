package storage

import (
	"BookHub/common"
	"BookHub/module/category/model"

	"context"
)

func (s *sqlStore) UpdateCategory(ctx context.Context, cond map[string]interface{}, data *model.CategoryUpdate) (err error) {
	db := s.db.Table(model.CategoryUpdate{}.TableName())
	if err := db.Where(cond).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
