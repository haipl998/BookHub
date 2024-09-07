package storage

import (
	"BookHub/common"
	"BookHub/module/author/model"
	"errors"

	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) GetAuthor(ctx context.Context, cond map[string]interface{}) (result *model.Author, err error) {
	cond["Authors.Deleted"] = false
	db := s.db.Table(model.Author{}.TableName())

	if err := db.Where(cond).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return result, nil
}
