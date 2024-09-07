package storage

import (
	"BookHub/common"
	"BookHub/module/author/model"

	"context"
)

func (s *sqlStore) ListAuthor(ctx context.Context) (result []model.Author, err error) {
	cond := make(map[string]interface{})
	cond["Authors.Deleted"] = false
	db := s.db.Table(model.Author{}.TableName())

	if err := db.Where(cond).Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return result, nil
}
