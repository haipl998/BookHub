package storage

import (
	"BookHub/common"
	"BookHub/module/author/model"

	"context"
)

func (s *sqlStore) CreateAuthor(ctx context.Context, data *model.Author) (err error) {
	db := s.db.Table(model.Author{}.TableName())
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
