package storage

import (
	"BookHub/common"
	"BookHub/module/Loan/model"
	"context"
)

func (s *sqlStore) UpdateLoan(ctx context.Context, data *model.LoanUpdate, cond map[string]interface{}) (err error) {
	if err := s.db.Where(cond).Updates(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
