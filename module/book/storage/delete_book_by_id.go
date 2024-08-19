package storage

import (
	"BookHub/common"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) DeleteBookById(ctx context.Context, id int) (err error) {
	if err = s.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("BookAuthors").Where("BookID = ?", id).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		if err = tx.Table("Loans").Where("BookID = ?", id).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		if err = tx.Table("Reviews").Where("BookID = ?", id).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		if err = tx.Table("Books").Where("BookID = ?", id).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		return nil
	}); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
