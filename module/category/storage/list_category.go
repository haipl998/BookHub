package storage

import (
	"BookHub/common"
	"BookHub/module/category/model"

	"context"
)

func (s *sqlStore) ListCategory(ctx context.Context) (result *[]model.Category, err error) {
	db := s.db.Table(model.Category{}.TableName())
	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
