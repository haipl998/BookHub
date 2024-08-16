package storage

import (
	"BookHub/common"
	"BookHub/module/category/model"
	"errors"

	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetCategory(ctx context.Context, cond map[string]interface{}) (result *model.Category, err error) {
	db := s.db.Table(model.Category{}.TableName())

	if err := db.First(&result, cond).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return result, nil
}
