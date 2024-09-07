package storage

import (
	"BookHub/common"
	"BookHub/module/category/model"

	"context"
)

func (s *sqlStore) ListCategory(ctx context.Context) (result *[]model.Category, err error) {
	cond := make(map[string]interface{})
	cond["Categories.Deleted"] = false

	db := s.db.Table(model.Category{}.TableName())

	if err := db.Where(cond).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
