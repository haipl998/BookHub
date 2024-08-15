package storage

import (
	"BookHub/common"
	"BookHub/module/author/model"
	"errors"

	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetAuthor(ctx context.Context, cond map[string]interface{}) (result *model.Authors, err error) {
	db := s.db.Table(model.Authors{}.TableName())

	if err := db.First(&result, cond).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return result, nil
}
