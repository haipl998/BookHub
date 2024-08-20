package storage

import (
	"BookHub/common"
	"BookHub/module/Loan/model"

	"context"
)

func (s *sqlStore) CreateLoan(ctx context.Context, data *model.LoanCreation) (err error) {
	cond := make(map[string]interface{})
	cond["Loans.Deleted"] = false
	if err := s.db.Where(cond).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
