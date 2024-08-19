package storage

import (
	"BookHub/common"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) DeleteCategory(ctx context.Context, cond map[string]interface{}) (err error) {
	if err = s.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("Categories").Where(cond).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}

		if err = tx.Table("Books").Where(cond).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}

		if err := tx.Table("BookAuthors").Where("BookID IN (?)", tx.Table("Books").Select("BookID").Where(cond)).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}

		if err := tx.Table("Loans").Where("BookID IN (?)", tx.Table("Books").Select("BookID").Where(cond)).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}

		if err := tx.Table("Reviews").Where("BookID IN (?)", tx.Table("Books").Select("BookID").Where(cond)).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}

		return nil
	}); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
