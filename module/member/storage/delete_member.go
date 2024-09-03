package storage_member

import (
	"BookHub/common"
	"context"

	"gorm.io/gorm"
)

func (s *sqlStore) DeleteMemberById(ctx context.Context, cond map[string]interface{}) (err error) {
	if err = s.db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Table("Members").Where(cond).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		if err = tx.Table("Loans").Where(cond).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		if err = tx.Table("Reviews").Where(cond).Update("Deleted", true).Error; err != nil {
			return common.ErrDB(err)
		}
		return nil
	}); err != nil {
		return common.ErrDB(err)
	}

	return nil
}
