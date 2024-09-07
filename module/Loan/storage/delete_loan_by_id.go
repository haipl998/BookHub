package storage

import (
	"BookHub/common"
	"context"
)

func (s *sqlStore) DeleteLoan(ctx context.Context, cond map[string]interface{}) (err error) {
	if err = s.db.Table("Loans").Where(cond).Update("Deleted", true).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
