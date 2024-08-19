package storage

import (
	"BookHub/common"
	"BookHub/module/author/model"

	"context"
)

func (s *sqlStore) UpdateAuthorById(ctx context.Context, data *model.AuthorUpdate, cond map[string]interface{}) (err error) {
	db := s.db.Table(model.AuthorUpdate{}.TableName())
	if err := db.Where(cond).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
