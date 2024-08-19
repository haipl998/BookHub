package storage

import (
	"BookHub/common"
	"BookHub/module/category/model"
	"context"
)

func (s *sqlStore) CreateCategory(ctx context.Context, data *model.Category) (err error) {
	db := s.db.Table(model.Category{}.TableName())

	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
