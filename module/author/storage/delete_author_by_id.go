package storage

import (
	"BookHub/common"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) DeleteAuthorById(ctx context.Context, cond map[string]interface{}) (err error) {
	if err = s.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("Authors").Where(cond).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		if err = tx.Table("BookAuthors").Where(cond).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		return nil
	}); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
