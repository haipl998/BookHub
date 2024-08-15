package storage

import (
	"BookHub/common"
	"BookHub/module/author/model"

	"context"
)

func (s *sqlStore) ListAuthor(ctx context.Context) (result []model.Authors, err error) {
	db := s.db.Table(model.Authors{}.TableName())

	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
