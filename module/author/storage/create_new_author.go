package storage

import (
	"BookHub/common"
	"BookHub/module/author/model"

	"context"
)

func (s *sqlStore) CreateAuthor(ctx context.Context, data *model.Authors) (err error) {
	db := s.db.Table(model.Authors{}.TableName())
	var existingAuthor model.Authors
	result := db.Where("FirstName = ? AND LastName = ?", data.FirstName, data.LastName).Find(&existingAuthor)

	if result.RowsAffected > 0 {
		return common.EntityExisted
	}
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
