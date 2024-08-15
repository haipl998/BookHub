package storage

import (
	"BookHub/common"
	"BookHub/module/book/model"
	"context"
)

func (s *sqlStore) UpdateBook(ctx context.Context, cond map[string]interface{}, data *model.BookUpdate) (err error) {
	if err = s.db.Table(model.BookUpdate{}.TableName()).Where(cond).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
